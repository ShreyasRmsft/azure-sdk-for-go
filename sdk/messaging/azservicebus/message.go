// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-amqp-common-go/v3/uuid"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/go-amqp"
	"github.com/devigned/tab"
)

type (
	// ReceivedMessage is a received message from a Client.NewReceiver().
	ReceivedMessage struct {
		Message

		LockToken              [16]byte
		DeliveryCount          uint32
		LockedUntil            *time.Time // `mapstructure:"x-opt-locked-until"`
		SequenceNumber         *int64     // :"x-opt-sequence-number"`
		EnqueuedSequenceNumber *int64     // :"x-opt-enqueue-sequence-number"`
		EnqueuedTime           *time.Time // :"x-opt-enqueued-time"`
		DeadLetterSource       *string    // :"x-opt-deadletter-source"`

		// available in the raw AMQP message, but not exported by default
		// GroupSequence  *uint32

		rawAMQPMessage *amqp.Message

		// deferred indicates we received it using ReceiveDeferredMessages. These messages
		// will still go through the normal Receiver.Settle functions but internally will
		// always be settled with the management link.
		deferred bool
	}

	// Message is a SendableMessage which can be sent using a Client.NewSender().
	Message struct {
		ID string

		ContentType   string
		CorrelationID string
		// Body corresponds to the first []byte array in the Data section of an AMQP message.
		Body             []byte
		SessionID        *string
		Subject          string
		ReplyTo          string
		ReplyToSessionID string
		To               string
		TimeToLive       *time.Duration

		PartitionKey            *string
		TransactionPartitionKey *string
		ScheduledEnqueueTime    *time.Time

		ApplicationProperties map[string]interface{}
		Format                uint32
	}
)

// Service Bus custom properties
const (
	// DeliveryAnnotation properties
	lockTokenDeliveryAnnotation = "x-opt-lock-token"

	// Annotation properties
	partitionKeyAnnotation           = "x-opt-partition-key"
	viaPartitionKeyAnnotation        = "x-opt-via-partition-key"
	scheduledEnqueuedTimeAnnotation  = "x-opt-scheduled-enqueue-time"
	lockedUntilAnnotation            = "x-opt-locked-until"
	sequenceNumberAnnotation         = "x-opt-sequence-number"
	enqueuedTimeAnnotation           = "x-opt-enqueued-time"
	deadLetterSourceAnnotation       = "x-opt-deadletter-source"
	enqueuedSequenceNumberAnnotation = "x-opt-enqueue-sequence-number"
)

func (m *Message) messageType() string {
	return "Message"
}

func (m *Message) toAMQPMessage() *amqp.Message {
	amqpMsg := amqp.NewMessage(m.Body)

	if m.TimeToLive != nil {
		if amqpMsg.Header == nil {
			amqpMsg.Header = new(amqp.MessageHeader)
		}
		amqpMsg.Header.TTL = *m.TimeToLive
	}

	// TODO: I don't think this should be strictly required. Need to
	// look into why it won't send properly without one.
	var messageID = m.ID

	if messageID == "" {
		uuid, err := uuid.NewV4()

		if err == nil {
			messageID = uuid.String()
		}
	}

	amqpMsg.Properties = &amqp.MessageProperties{
		MessageID: messageID,
	}

	if m.SessionID != nil {
		amqpMsg.Properties.GroupID = *m.SessionID
	}

	// if m.GroupSequence != nil {
	// 	amqpMsg.Properties.GroupSequence = *m.GroupSequence
	// }

	amqpMsg.Properties.CorrelationID = m.CorrelationID
	amqpMsg.Properties.ContentType = m.ContentType
	amqpMsg.Properties.Subject = m.Subject
	amqpMsg.Properties.To = m.To
	amqpMsg.Properties.ReplyTo = m.ReplyTo
	amqpMsg.Properties.ReplyToGroupID = m.ReplyToSessionID

	if len(m.ApplicationProperties) > 0 {
		amqpMsg.ApplicationProperties = make(map[string]interface{})
		for key, value := range m.ApplicationProperties {
			amqpMsg.ApplicationProperties[key] = value
		}
	}

	amqpMsg.Annotations = map[interface{}]interface{}{}

	if m.PartitionKey != nil {
		amqpMsg.Annotations[partitionKeyAnnotation] = *m.PartitionKey
	}

	if m.TransactionPartitionKey != nil {
		amqpMsg.Annotations[viaPartitionKeyAnnotation] = *m.TransactionPartitionKey
	}

	if m.ScheduledEnqueueTime != nil {
		amqpMsg.Annotations[scheduledEnqueuedTimeAnnotation] = *m.ScheduledEnqueueTime
	}

	// TODO: These are 'received' message properties so I believe their inclusion here was just an artifact of only
	// having one message type.

	// if m.SystemProperties != nil {
	// 	// Set the raw annotations first (they may be nil) and add the explicit
	// 	// system properties second to ensure they're set properly.
	// 	amqpMsg.Annotations = addMapToAnnotations(amqpMsg.Annotations, m.SystemProperties.Annotations)

	// 	sysPropMap, err := encodeStructureToMap(m.SystemProperties)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	amqpMsg.Annotations = addMapToAnnotations(amqpMsg.Annotations, sysPropMap)
	// }

	// if m.LockToken != nil {
	// 	if amqpMsg.DeliveryAnnotations == nil {
	// 		amqpMsg.DeliveryAnnotations = make(amqp.Annotations)
	// 	}
	// 	amqpMsg.DeliveryAnnotations[lockTokenName] = *m.LockToken
	// }

	return amqpMsg
}

// newReceivedMessage creates a received message from an AMQP message.
// NOTE: this converter assumes that the Body of this message will be the first
// serialized byte array in the Data section of the messsage.
func newReceivedMessage(ctxForLogging context.Context, amqpMsg *amqp.Message) *ReceivedMessage {
	msg := &ReceivedMessage{
		Message: Message{
			Body: amqpMsg.GetData(),
		},
		rawAMQPMessage: amqpMsg,
	}

	if amqpMsg.Properties != nil {
		if id, ok := amqpMsg.Properties.MessageID.(string); ok {
			msg.ID = id
		}
		msg.SessionID = &amqpMsg.Properties.GroupID
		//msg.GroupSequence = &amqpMsg.Properties.GroupSequence

		if id, ok := amqpMsg.Properties.CorrelationID.(string); ok {
			msg.CorrelationID = id
		}
		msg.ContentType = amqpMsg.Properties.ContentType
		msg.Subject = amqpMsg.Properties.Subject
		msg.To = amqpMsg.Properties.To
		msg.ReplyTo = amqpMsg.Properties.ReplyTo
		msg.ReplyToSessionID = amqpMsg.Properties.ReplyToGroupID
		if amqpMsg.Header != nil {
			msg.DeliveryCount = amqpMsg.Header.DeliveryCount + 1
			msg.TimeToLive = &amqpMsg.Header.TTL
		}
	}

	if amqpMsg.ApplicationProperties != nil {
		msg.ApplicationProperties = make(map[string]interface{}, len(amqpMsg.ApplicationProperties))
		for key, value := range amqpMsg.ApplicationProperties {
			msg.ApplicationProperties[key] = value
		}
	}

	if amqpMsg.Annotations != nil {
		if lockedUntil, ok := amqpMsg.Annotations[lockedUntilAnnotation]; ok {
			t := lockedUntil.(time.Time)
			msg.LockedUntil = &t
		}

		if sequenceNumber, ok := amqpMsg.Annotations[sequenceNumberAnnotation]; ok {
			msg.SequenceNumber = to.Int64Ptr(sequenceNumber.(int64))
		}

		if partitionKey, ok := amqpMsg.Annotations[partitionKeyAnnotation]; ok {
			msg.PartitionKey = to.StringPtr(partitionKey.(string))
		}

		if enqueuedTime, ok := amqpMsg.Annotations[enqueuedTimeAnnotation]; ok {
			t := enqueuedTime.(time.Time)
			msg.EnqueuedTime = &t
		}

		if deadLetterSource, ok := amqpMsg.Annotations[deadLetterSourceAnnotation]; ok {
			msg.DeadLetterSource = to.StringPtr(deadLetterSource.(string))
		}

		if scheduledEnqueueTime, ok := amqpMsg.Annotations[scheduledEnqueuedTimeAnnotation]; ok {
			t := scheduledEnqueueTime.(time.Time)
			msg.ScheduledEnqueueTime = &t
		}

		if enqueuedSequenceNumber, ok := amqpMsg.Annotations[enqueuedSequenceNumberAnnotation]; ok {
			msg.EnqueuedSequenceNumber = to.Int64Ptr(enqueuedSequenceNumber.(int64))
		}

		if viaPartitionKey, ok := amqpMsg.Annotations[viaPartitionKeyAnnotation]; ok {
			msg.TransactionPartitionKey = to.StringPtr(viaPartitionKey.(string))
		}

		// TODO: annotation propagation is a thing. Currently these are only stored inside
		// of the underlying AMQP message, but not inside of the message itself.

		// If we didn't populate any system properties, set up the struct so we
		// can put the annotations in it
		// if msg.SystemProperties == nil {
		// 	msg.SystemProperties = new(SystemProperties)
		// }

		// Take all string-keyed annotations because the protocol reserves all
		// numeric keys for itself and there are no numeric keys defined in the
		// protocol today:
		//
		//	http://www.amqp.org/sites/amqp.org/files/amqp.pdf (section 3.2.10)
		//
		// This approach is also consistent with the behavior of .NET:
		//
		//	https://docs.microsoft.com/en-us/dotnet/api/azure.messaging.eventhubs.eventdata.systemproperties?view=azure-dotnet#Azure_Messaging_EventHubs_EventData_SystemProperties
		// msg.SystemProperties.Annotations = make(map[string]interface{})
		// for key, val := range amqpMsg.Annotations {
		// 	if s, ok := key.(string); ok {
		// 		msg.SystemProperties.Annotations[s] = val
		// 	}
		// }
	}

	if amqpMsg.DeliveryTag != nil && len(amqpMsg.DeliveryTag) > 0 {
		lockToken, err := lockTokenFromMessageTag(amqpMsg)

		if err == nil {
			msg.LockToken = *(*amqp.UUID)(lockToken)
		} else {
			tab.For(ctxForLogging).Info(fmt.Sprintf("msg.DeliveryTag could not be converted into a UUID: %s", err.Error()))
		}
	}

	if token, ok := amqpMsg.DeliveryAnnotations[lockTokenDeliveryAnnotation]; ok {
		if id, ok := token.(amqp.UUID); ok {
			msg.LockToken = [16]byte(id)
		}
	}

	msg.Format = amqpMsg.Format
	return msg
}

func lockTokenFromMessageTag(msg *amqp.Message) (*amqp.UUID, error) {
	return uuidFromLockTokenBytes(msg.DeliveryTag)
}

func uuidFromLockTokenBytes(bytes []byte) (*amqp.UUID, error) {
	if len(bytes) != 16 {
		return nil, fmt.Errorf("invalid lock token, token was not 16 bytes long")
	}

	var swapIndex = func(indexOne, indexTwo int, array *[16]byte) {
		v1 := array[indexOne]
		array[indexOne] = array[indexTwo]
		array[indexTwo] = v1
	}

	// Get lock token from the deliveryTag
	var lockTokenBytes [16]byte
	copy(lockTokenBytes[:], bytes[:16])
	// translate from .net guid byte serialisation format to amqp rfc standard
	swapIndex(0, 3, &lockTokenBytes)
	swapIndex(1, 2, &lockTokenBytes)
	swapIndex(4, 5, &lockTokenBytes)
	swapIndex(6, 7, &lockTokenBytes)
	amqpUUID := amqp.UUID(lockTokenBytes)

	return &amqpUUID, nil
}