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

import React, { useCallback, useState } from 'react'
import { useParams } from 'react-router-dom'
import { defineMessages } from 'react-intl'
import { useSelector } from 'react-redux'

import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import PageTitle from '@ttn-lw/components/page-title'
import Form from '@ttn-lw/components/form'
import SubmitButton from '@ttn-lw/components/submit-button'
import SubmitBar from '@ttn-lw/components/submit-bar'

import { initialWifiProfile } from '@console/containers/gateway-managed-gateway/shared/utils'
import GatewayWifiProfilesFormFields from '@console/containers/gateway-managed-gateway/shared/wifi-profiles-form-fields'
import { wifiValidationSchema } from '@console/containers/gateway-managed-gateway/shared/validation-schema'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { selectFetchingEntry } from '@ttn-lw/lib/store/selectors/fetching'

const m = defineMessages({
  updateWifiProfile: 'Update WiFi profile',
})

const GatewayWifiProfilesForm = () => {
  const [error, setError] = useState(undefined)
  const { gtwId, profileId } = useParams()
  const isLoading = useSelector(state => selectFetchingEntry(state, 'GET_ACCESS_POINTS'))

  const isEdit = Boolean(profileId)

  const baseUrl = `/gateways/${gtwId}/managed-gateway/wifi-profiles`

  useBreadcrumbs(
    'gtws.single.managed-gateway.wifi-profiles.form',
    <Breadcrumb
      path={isEdit ? `${baseUrl}/edit/${profileId}` : `${baseUrl}/add`}
      content={isEdit ? m.updateWifiProfile : sharedMessages.addWifiProfile}
    />,
  )

  const handleSubmit = useCallback(values => {
    try {
      console.log(values)
    } catch (e) {
      setError(e)
    }
  }, [])

  return (
    <>
      <PageTitle title={isEdit ? m.updateWifiProfile : sharedMessages.addWifiProfile} />
      <Form
        error={error}
        onSubmit={handleSubmit}
        initialValues={initialWifiProfile}
        validationSchema={wifiValidationSchema}
      >
        <GatewayWifiProfilesFormFields isEdit={isEdit} />

        <SubmitBar>
          <Form.Submit
            component={SubmitButton}
            message={sharedMessages.saveChanges}
            disabled={isLoading}
          />
        </SubmitBar>
      </Form>
    </>
  )
}

export default GatewayWifiProfilesForm
