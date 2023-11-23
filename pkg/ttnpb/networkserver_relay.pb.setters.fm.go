// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *RelayConfiguration) SetFields(src *RelayConfiguration, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {

		case "mode":
			if len(subs) == 0 && src == nil {
				dst.Mode = nil
				continue
			} else if len(subs) == 0 {
				dst.Mode = src.Mode
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "serving":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.Mode.(*RelayConfiguration_Serving_)
					}
					if srcValid := srcTypeOk || src == nil || src.Mode == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'serving', while different oneof is set in source")
					}
					_, dstTypeOk := dst.Mode.(*RelayConfiguration_Serving_)
					if dstValid := dstTypeOk || dst.Mode == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'serving', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *RelayConfiguration_Serving
						if srcTypeOk {
							newSrc = src.Mode.(*RelayConfiguration_Serving_).Serving
						}
						if dstTypeOk {
							newDst = dst.Mode.(*RelayConfiguration_Serving_).Serving
						} else if srcTypeOk {
							newDst = &RelayConfiguration_Serving{}
							dst.Mode = &RelayConfiguration_Serving_{Serving: newDst}
						} else {
							dst.Mode = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.Mode = src.Mode
						} else {
							dst.Mode = nil
						}
					}
				case "served":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.Mode.(*RelayConfiguration_Served_)
					}
					if srcValid := srcTypeOk || src == nil || src.Mode == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'served', while different oneof is set in source")
					}
					_, dstTypeOk := dst.Mode.(*RelayConfiguration_Served_)
					if dstValid := dstTypeOk || dst.Mode == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'served', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *RelayConfiguration_Served
						if srcTypeOk {
							newSrc = src.Mode.(*RelayConfiguration_Served_).Served
						}
						if dstTypeOk {
							newDst = dst.Mode.(*RelayConfiguration_Served_).Served
						} else if srcTypeOk {
							newDst = &RelayConfiguration_Served{}
							dst.Mode = &RelayConfiguration_Served_{Served: newDst}
						} else {
							dst.Mode = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.Mode = src.Mode
						} else {
							dst.Mode = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateRelayRequest) SetFields(src *CreateRelayRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "configuration":
			if len(subs) > 0 {
				var newDst, newSrc *RelayConfiguration
				if (src == nil || src.Configuration == nil) && dst.Configuration == nil {
					continue
				}
				if src != nil {
					newSrc = src.Configuration
				}
				if dst.Configuration != nil {
					newDst = dst.Configuration
				} else {
					newDst = &RelayConfiguration{}
					dst.Configuration = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Configuration = src.Configuration
				} else {
					dst.Configuration = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateRelayResponse) SetFields(src *CreateRelayResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "configuration":
			if len(subs) > 0 {
				var newDst, newSrc *RelayConfiguration
				if (src == nil || src.Configuration == nil) && dst.Configuration == nil {
					continue
				}
				if src != nil {
					newSrc = src.Configuration
				}
				if dst.Configuration != nil {
					newDst = dst.Configuration
				} else {
					newDst = &RelayConfiguration{}
					dst.Configuration = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Configuration = src.Configuration
				} else {
					dst.Configuration = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetRelayRequest) SetFields(src *GetRelayRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetRelayResponse) SetFields(src *GetRelayResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "configuration":
			if len(subs) > 0 {
				var newDst, newSrc *RelayConfiguration
				if (src == nil || src.Configuration == nil) && dst.Configuration == nil {
					continue
				}
				if src != nil {
					newSrc = src.Configuration
				}
				if dst.Configuration != nil {
					newDst = dst.Configuration
				} else {
					newDst = &RelayConfiguration{}
					dst.Configuration = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Configuration = src.Configuration
				} else {
					dst.Configuration = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateRelayRequest) SetFields(src *UpdateRelayRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "configuration":
			if len(subs) > 0 {
				var newDst, newSrc *RelayConfiguration
				if (src == nil || src.Configuration == nil) && dst.Configuration == nil {
					continue
				}
				if src != nil {
					newSrc = src.Configuration
				}
				if dst.Configuration != nil {
					newDst = dst.Configuration
				} else {
					newDst = &RelayConfiguration{}
					dst.Configuration = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Configuration = src.Configuration
				} else {
					dst.Configuration = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateRelayResponse) SetFields(src *UpdateRelayResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "configuration":
			if len(subs) > 0 {
				var newDst, newSrc *RelayConfiguration
				if (src == nil || src.Configuration == nil) && dst.Configuration == nil {
					continue
				}
				if src != nil {
					newSrc = src.Configuration
				}
				if dst.Configuration != nil {
					newDst = dst.Configuration
				} else {
					newDst = &RelayConfiguration{}
					dst.Configuration = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Configuration = src.Configuration
				} else {
					dst.Configuration = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DeleteRelayRequest) SetFields(src *DeleteRelayRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DeleteRelayResponse) SetFields(src *DeleteRelayResponse, paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message DeleteRelayResponse has no fields, but paths %s were specified", paths)
	}
	return nil
}

func (dst *RelayConfigurationUplinkForwardingRule) SetFields(src *RelayConfigurationUplinkForwardingRule, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "limits":
			if len(subs) > 0 {
				var newDst, newSrc *RelayUplinkForwardLimits
				if (src == nil || src.Limits == nil) && dst.Limits == nil {
					continue
				}
				if src != nil {
					newSrc = src.Limits
				}
				if dst.Limits != nil {
					newDst = dst.Limits
				} else {
					newDst = &RelayUplinkForwardLimits{}
					dst.Limits = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Limits = src.Limits
				} else {
					dst.Limits = nil
				}
			}
		case "last_w_f_cnt":
			if len(subs) > 0 {
				return fmt.Errorf("'last_w_f_cnt' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.LastWFCnt = src.LastWFCnt
			} else {
				var zero uint32
				dst.LastWFCnt = zero
			}
		case "device_id":
			if len(subs) > 0 {
				return fmt.Errorf("'device_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DeviceId = src.DeviceId
			} else {
				var zero string
				dst.DeviceId = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateRelayUplinkForwardingRuleRequest) SetFields(src *CreateRelayUplinkForwardingRuleRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "index":
			if len(subs) > 0 {
				return fmt.Errorf("'index' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Index = src.Index
			} else {
				var zero uint32
				dst.Index = zero
			}
		case "uplink_forwarding_rule":
			if len(subs) > 0 {
				var newDst, newSrc *RelayConfigurationUplinkForwardingRule
				if (src == nil || src.UplinkForwardingRule == nil) && dst.UplinkForwardingRule == nil {
					continue
				}
				if src != nil {
					newSrc = src.UplinkForwardingRule
				}
				if dst.UplinkForwardingRule != nil {
					newDst = dst.UplinkForwardingRule
				} else {
					newDst = &RelayConfigurationUplinkForwardingRule{}
					dst.UplinkForwardingRule = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UplinkForwardingRule = src.UplinkForwardingRule
				} else {
					dst.UplinkForwardingRule = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateRelayUplinkForwardingRuleResponse) SetFields(src *CreateRelayUplinkForwardingRuleResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "uplink_forwarding_rule":
			if len(subs) > 0 {
				var newDst, newSrc *RelayConfigurationUplinkForwardingRule
				if (src == nil || src.UplinkForwardingRule == nil) && dst.UplinkForwardingRule == nil {
					continue
				}
				if src != nil {
					newSrc = src.UplinkForwardingRule
				}
				if dst.UplinkForwardingRule != nil {
					newDst = dst.UplinkForwardingRule
				} else {
					newDst = &RelayConfigurationUplinkForwardingRule{}
					dst.UplinkForwardingRule = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UplinkForwardingRule = src.UplinkForwardingRule
				} else {
					dst.UplinkForwardingRule = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetRelayUplinkForwardingRuleRequest) SetFields(src *GetRelayUplinkForwardingRuleRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "index":
			if len(subs) > 0 {
				return fmt.Errorf("'index' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Index = src.Index
			} else {
				var zero uint32
				dst.Index = zero
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetRelayUplinkForwardingRuleResponse) SetFields(src *GetRelayUplinkForwardingRuleResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "uplink_forwarding_rule":
			if len(subs) > 0 {
				var newDst, newSrc *RelayConfigurationUplinkForwardingRule
				if (src == nil || src.UplinkForwardingRule == nil) && dst.UplinkForwardingRule == nil {
					continue
				}
				if src != nil {
					newSrc = src.UplinkForwardingRule
				}
				if dst.UplinkForwardingRule != nil {
					newDst = dst.UplinkForwardingRule
				} else {
					newDst = &RelayConfigurationUplinkForwardingRule{}
					dst.UplinkForwardingRule = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UplinkForwardingRule = src.UplinkForwardingRule
				} else {
					dst.UplinkForwardingRule = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListRelayUplinkForwardingRulesRequest) SetFields(src *ListRelayUplinkForwardingRulesRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListRelayUplinkForwardingRulesResponse) SetFields(src *ListRelayUplinkForwardingRulesResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "uplink_forwarding_rules":
			if len(subs) > 0 {
				return fmt.Errorf("'uplink_forwarding_rules' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UplinkForwardingRules = src.UplinkForwardingRules
			} else {
				dst.UplinkForwardingRules = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateRelayUplinkForwardingRuleRequest) SetFields(src *UpdateRelayUplinkForwardingRuleRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "index":
			if len(subs) > 0 {
				return fmt.Errorf("'index' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Index = src.Index
			} else {
				var zero uint32
				dst.Index = zero
			}
		case "uplink_forwarding_rule":
			if len(subs) > 0 {
				var newDst, newSrc *RelayConfigurationUplinkForwardingRule
				if (src == nil || src.UplinkForwardingRule == nil) && dst.UplinkForwardingRule == nil {
					continue
				}
				if src != nil {
					newSrc = src.UplinkForwardingRule
				}
				if dst.UplinkForwardingRule != nil {
					newDst = dst.UplinkForwardingRule
				} else {
					newDst = &RelayConfigurationUplinkForwardingRule{}
					dst.UplinkForwardingRule = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UplinkForwardingRule = src.UplinkForwardingRule
				} else {
					dst.UplinkForwardingRule = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateRelayUplinkForwardingRuleResponse) SetFields(src *UpdateRelayUplinkForwardingRuleResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "uplink_forwarding_rule":
			if len(subs) > 0 {
				var newDst, newSrc *RelayConfigurationUplinkForwardingRule
				if (src == nil || src.UplinkForwardingRule == nil) && dst.UplinkForwardingRule == nil {
					continue
				}
				if src != nil {
					newSrc = src.UplinkForwardingRule
				}
				if dst.UplinkForwardingRule != nil {
					newDst = dst.UplinkForwardingRule
				} else {
					newDst = &RelayConfigurationUplinkForwardingRule{}
					dst.UplinkForwardingRule = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UplinkForwardingRule = src.UplinkForwardingRule
				} else {
					dst.UplinkForwardingRule = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DeleteRelayUplinkForwardingRuleRequest) SetFields(src *DeleteRelayUplinkForwardingRuleRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "index":
			if len(subs) > 0 {
				return fmt.Errorf("'index' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Index = src.Index
			} else {
				var zero uint32
				dst.Index = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DeleteRelayUplinkForwardingRuleResponse) SetFields(src *DeleteRelayUplinkForwardingRuleResponse, paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message DeleteRelayUplinkForwardingRuleResponse has no fields, but paths %s were specified", paths)
	}
	return nil
}

func (dst *RelayConfiguration_Serving) SetFields(src *RelayConfiguration_Serving, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "second_channel":
			if len(subs) > 0 {
				var newDst, newSrc *RelaySecondChannel
				if (src == nil || src.SecondChannel == nil) && dst.SecondChannel == nil {
					continue
				}
				if src != nil {
					newSrc = src.SecondChannel
				}
				if dst.SecondChannel != nil {
					newDst = dst.SecondChannel
				} else {
					newDst = &RelaySecondChannel{}
					dst.SecondChannel = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.SecondChannel = src.SecondChannel
				} else {
					dst.SecondChannel = nil
				}
			}
		case "default_channel_index":
			if len(subs) > 0 {
				return fmt.Errorf("'default_channel_index' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DefaultChannelIndex = src.DefaultChannelIndex
			} else {
				var zero uint32
				dst.DefaultChannelIndex = zero
			}
		case "cad_periodicity":
			if len(subs) > 0 {
				return fmt.Errorf("'cad_periodicity' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CadPeriodicity = src.CadPeriodicity
			} else {
				dst.CadPeriodicity = 0
			}
		case "limits":
			if len(subs) > 0 {
				var newDst, newSrc *ServingRelayForwardingLimits
				if (src == nil || src.Limits == nil) && dst.Limits == nil {
					continue
				}
				if src != nil {
					newSrc = src.Limits
				}
				if dst.Limits != nil {
					newDst = dst.Limits
				} else {
					newDst = &ServingRelayForwardingLimits{}
					dst.Limits = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Limits = src.Limits
				} else {
					dst.Limits = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *RelayConfiguration_Served) SetFields(src *RelayConfiguration_Served, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "backoff":
			if len(subs) > 0 {
				return fmt.Errorf("'backoff' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Backoff = src.Backoff
			} else {
				var zero uint32
				dst.Backoff = zero
			}
		case "second_channel":
			if len(subs) > 0 {
				var newDst, newSrc *RelaySecondChannel
				if (src == nil || src.SecondChannel == nil) && dst.SecondChannel == nil {
					continue
				}
				if src != nil {
					newSrc = src.SecondChannel
				}
				if dst.SecondChannel != nil {
					newDst = dst.SecondChannel
				} else {
					newDst = &RelaySecondChannel{}
					dst.SecondChannel = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.SecondChannel = src.SecondChannel
				} else {
					dst.SecondChannel = nil
				}
			}
		case "serving_device_id":
			if len(subs) > 0 {
				return fmt.Errorf("'serving_device_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ServingDeviceId = src.ServingDeviceId
			} else {
				var zero string
				dst.ServingDeviceId = zero
			}

		case "mode":
			if len(subs) == 0 && src == nil {
				dst.Mode = nil
				continue
			} else if len(subs) == 0 {
				dst.Mode = src.Mode
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "always":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.Mode.(*RelayConfiguration_Served_Always)
					}
					if srcValid := srcTypeOk || src == nil || src.Mode == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'always', while different oneof is set in source")
					}
					_, dstTypeOk := dst.Mode.(*RelayConfiguration_Served_Always)
					if dstValid := dstTypeOk || dst.Mode == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'always', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *RelayEndDeviceAlwaysMode
						if srcTypeOk {
							newSrc = src.Mode.(*RelayConfiguration_Served_Always).Always
						}
						if dstTypeOk {
							newDst = dst.Mode.(*RelayConfiguration_Served_Always).Always
						} else if srcTypeOk {
							newDst = &RelayEndDeviceAlwaysMode{}
							dst.Mode = &RelayConfiguration_Served_Always{Always: newDst}
						} else {
							dst.Mode = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.Mode = src.Mode
						} else {
							dst.Mode = nil
						}
					}
				case "dynamic":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.Mode.(*RelayConfiguration_Served_Dynamic)
					}
					if srcValid := srcTypeOk || src == nil || src.Mode == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'dynamic', while different oneof is set in source")
					}
					_, dstTypeOk := dst.Mode.(*RelayConfiguration_Served_Dynamic)
					if dstValid := dstTypeOk || dst.Mode == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'dynamic', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *RelayEndDeviceDynamicMode
						if srcTypeOk {
							newSrc = src.Mode.(*RelayConfiguration_Served_Dynamic).Dynamic
						}
						if dstTypeOk {
							newDst = dst.Mode.(*RelayConfiguration_Served_Dynamic).Dynamic
						} else if srcTypeOk {
							newDst = &RelayEndDeviceDynamicMode{}
							dst.Mode = &RelayConfiguration_Served_Dynamic{Dynamic: newDst}
						} else {
							dst.Mode = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.Mode = src.Mode
						} else {
							dst.Mode = nil
						}
					}
				case "end_device_controlled":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.Mode.(*RelayConfiguration_Served_EndDeviceControlled)
					}
					if srcValid := srcTypeOk || src == nil || src.Mode == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'end_device_controlled', while different oneof is set in source")
					}
					_, dstTypeOk := dst.Mode.(*RelayConfiguration_Served_EndDeviceControlled)
					if dstValid := dstTypeOk || dst.Mode == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'end_device_controlled', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *RelayEndDeviceControlledMode
						if srcTypeOk {
							newSrc = src.Mode.(*RelayConfiguration_Served_EndDeviceControlled).EndDeviceControlled
						}
						if dstTypeOk {
							newDst = dst.Mode.(*RelayConfiguration_Served_EndDeviceControlled).EndDeviceControlled
						} else if srcTypeOk {
							newDst = &RelayEndDeviceControlledMode{}
							dst.Mode = &RelayConfiguration_Served_EndDeviceControlled{EndDeviceControlled: newDst}
						} else {
							dst.Mode = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.Mode = src.Mode
						} else {
							dst.Mode = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
