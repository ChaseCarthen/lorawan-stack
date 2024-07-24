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
import classnames from 'classnames'
import { useSelector } from 'react-redux'

import gatewayIcon from '@assets/misc/gateway.svg'

import Icon from '@ttn-lw/components/icon'
import Link from '@ttn-lw/components/link'

import Message from '@ttn-lw/lib/components/message'

import {
  connectionMessageMap,
  getCellularDetails,
  getConnectionType,
  getDetails,
  getEthernetDetails,
  getWifiDetails,
  NETWORK_INTERFACE_TYPES,
} from '@console/containers/gateway-managed-gateway/connection-settings/connections/utils'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

import { selectSelectedManagedGateway } from '@console/store/selectors/gateways'

import m from './messages'

import style from './connections.styl'

const ManagedGatewayConnections = ({ connectionsData }) => {
  const selectedManagedGateway = useSelector(selectSelectedManagedGateway)
  const {
    systemStatus,
    controllerConnection,
    serverConnection,
    cellularBackhaul,
    wifiBackhaul,
    ethernetBackhaul,
    updatedManagedGateway,
    isCellularConnected,
    isWifiConnected,
    isEthernetConnected,
  } = connectionsData

  const gatewayControllerConnection =
    controllerConnection?.network_interface_type ?? NETWORK_INTERFACE_TYPES.UNSPECIFIED
  const gatewayServerConnection =
    serverConnection?.network_interface_type ?? NETWORK_INTERFACE_TYPES.UNSPECIFIED

  const controllerConnectionIsSpecified =
    gatewayControllerConnection !== NETWORK_INTERFACE_TYPES.UNSPECIFIED
  const serverConnectionIsSpecified =
    gatewayServerConnection !== NETWORK_INTERFACE_TYPES.UNSPECIFIED

  const cellularDetails = isCellularConnected && getCellularDetails(cellularBackhaul)

  const managedGateway = updatedManagedGateway ?? selectedManagedGateway

  const getIsConnectedDiv = isConnected => (
    <div className="d-flex al-center gap-cs-xxs">
      <Icon
        icon={isConnected ? 'check_circle_outline' : 'highlight_remove'}
        className={isConnected ? 'c-success' : 'c-error'}
      />
      <Message content={isConnected ? sharedMessages.connected : sharedMessages.disconnected} />
    </div>
  )

  return (
    <div className={style.root}>
      <Message className="fw-bold mt-0" component="h2" content={sharedMessages.managedGateway} />
      <div className={classnames(style.top, 'd-flex j-between al-center pb-cs-xs mb-cs-xs')}>
        <div>
          <Message
            component="p"
            className="m-0"
            content={m.hardwareVersion}
            values={{ version: managedGateway.version_ids.hardware_version }}
          />
          <Message
            component="p"
            className="m-0"
            content={m.firmwareVersion}
            values={{ version: managedGateway.version_ids.firmware_version }}
          />

          {systemStatus?.cpu_temperature && (
            <div className="d-flex al-center gap-cs-xxs mb-cs-s mt-cs-xxs">
              <Icon icon="device_thermostat" />
              {systemStatus?.cpu_temperature} &deg;C
            </div>
          )}

          <Link.Anchor primary href="#">
            <Message content={m.officialDocumentation} />
          </Link.Anchor>
        </div>
        <img className={style.image} src={gatewayIcon} alt="managed-gateway" />
      </div>

      <Message className="fw-bold mt-0 mb-cs-xs" component="h3" content={m.connections} />
      {controllerConnectionIsSpecified && (
        <div className="d-flex al-center gap-cs-xxs">
          <Icon icon="cloud" />
          <Message
            content={m.connectedToGatewayController}
            values={{
              type: connectionMessageMap[
                getConnectionType(gatewayControllerConnection)
              ]?.defaultMessage.toLowerCase(),
            }}
          />
        </div>
      )}
      {serverConnectionIsSpecified && (
        <div className="d-flex al-center gap-cs-xxs">
          <Icon icon="cloud" />
          <Message
            content={m.connectedToGatewayServer}
            values={{
              type: connectionMessageMap[
                getConnectionType(gatewayServerConnection)
              ]?.defaultMessage.toLowerCase(),
            }}
          />
        </div>
      )}

      <Message className="fw-bold mb-cs-xs" component="h4" content={m.cellular} />
      <div className="d-flex al-center gap-cs-m">
        {getIsConnectedDiv(isCellularConnected)}
        {isCellularConnected && (
          <div className="d-flex al-center gap-cs-xxs">
            <Icon icon="signal_cellular_alt" />
            {cellularBackhaul?.operator}
          </div>
        )}
      </div>
      {Boolean(cellularDetails?.items?.length) && getDetails(cellularDetails)}

      <Message className="fw-bold mb-cs-xs" component="h4" content={m.wifi} />
      <div className="d-flex al-center gap-cs-m">
        {getIsConnectedDiv(isWifiConnected)}
        {isWifiConnected && (
          <div className="d-flex al-center gap-cs-xxs">
            <Icon icon="wifi" />
            {wifiBackhaul?.ssid}
          </div>
        )}
      </div>
      <div>
        <Message content={m.macAddress} values={{ address: managedGateway.wifi_mac_address }} />
      </div>
      {isWifiConnected && getDetails(getWifiDetails(wifiBackhaul))}

      <Message className="fw-bold mb-cs-xs" component="h4" content={m.ethernet} />
      <div className="d-flex al-center gap-cs-m">
        {getIsConnectedDiv(isEthernetConnected)}
        {isEthernetConnected && (
          <div className="d-flex al-center gap-cs-xxs">
            <Icon icon="router" />
            {/* TODO: Check which property is displayed here*/}
            {ethernetBackhaul.network_interface.addresses.gateway}
          </div>
        )}
      </div>
      <div>
        <Message content={m.macAddress} values={{ address: managedGateway.ethernet_mac_address }} />
      </div>
      {isEthernetConnected && getDetails(getEthernetDetails(ethernetBackhaul))}
    </div>
  )
}

ManagedGatewayConnections.propTypes = {
  connectionsData: PropTypes.shape({
    systemStatus: PropTypes.object,
    controllerConnection: PropTypes.object,
    serverConnection: PropTypes.object,
    cellularBackhaul: PropTypes.object,
    wifiBackhaul: PropTypes.object,
    ethernetBackhaul: PropTypes.object,
    updatedManagedGateway: PropTypes.object,
    isCellularConnected: PropTypes.bool,
    isWifiConnected: PropTypes.bool,
    isEthernetConnected: PropTypes.bool,
  }).isRequired,
}

export default ManagedGatewayConnections
