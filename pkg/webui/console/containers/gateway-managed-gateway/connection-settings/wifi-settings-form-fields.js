// Copyright © 2024 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React, { useCallback, useMemo } from 'react'
import { defineMessages } from 'react-intl'
import { useDispatch, useSelector } from 'react-redux'
import { isEqual, omit } from 'lodash'

import Form, { useFormContext } from '@ttn-lw/components/form'
import Select from '@ttn-lw/components/select'
import Icon from '@ttn-lw/components/icon'
import Notification from '@ttn-lw/components/notification'
import Button from '@ttn-lw/components/button'
import Link from '@ttn-lw/components/link'

import Message from '@ttn-lw/lib/components/message'
import RequireRequest from '@ttn-lw/lib/components/require-request'

import GatewayWifiProfilesFormFields from '@console/containers/gateway-managed-gateway/shared/wifi-profiles-form-fields'
import ShowProfilesSelect from '@console/containers/gateway-managed-gateway/shared/show-profiles-select'
import { CONNECTION_TYPES } from '@console/containers/gateway-managed-gateway/shared/utils'

import tooltipIds from '@ttn-lw/lib/constants/tooltip-ids'
import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'
import PropTypes from '@ttn-lw/lib/prop-types'

import { getConnectionProfilesList } from '@console/store/actions/connection-profiles'

import { selectConnectionProfilesByType } from '@console/store/selectors/connection-profiles'

const m = defineMessages({
  settingsProfile: 'Settings profile',
  profileDescription: 'Connection settings profiles can be shared within the same organization',
  wifiConnection: 'WiFi connection',
  selectAProfile: 'Select a profile',
  connected: 'The gateway WiFi successfully connected using this profile',
  connectedCollaborator: "The gateway WiFi successfully connected using a collaborator's profile",
  unableToConnect: 'The gateway WiFi is currently unable to connect using this profile',
  unableToConnectCollaborator:
    "The gateway WiFi is currently unable to connect using  a collaborator's profile",
  saveToConnect: 'Please click "Save changes" to start using this WiFi profile for the gateway',
  createNewSharedProfile: 'Create a new shared profile',
  setAConfigForThisGateway: 'Set a config for this gateway only',
  notificationContent:
    'This gateway already has a WiFi profile set by another collaborator. If wished, you can override this profile below.',
  overrideProfile: 'Override this profile',
  editProfile: 'Edit this profile',
})

const WifiSettingsFormFields = ({ initialValues, isWifiConnected }) => {
  const { values, setValues } = useFormContext()
  const dispatch = useDispatch()

  const profiles = useSelector(state =>
    selectConnectionProfilesByType(state, CONNECTION_TYPES.WIFI),
  )

  const hasChanged = useMemo(
    () =>
      !isEqual(
        omit(values.wifi_profile, ['_profile_of']),
        omit(initialValues.wifi_profile, ['_profile_of']),
      ),
    [initialValues, values],
  )

  const profileOptions = [
    ...profiles.map(p => ({
      value: p.profile_id,
      label: p.profile_name,
    })),
    { value: 'shared', label: m.createNewSharedProfile.defaultMessage },
    { value: 'non-shared', label: m.setAConfigForThisGateway.defaultMessage },
  ]

  const connectionStatus = useMemo(() => {
    if (!values.wifi_profile.profile_id) return null
    if (hasChanged) {
      return { message: m.saveToConnect, icon: 'more_horiz' }
    }
    if (isWifiConnected) {
      if (values.wifi_profile._override) {
        return { message: m.connectedCollaborator, icon: 'valid' }
      }
      return { message: m.connected, icon: 'valid' }
    }
    if (!isWifiConnected) {
      if (values.wifi_profile._override) {
        return { message: m.unableToConnectCollaborator, icon: 'cancel' }
      }
      return { message: m.unableToConnect, icon: 'cancel' }
    }

    return null
  }, [hasChanged, isWifiConnected, values.wifi_profile._override, values.wifi_profile.profile_id])

  const handleChangeProfile = useCallback(
    async value => {
      setValues(values => ({
        ...values,
        wifi_profile: {
          ...values.wifi_profile,
          profile_id: '',
        },
      }))

      await dispatch(
        attachPromise(
          getConnectionProfilesList({
            entityId: value,
            type: CONNECTION_TYPES.WIFI,
          }),
        ),
      )
    },
    [dispatch, setValues],
  )

  const handleOverrideProfile = useCallback(
    () =>
      setValues(oldValues => ({
        ...oldValues,
        wifi_profile: {
          ...oldValues.wifi_profile,
          _override: false,
          profile_id: '',
        },
      })),
    [setValues],
  )

  return (
    <>
      <Message component="h3" content={m.wifiConnection} />
      {!values.wifi_profile._override ? (
        <>
          <div className="d-flex al-center gap-cs-m">
            <ShowProfilesSelect name={`wifi_profile._profile_of`} onChange={handleChangeProfile} />
            {Boolean(values.wifi_profile._profile_of) && (
              <RequireRequest
                requestAction={getConnectionProfilesList({
                  entityId: values.wifi_profile._profile_of,
                  type: CONNECTION_TYPES.WIFI,
                })}
              >
                <Form.Field
                  name={`wifi_profile.profile_id`}
                  title={m.settingsProfile}
                  component={Select}
                  options={profileOptions}
                  tooltipId={tooltipIds.GATEWAY_SHOW_PROFILES}
                  placeholder={m.selectAProfile}
                />
              </RequireRequest>
            )}
          </div>
          <Message
            component="div"
            content={m.profileDescription}
            className="tc-subtle-gray mb-cs-m"
          />
          {values.wifi_profile.profile_id.includes('shared') && (
            <GatewayWifiProfilesFormFields namePrefix={`wifi_profile.`} />
          )}
        </>
      ) : (
        <>
          <Notification info small content={m.notificationContent} />
          <Button
            type="button"
            className="mb-cs-m"
            message={m.overrideProfile}
            onClick={handleOverrideProfile}
          />
        </>
      )}

      {connectionStatus !== null && (
        <div className="d-flex al-center gap-cs-xs">
          <Icon icon={connectionStatus.icon} />
          <Message content={connectionStatus.message} />
        </div>
      )}
      {Boolean(values.wifi_profile.profile_id) &&
        !values.wifi_profile.profile_id.includes('shared') && (
          <Link.Anchor
            primary
            href={`wifi-profiles/edit/${values.wifi_profile.profile_id}?profileOf=${values.wifi_profile._profile_of}`}
          >
            <Message content={m.editProfile} />
          </Link.Anchor>
        )}
    </>
  )
}

WifiSettingsFormFields.propTypes = {
  initialValues: PropTypes.shape({
    wifi_profile: PropTypes.shape({
      _override: PropTypes.bool,
      profile_id: PropTypes.string,
      _profile_of: PropTypes.string,
    }),
  }).isRequired,
  isWifiConnected: PropTypes.bool,
}

WifiSettingsFormFields.defaultProps = {
  isWifiConnected: false,
}

export default WifiSettingsFormFields
