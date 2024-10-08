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

import React from 'react'
import PropTypes from 'prop-types'
import { defineMessages } from 'react-intl'

import Form from '@ttn-lw/components/form'
import KeyValueMap from '@ttn-lw/components/key-value-map'
import Input from '@ttn-lw/components/input'

import sharedMessages from '@ttn-lw/lib/shared-messages'

const m = defineMessages({
  ipAddresses: 'IP addresses',
  dnsServers: 'DNS servers',
  addServerAddress: 'Add server address',
})

const NetworkInterfaceAddressesFormFields = ({ namePrefix, showOnlyDns }) => (
  <>
    {!showOnlyDns && (
      <>
        <Form.Field
          name={`${namePrefix}network_interface_addresses.ip_addresses`}
          title={m.ipAddresses}
          addMessage={m.addServerAddress}
          component={KeyValueMap}
          indexAsKey
          valuePlaceholder="192.168.100.5"
          required
        />
        <Form.Field
          title={sharedMessages.subnetMask}
          name={`${namePrefix}network_interface_addresses.subnet_mask`}
          component={Input}
          placeholder="255.255.255.0"
          required
        />
        <Form.Field
          title={sharedMessages.gateway}
          name={`${namePrefix}network_interface_addresses.gateway`}
          component={Input}
          required
          placeholder="192.168.1.0"
        />
      </>
    )}

    <Form.Field
      name={`${namePrefix}network_interface_addresses.dns_servers`}
      title={m.dnsServers}
      addMessage={m.addServerAddress}
      component={KeyValueMap}
      indexAsKey
      valuePlaceholder="8.8.8.8"
    />
  </>
)

NetworkInterfaceAddressesFormFields.propTypes = {
  namePrefix: PropTypes.string,
  showOnlyDns: PropTypes.bool,
}

NetworkInterfaceAddressesFormFields.defaultProps = {
  namePrefix: '',
  showOnlyDns: false,
}

export default NetworkInterfaceAddressesFormFields
