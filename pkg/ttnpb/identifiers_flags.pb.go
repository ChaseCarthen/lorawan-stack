// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v1.0.0
// - protoc              v3.9.1
// source: lorawan-stack/api/identifiers.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	pflag "github.com/spf13/pflag"
	customflags "go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/customflags"
	types "go.thethings.network/lorawan-stack/v3/pkg/types"
)

// AddSelectFlagsForApplicationIdentifiers adds flags to select fields in ApplicationIdentifiers.
func AddSelectFlagsForApplicationIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("application-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("application-id", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forApplicationIdentifiers message from select flags.
func PathsFromSelectFlagsForApplicationIdentifiers(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("application_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("application_id", prefix))
	}
	return paths, nil
}

// AddSetFlagsForApplicationIdentifiers adds flags to select fields in ApplicationIdentifiers.
func AddSetFlagsForApplicationIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("application-id", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the ApplicationIdentifiers message from flags.
func (m *ApplicationIdentifiers) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("application_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.ApplicationId = val
		paths = append(paths, flagsplugin.Prefix("application_id", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForClientIdentifiers adds flags to select fields in ClientIdentifiers.
func AddSelectFlagsForClientIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("client-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("client-id", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forClientIdentifiers message from select flags.
func PathsFromSelectFlagsForClientIdentifiers(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("client_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("client_id", prefix))
	}
	return paths, nil
}

// AddSetFlagsForClientIdentifiers adds flags to select fields in ClientIdentifiers.
func AddSetFlagsForClientIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("client-id", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the ClientIdentifiers message from flags.
func (m *ClientIdentifiers) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("client_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.ClientId = val
		paths = append(paths, flagsplugin.Prefix("client_id", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForEndDeviceIdentifiers adds flags to select fields in EndDeviceIdentifiers.
func AddSelectFlagsForEndDeviceIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("device-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("device-id", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("application-ids", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("application-ids", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application-ids", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("dev-eui", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("dev-eui", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("join-eui", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("join-eui", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("dev-addr", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("dev-addr", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forEndDeviceIdentifiers message from select flags.
func PathsFromSelectFlagsForEndDeviceIdentifiers(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("device_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("device_id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("application_ids", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("dev_eui", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("dev_eui", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("join_eui", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("join_eui", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("dev_addr", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("dev_addr", prefix))
	}
	return paths, nil
}

// AddSetFlagsForEndDeviceIdentifiers adds flags to select fields in EndDeviceIdentifiers.
func AddSetFlagsForEndDeviceIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("device-id", prefix), "", flagsplugin.WithHidden(hidden)))
	AddSetFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application-ids", prefix), hidden)
	flags.AddFlag(customflags.New8BytesFlag(flagsplugin.Prefix("dev-eui", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(customflags.New8BytesFlag(flagsplugin.Prefix("join-eui", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(customflags.New4BytesFlag(flagsplugin.Prefix("dev-addr", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the EndDeviceIdentifiers message from flags.
func (m *EndDeviceIdentifiers) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("device_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.DeviceId = val
		paths = append(paths, flagsplugin.Prefix("device_id", prefix))
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("application_ids", prefix)); changed {
		m.ApplicationIds = &ApplicationIdentifiers{}
		if setPaths, err := m.ApplicationIds.SetFromFlags(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, changed, err := types.GetEUI64(flags, flagsplugin.Prefix("dev_eui", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.DevEui = &val
		paths = append(paths, flagsplugin.Prefix("dev_eui", prefix))
	}
	if val, changed, err := types.GetEUI64(flags, flagsplugin.Prefix("join_eui", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.JoinEui = &val
		paths = append(paths, flagsplugin.Prefix("join_eui", prefix))
	}
	if val, changed, err := types.GetDevAddr(flags, flagsplugin.Prefix("dev_addr", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.DevAddr = &val
		paths = append(paths, flagsplugin.Prefix("dev_addr", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForGatewayIdentifiers adds flags to select fields in GatewayIdentifiers.
func AddSelectFlagsForGatewayIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("gateway-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("gateway-id", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("eui", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("eui", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forGatewayIdentifiers message from select flags.
func PathsFromSelectFlagsForGatewayIdentifiers(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("gateway_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("gateway_id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("eui", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("eui", prefix))
	}
	return paths, nil
}

// AddSetFlagsForGatewayIdentifiers adds flags to select fields in GatewayIdentifiers.
func AddSetFlagsForGatewayIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("gateway-id", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(customflags.New8BytesFlag(flagsplugin.Prefix("eui", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the GatewayIdentifiers message from flags.
func (m *GatewayIdentifiers) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("gateway_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.GatewayId = val
		paths = append(paths, flagsplugin.Prefix("gateway_id", prefix))
	}
	if val, changed, err := types.GetEUI64(flags, flagsplugin.Prefix("eui", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Eui = &val
		paths = append(paths, flagsplugin.Prefix("eui", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForOrganizationIdentifiers adds flags to select fields in OrganizationIdentifiers.
func AddSelectFlagsForOrganizationIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("organization-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("organization-id", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forOrganizationIdentifiers message from select flags.
func PathsFromSelectFlagsForOrganizationIdentifiers(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("organization_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("organization_id", prefix))
	}
	return paths, nil
}

// AddSetFlagsForOrganizationIdentifiers adds flags to select fields in OrganizationIdentifiers.
func AddSetFlagsForOrganizationIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("organization-id", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the OrganizationIdentifiers message from flags.
func (m *OrganizationIdentifiers) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("organization_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.OrganizationId = val
		paths = append(paths, flagsplugin.Prefix("organization_id", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForUserIdentifiers adds flags to select fields in UserIdentifiers.
func AddSelectFlagsForUserIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("user-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("user-id", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("email", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("email", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forUserIdentifiers message from select flags.
func PathsFromSelectFlagsForUserIdentifiers(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("user_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("user_id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("email", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("email", prefix))
	}
	return paths, nil
}

// AddSetFlagsForUserIdentifiers adds flags to select fields in UserIdentifiers.
func AddSetFlagsForUserIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("user-id", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("email", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the UserIdentifiers message from flags.
func (m *UserIdentifiers) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("user_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.UserId = val
		paths = append(paths, flagsplugin.Prefix("user_id", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("email", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Email = val
		paths = append(paths, flagsplugin.Prefix("email", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForOrganizationOrUserIdentifiers adds flags to select fields in OrganizationOrUserIdentifiers.
func AddSelectFlagsForOrganizationOrUserIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("ids.organization-ids", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("ids.organization-ids", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForOrganizationIdentifiers(flags, flagsplugin.Prefix("ids.organization-ids", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("ids.user-ids", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("ids.user-ids", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForUserIdentifiers(flags, flagsplugin.Prefix("ids.user-ids", prefix), hidden)
}

// SelectFromFlags outputs the fieldmask paths forOrganizationOrUserIdentifiers message from select flags.
func PathsFromSelectFlagsForOrganizationOrUserIdentifiers(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("ids.organization_ids", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("ids.organization_ids", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForOrganizationIdentifiers(flags, flagsplugin.Prefix("ids.organization_ids", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("ids.user_ids", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("ids.user_ids", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForUserIdentifiers(flags, flagsplugin.Prefix("ids.user_ids", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	return paths, nil
}

// AddSetFlagsForOrganizationOrUserIdentifiers adds flags to select fields in OrganizationOrUserIdentifiers.
func AddSetFlagsForOrganizationOrUserIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForOrganizationIdentifiers(flags, flagsplugin.Prefix("ids.organization-ids", prefix), hidden)
	AddSetFlagsForUserIdentifiers(flags, flagsplugin.Prefix("ids.user-ids", prefix), hidden)
}

// SetFromFlags sets the OrganizationOrUserIdentifiers message from flags.
func (m *OrganizationOrUserIdentifiers) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("ids.organization_ids", prefix)); changed {
		ov := &OrganizationOrUserIdentifiers_OrganizationIds{}
		ov.OrganizationIds = &OrganizationIdentifiers{}
		if setPaths, err := ov.OrganizationIds.SetFromFlags(flags, flagsplugin.Prefix("ids.organization_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
		m.Ids = ov
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("ids.user_ids", prefix)); changed {
		ov := &OrganizationOrUserIdentifiers_UserIds{}
		ov.UserIds = &UserIdentifiers{}
		if setPaths, err := ov.UserIds.SetFromFlags(flags, flagsplugin.Prefix("ids.user_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
		m.Ids = ov
	}
	return paths, nil
}

// AddSelectFlagsForEndDeviceVersionIdentifiers adds flags to select fields in EndDeviceVersionIdentifiers.
func AddSelectFlagsForEndDeviceVersionIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("brand-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("brand-id", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("model-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("model-id", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("hardware-version", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("hardware-version", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("firmware-version", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("firmware-version", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("band-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("band-id", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("vendor-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("vendor-id", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("vendor-profile-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("vendor-profile-id", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("serial-number", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("serial-number", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forEndDeviceVersionIdentifiers message from select flags.
func PathsFromSelectFlagsForEndDeviceVersionIdentifiers(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("brand_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("brand_id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("model_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("model_id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("hardware_version", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("hardware_version", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("firmware_version", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("firmware_version", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("band_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("band_id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("vendor_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("vendor_id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("vendor_profile_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("vendor_profile_id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("serial_number", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("serial_number", prefix))
	}
	return paths, nil
}

// AddSetFlagsForEndDeviceVersionIdentifiers adds flags to select fields in EndDeviceVersionIdentifiers.
func AddSetFlagsForEndDeviceVersionIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("brand-id", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("model-id", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("hardware-version", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("firmware-version", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("band-id", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("vendor-id", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("vendor-profile-id", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("serial-number", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the EndDeviceVersionIdentifiers message from flags.
func (m *EndDeviceVersionIdentifiers) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("brand_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.BrandId = val
		paths = append(paths, flagsplugin.Prefix("brand_id", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("model_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.ModelId = val
		paths = append(paths, flagsplugin.Prefix("model_id", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("hardware_version", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.HardwareVersion = val
		paths = append(paths, flagsplugin.Prefix("hardware_version", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("firmware_version", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.FirmwareVersion = val
		paths = append(paths, flagsplugin.Prefix("firmware_version", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("band_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.BandId = val
		paths = append(paths, flagsplugin.Prefix("band_id", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("vendor_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.VendorId = val
		paths = append(paths, flagsplugin.Prefix("vendor_id", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("vendor_profile_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.VendorProfileId = val
		paths = append(paths, flagsplugin.Prefix("vendor_profile_id", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("serial_number", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.SerialNumber = val
		paths = append(paths, flagsplugin.Prefix("serial_number", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForNetworkIdentifiers adds flags to select fields in NetworkIdentifiers.
func AddSelectFlagsForNetworkIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("net-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("net-id", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("tenant-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("tenant-id", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("cluster-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("cluster-id", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("cluster-address", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("cluster-address", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("tenant-address", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("tenant-address", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forNetworkIdentifiers message from select flags.
func PathsFromSelectFlagsForNetworkIdentifiers(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("net_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("net_id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("tenant_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("tenant_id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("cluster_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("cluster_id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("cluster_address", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("cluster_address", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("tenant_address", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("tenant_address", prefix))
	}
	return paths, nil
}
