// Code generated by protoc-gen-gogo.
// source: envelope.proto
// DO NOT EDIT!

/*
Package events is a generated protocol buffer package.

It is generated from these files:
	envelope.proto
	error.proto
	heartbeat.proto
	http.proto
	log.proto
	metric.proto
	uuid.proto

It has these top-level messages:
	Envelope
*/
package events

import proto "github.com/gogo/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// / Type of the wrapped event.
type Envelope_EventType int32

const (
	Envelope_Heartbeat       Envelope_EventType = 1
	Envelope_HttpStart       Envelope_EventType = 2
	Envelope_HttpStop        Envelope_EventType = 3
	Envelope_HttpStartStop   Envelope_EventType = 4
	Envelope_LogMessage      Envelope_EventType = 5
	Envelope_ValueMetric     Envelope_EventType = 6
	Envelope_CounterEvent    Envelope_EventType = 7
	Envelope_Error           Envelope_EventType = 8
	Envelope_ContainerMetric Envelope_EventType = 9
)

var Envelope_EventType_name = map[int32]string{
	1: "Heartbeat",
	2: "HttpStart",
	3: "HttpStop",
	4: "HttpStartStop",
	5: "LogMessage",
	6: "ValueMetric",
	7: "CounterEvent",
	8: "Error",
	9: "ContainerMetric",
}
var Envelope_EventType_value = map[string]int32{
	"Heartbeat":       1,
	"HttpStart":       2,
	"HttpStop":        3,
	"HttpStartStop":   4,
	"LogMessage":      5,
	"ValueMetric":     6,
	"CounterEvent":    7,
	"Error":           8,
	"ContainerMetric": 9,
}

func (x Envelope_EventType) Enum() *Envelope_EventType {
	p := new(Envelope_EventType)
	*p = x
	return p
}
func (x Envelope_EventType) String() string {
	return proto.EnumName(Envelope_EventType_name, int32(x))
}
func (x *Envelope_EventType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Envelope_EventType_value, data, "Envelope_EventType")
	if err != nil {
		return err
	}
	*x = Envelope_EventType(value)
	return nil
}

// / Envelope wraps an Event and adds metadata.
type Envelope struct {
	Origin           *string             `protobuf:"bytes,1,req,name=origin" json:"origin,omitempty"`
	EventType        *Envelope_EventType `protobuf:"varint,2,req,name=eventType,enum=events.Envelope_EventType" json:"eventType,omitempty"`
	Timestamp        *int64              `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	Heartbeat        *Heartbeat          `protobuf:"bytes,3,opt,name=heartbeat" json:"heartbeat,omitempty"`
	HttpStart        *HttpStart          `protobuf:"bytes,4,opt,name=httpStart" json:"httpStart,omitempty"`
	HttpStop         *HttpStop           `protobuf:"bytes,5,opt,name=httpStop" json:"httpStop,omitempty"`
	HttpStartStop    *HttpStartStop      `protobuf:"bytes,7,opt,name=httpStartStop" json:"httpStartStop,omitempty"`
	LogMessage       *LogMessage         `protobuf:"bytes,8,opt,name=logMessage" json:"logMessage,omitempty"`
	ValueMetric      *ValueMetric        `protobuf:"bytes,9,opt,name=valueMetric" json:"valueMetric,omitempty"`
	CounterEvent     *CounterEvent       `protobuf:"bytes,10,opt,name=counterEvent" json:"counterEvent,omitempty"`
	Error            *Error              `protobuf:"bytes,11,opt,name=error" json:"error,omitempty"`
	ContainerMetric  *ContainerMetric    `protobuf:"bytes,12,opt,name=containerMetric" json:"containerMetric,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}

func (m *Envelope) GetOrigin() string {
	if m != nil && m.Origin != nil {
		return *m.Origin
	}
	return ""
}

func (m *Envelope) GetEventType() Envelope_EventType {
	if m != nil && m.EventType != nil {
		return *m.EventType
	}
	return Envelope_Heartbeat
}

func (m *Envelope) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Envelope) GetHeartbeat() *Heartbeat {
	if m != nil {
		return m.Heartbeat
	}
	return nil
}

func (m *Envelope) GetHttpStart() *HttpStart {
	if m != nil {
		return m.HttpStart
	}
	return nil
}

func (m *Envelope) GetHttpStop() *HttpStop {
	if m != nil {
		return m.HttpStop
	}
	return nil
}

func (m *Envelope) GetHttpStartStop() *HttpStartStop {
	if m != nil {
		return m.HttpStartStop
	}
	return nil
}

func (m *Envelope) GetLogMessage() *LogMessage {
	if m != nil {
		return m.LogMessage
	}
	return nil
}

func (m *Envelope) GetValueMetric() *ValueMetric {
	if m != nil {
		return m.ValueMetric
	}
	return nil
}

func (m *Envelope) GetCounterEvent() *CounterEvent {
	if m != nil {
		return m.CounterEvent
	}
	return nil
}

func (m *Envelope) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Envelope) GetContainerMetric() *ContainerMetric {
	if m != nil {
		return m.ContainerMetric
	}
	return nil
}

func init() {
	proto.RegisterEnum("events.Envelope_EventType", Envelope_EventType_name, Envelope_EventType_value)
}
