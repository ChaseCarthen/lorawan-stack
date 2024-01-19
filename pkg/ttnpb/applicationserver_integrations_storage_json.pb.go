// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.6.0
// - protoc             v4.23.4
// source: ttn/lorawan/v3/applicationserver_integrations_storage.proto

package ttnpb

import (
	golang "github.com/TheThingsIndustries/protoc-gen-go-json/golang"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// MarshalProtoJSON marshals the ContinuationTokenPayload message to JSON.
func (x *ContinuationTokenPayload) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Limit != nil || s.HasField("limit") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("limit")
		if x.Limit == nil {
			s.WriteNil()
		} else {
			s.WriteUint32(x.Limit.Value)
		}
	}
	if x.After != nil || s.HasField("after") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("after")
		if x.After == nil {
			s.WriteNil()
		} else {
			golang.MarshalTimestamp(s, x.After)
		}
	}
	if x.Before != nil || s.HasField("before") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("before")
		if x.Before == nil {
			s.WriteNil()
		} else {
			golang.MarshalTimestamp(s, x.Before)
		}
	}
	if x.FPort != nil || s.HasField("f_port") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("f_port")
		if x.FPort == nil {
			s.WriteNil()
		} else {
			s.WriteUint32(x.FPort.Value)
		}
	}
	if x.Order != "" || s.HasField("order") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("order")
		s.WriteString(x.Order)
	}
	if x.FieldMask != nil || s.HasField("field_mask") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field_mask")
		if x.FieldMask == nil {
			s.WriteNil()
		} else {
			golang.MarshalLegacyFieldMask(s, x.FieldMask)
		}
	}
	if x.Last != nil || s.HasField("last") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("last")
		if x.Last == nil {
			s.WriteNil()
		} else {
			golang.MarshalDuration(s, x.Last)
		}
	}
	if x.LastReceivedId != 0 || s.HasField("last_received_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("last_received_id")
		s.WriteInt64(x.LastReceivedId)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ContinuationTokenPayload to JSON.
func (x *ContinuationTokenPayload) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the ContinuationTokenPayload message from JSON.
func (x *ContinuationTokenPayload) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "limit":
			s.AddField("limit")
			if s.ReadNil() {
				x.Limit = nil
				return
			}
			v := s.ReadWrappedUint32()
			if s.Err() != nil {
				return
			}
			x.Limit = &wrapperspb.UInt32Value{Value: v}
		case "after":
			s.AddField("after")
			if s.ReadNil() {
				x.After = nil
				return
			}
			v := golang.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.After = v
		case "before":
			s.AddField("before")
			if s.ReadNil() {
				x.Before = nil
				return
			}
			v := golang.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.Before = v
		case "f_port", "fPort":
			s.AddField("f_port")
			if s.ReadNil() {
				x.FPort = nil
				return
			}
			v := s.ReadWrappedUint32()
			if s.Err() != nil {
				return
			}
			x.FPort = &wrapperspb.UInt32Value{Value: v}
		case "order":
			s.AddField("order")
			x.Order = s.ReadString()
		case "field_mask", "fieldMask":
			s.AddField("field_mask")
			if s.ReadNil() {
				x.FieldMask = nil
				return
			}
			v := golang.UnmarshalFieldMask(s)
			if s.Err() != nil {
				return
			}
			x.FieldMask = v
		case "last":
			s.AddField("last")
			if s.ReadNil() {
				x.Last = nil
				return
			}
			v := golang.UnmarshalDuration(s)
			if s.Err() != nil {
				return
			}
			x.Last = v
		case "last_received_id", "lastReceivedId":
			s.AddField("last_received_id")
			x.LastReceivedId = s.ReadInt64()
		}
	})
}

// UnmarshalJSON unmarshals the ContinuationTokenPayload from JSON.
func (x *ContinuationTokenPayload) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the GetStoredApplicationUpRequest message to JSON.
func (x *GetStoredApplicationUpRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ApplicationIds != nil || s.HasField("application_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application_ids")
		// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.ApplicationIds)
	}
	if x.EndDeviceIds != nil || s.HasField("end_device_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("end_device_ids")
		x.EndDeviceIds.MarshalProtoJSON(s.WithField("end_device_ids"))
	}
	if x.Type != "" || s.HasField("type") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("type")
		s.WriteString(x.Type)
	}
	if x.Limit != nil || s.HasField("limit") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("limit")
		if x.Limit == nil {
			s.WriteNil()
		} else {
			s.WriteUint32(x.Limit.Value)
		}
	}
	if x.After != nil || s.HasField("after") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("after")
		if x.After == nil {
			s.WriteNil()
		} else {
			golang.MarshalTimestamp(s, x.After)
		}
	}
	if x.Before != nil || s.HasField("before") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("before")
		if x.Before == nil {
			s.WriteNil()
		} else {
			golang.MarshalTimestamp(s, x.Before)
		}
	}
	if x.FPort != nil || s.HasField("f_port") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("f_port")
		if x.FPort == nil {
			s.WriteNil()
		} else {
			s.WriteUint32(x.FPort.Value)
		}
	}
	if x.Order != "" || s.HasField("order") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("order")
		s.WriteString(x.Order)
	}
	if x.FieldMask != nil || s.HasField("field_mask") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field_mask")
		if x.FieldMask == nil {
			s.WriteNil()
		} else {
			golang.MarshalLegacyFieldMask(s, x.FieldMask)
		}
	}
	if x.Last != nil || s.HasField("last") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("last")
		if x.Last == nil {
			s.WriteNil()
		} else {
			golang.MarshalDuration(s, x.Last)
		}
	}
	if x.ContinuationToken != "" || s.HasField("continuation_token") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("continuation_token")
		s.WriteString(x.ContinuationToken)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GetStoredApplicationUpRequest to JSON.
func (x *GetStoredApplicationUpRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GetStoredApplicationUpRequest message from JSON.
func (x *GetStoredApplicationUpRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application_ids", "applicationIds":
			s.AddField("application_ids")
			if s.ReadNil() {
				x.ApplicationIds = nil
				return
			}
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.ApplicationIds = &v
		case "end_device_ids", "endDeviceIds":
			if s.ReadNil() {
				x.EndDeviceIds = nil
				return
			}
			x.EndDeviceIds = &EndDeviceIdentifiers{}
			x.EndDeviceIds.UnmarshalProtoJSON(s.WithField("end_device_ids", true))
		case "type":
			s.AddField("type")
			x.Type = s.ReadString()
		case "limit":
			s.AddField("limit")
			if s.ReadNil() {
				x.Limit = nil
				return
			}
			v := s.ReadWrappedUint32()
			if s.Err() != nil {
				return
			}
			x.Limit = &wrapperspb.UInt32Value{Value: v}
		case "after":
			s.AddField("after")
			if s.ReadNil() {
				x.After = nil
				return
			}
			v := golang.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.After = v
		case "before":
			s.AddField("before")
			if s.ReadNil() {
				x.Before = nil
				return
			}
			v := golang.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.Before = v
		case "f_port", "fPort":
			s.AddField("f_port")
			if s.ReadNil() {
				x.FPort = nil
				return
			}
			v := s.ReadWrappedUint32()
			if s.Err() != nil {
				return
			}
			x.FPort = &wrapperspb.UInt32Value{Value: v}
		case "order":
			s.AddField("order")
			x.Order = s.ReadString()
		case "field_mask", "fieldMask":
			s.AddField("field_mask")
			if s.ReadNil() {
				x.FieldMask = nil
				return
			}
			v := golang.UnmarshalFieldMask(s)
			if s.Err() != nil {
				return
			}
			x.FieldMask = v
		case "last":
			s.AddField("last")
			if s.ReadNil() {
				x.Last = nil
				return
			}
			v := golang.UnmarshalDuration(s)
			if s.Err() != nil {
				return
			}
			x.Last = v
		case "continuation_token", "continuationToken":
			s.AddField("continuation_token")
			x.ContinuationToken = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the GetStoredApplicationUpRequest from JSON.
func (x *GetStoredApplicationUpRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the GetStoredApplicationUpCountRequest message to JSON.
func (x *GetStoredApplicationUpCountRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ApplicationIds != nil || s.HasField("application_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application_ids")
		// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.ApplicationIds)
	}
	if x.EndDeviceIds != nil || s.HasField("end_device_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("end_device_ids")
		x.EndDeviceIds.MarshalProtoJSON(s.WithField("end_device_ids"))
	}
	if x.Type != "" || s.HasField("type") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("type")
		s.WriteString(x.Type)
	}
	if x.After != nil || s.HasField("after") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("after")
		if x.After == nil {
			s.WriteNil()
		} else {
			golang.MarshalTimestamp(s, x.After)
		}
	}
	if x.Before != nil || s.HasField("before") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("before")
		if x.Before == nil {
			s.WriteNil()
		} else {
			golang.MarshalTimestamp(s, x.Before)
		}
	}
	if x.FPort != nil || s.HasField("f_port") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("f_port")
		if x.FPort == nil {
			s.WriteNil()
		} else {
			s.WriteUint32(x.FPort.Value)
		}
	}
	if x.Last != nil || s.HasField("last") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("last")
		if x.Last == nil {
			s.WriteNil()
		} else {
			golang.MarshalDuration(s, x.Last)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GetStoredApplicationUpCountRequest to JSON.
func (x *GetStoredApplicationUpCountRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GetStoredApplicationUpCountRequest message from JSON.
func (x *GetStoredApplicationUpCountRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application_ids", "applicationIds":
			s.AddField("application_ids")
			if s.ReadNil() {
				x.ApplicationIds = nil
				return
			}
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.ApplicationIds = &v
		case "end_device_ids", "endDeviceIds":
			if s.ReadNil() {
				x.EndDeviceIds = nil
				return
			}
			x.EndDeviceIds = &EndDeviceIdentifiers{}
			x.EndDeviceIds.UnmarshalProtoJSON(s.WithField("end_device_ids", true))
		case "type":
			s.AddField("type")
			x.Type = s.ReadString()
		case "after":
			s.AddField("after")
			if s.ReadNil() {
				x.After = nil
				return
			}
			v := golang.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.After = v
		case "before":
			s.AddField("before")
			if s.ReadNil() {
				x.Before = nil
				return
			}
			v := golang.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.Before = v
		case "f_port", "fPort":
			s.AddField("f_port")
			if s.ReadNil() {
				x.FPort = nil
				return
			}
			v := s.ReadWrappedUint32()
			if s.Err() != nil {
				return
			}
			x.FPort = &wrapperspb.UInt32Value{Value: v}
		case "last":
			s.AddField("last")
			if s.ReadNil() {
				x.Last = nil
				return
			}
			v := golang.UnmarshalDuration(s)
			if s.Err() != nil {
				return
			}
			x.Last = v
		}
	})
}

// UnmarshalJSON unmarshals the GetStoredApplicationUpCountRequest from JSON.
func (x *GetStoredApplicationUpCountRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
