// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.6.0
// - protoc             v4.23.4
// source: ttn/lorawan/v3/enums.proto

package ttnpb

import (
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the DownlinkPathConstraint to JSON.
func (x DownlinkPathConstraint) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	s.WriteEnumString(int32(x), DownlinkPathConstraint_name)
}

// MarshalText marshals the DownlinkPathConstraint to text.
func (x DownlinkPathConstraint) MarshalText() ([]byte, error) {
	return []byte(jsonplugin.GetEnumString(int32(x), DownlinkPathConstraint_name)), nil
}

// MarshalJSON marshals the DownlinkPathConstraint to JSON.
func (x DownlinkPathConstraint) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// DownlinkPathConstraint_customvalue contains custom string values that extend DownlinkPathConstraint_value.
var DownlinkPathConstraint_customvalue = map[string]int32{
	"NONE":         0,
	"PREFER_OTHER": 1,
	"NEVER":        2,
}

// UnmarshalProtoJSON unmarshals the DownlinkPathConstraint from JSON.
func (x *DownlinkPathConstraint) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	v := s.ReadEnum(DownlinkPathConstraint_value, DownlinkPathConstraint_customvalue)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read DownlinkPathConstraint enum: %v", err)
		return
	}
	*x = DownlinkPathConstraint(v)
}

// UnmarshalText unmarshals the DownlinkPathConstraint from text.
func (x *DownlinkPathConstraint) UnmarshalText(b []byte) error {
	i, err := jsonplugin.ParseEnumString(string(b), DownlinkPathConstraint_customvalue, DownlinkPathConstraint_value)
	if err != nil {
		return err
	}
	*x = DownlinkPathConstraint(i)
	return nil
}

// UnmarshalJSON unmarshals the DownlinkPathConstraint from JSON.
func (x *DownlinkPathConstraint) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the State to JSON.
func (x State) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	s.WriteEnumString(int32(x), State_name)
}

// MarshalText marshals the State to text.
func (x State) MarshalText() ([]byte, error) {
	return []byte(jsonplugin.GetEnumString(int32(x), State_name)), nil
}

// MarshalJSON marshals the State to JSON.
func (x State) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// State_customvalue contains custom string values that extend State_value.
var State_customvalue = map[string]int32{
	"REQUESTED": 0,
	"APPROVED":  1,
	"REJECTED":  2,
	"FLAGGED":   3,
	"SUSPENDED": 4,
}

// UnmarshalProtoJSON unmarshals the State from JSON.
func (x *State) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	v := s.ReadEnum(State_value, State_customvalue)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read State enum: %v", err)
		return
	}
	*x = State(v)
}

// UnmarshalText unmarshals the State from text.
func (x *State) UnmarshalText(b []byte) error {
	i, err := jsonplugin.ParseEnumString(string(b), State_customvalue, State_value)
	if err != nil {
		return err
	}
	*x = State(i)
	return nil
}

// UnmarshalJSON unmarshals the State from JSON.
func (x *State) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
