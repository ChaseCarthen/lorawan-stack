// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
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

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { activationEvent } from '@console/lib/regexp'

import messages from '../messages'

import DescriptionList from './shared/description-list'

const DefaultPreview = React.memo(({ event }) => {
  const { identifiers, name } = event

  if (identifiers && 'device_ids' in identifiers[0]) {
    return (
      <DescriptionList>
        <DescriptionList.Byte title={messages.devAddr} data={identifiers[0].device_ids.dev_addr} />
        {activationEvent.test(name) && (
          <>
            <DescriptionList.Byte
              title={sharedMessages.joinEUI}
              data={identifiers[0].device_ids.join_eui}
            />
            <DescriptionList.Byte
              title={sharedMessages.devEUI}
              data={identifiers[0].device_ids.dev_eui}
            />
          </>
        )}
      </DescriptionList>
    )
  }

  return null
})

DefaultPreview.propTypes = {
  event: PropTypes.event.isRequired,
}

export default DefaultPreview
