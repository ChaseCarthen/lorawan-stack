// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.6.0
// - protoc             v4.23.4
// source: ttn/lorawan/v3/gateway_configuration.proto

package ttnpb

import (
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the GetGatewayConfigurationRequest message to JSON.
func (x *GetGatewayConfigurationRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.GatewayIds != nil || s.HasField("gateway_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("gateway_ids")
		x.GatewayIds.MarshalProtoJSON(s.WithField("gateway_ids"))
	}
	if x.Format != "" || s.HasField("format") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("format")
		s.WriteString(x.Format)
	}
	if x.Type != "" || s.HasField("type") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("type")
		s.WriteString(x.Type)
	}
	if x.Filename != "" || s.HasField("filename") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("filename")
		s.WriteString(x.Filename)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GetGatewayConfigurationRequest to JSON.
func (x *GetGatewayConfigurationRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GetGatewayConfigurationRequest message from JSON.
func (x *GetGatewayConfigurationRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "gateway_ids", "gatewayIds":
			if s.ReadNil() {
				x.GatewayIds = nil
				return
			}
			x.GatewayIds = &GatewayIdentifiers{}
			x.GatewayIds.UnmarshalProtoJSON(s.WithField("gateway_ids", true))
		case "format":
			s.AddField("format")
			x.Format = s.ReadString()
		case "type":
			s.AddField("type")
			x.Type = s.ReadString()
		case "filename":
			s.AddField("filename")
			x.Filename = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the GetGatewayConfigurationRequest from JSON.
func (x *GetGatewayConfigurationRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
