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

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const GET_IS_CONFIGURATION_BASE = 'GET_IS_CONFIGURATION'
export const [
  {
    request: GET_IS_CONFIGURATION,
    success: GET_IS_CONFIGURATION_SUCCESS,
    failure: GET_IS_CONFIGURATION_FAILURE,
  },
  {
    request: getIsConfiguration,
    success: getIsConfigurationSuccess,
    failure: getIsConfigurationFailure,
  },
] = createRequestActions(GET_IS_CONFIGURATION_BASE)
