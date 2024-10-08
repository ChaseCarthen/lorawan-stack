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

import {
  GET_CONNECTION_PROFILES_LIST_SUCCESS,
  DELETE_CONNECTION_PROFILE_SUCCESS,
  GET_ACCESS_POINTS_SUCCESS,
  GET_ACCESS_POINTS_FAILURE,
  GET_CONNECTION_PROFILE_SUCCESS,
  CREATE_CONNECTION_PROFILE_SUCCESS,
  UPDATE_CONNECTION_PROFILE_SUCCESS,
} from '@console/store/actions/connection-profiles'

const defaultState = {
  profiles: {
    wifi: [],
    ethernet: [],
  },
  wifiSelectedProfile: undefined,
  ethernetSelectedProfile: undefined,
  accessPoints: [],
}

const connectionProfiles = (state = defaultState, action) => {
  const { type, payload } = action

  switch (type) {
    case GET_CONNECTION_PROFILES_LIST_SUCCESS:
      return {
        ...state,
        profiles: {
          ...state.profiles,
          [payload.type]: payload.entities,
        },
      }
    case GET_CONNECTION_PROFILE_SUCCESS:
      return {
        ...state,
        [`${payload.type}SelectedProfile`]: payload.data,
      }
    case CREATE_CONNECTION_PROFILE_SUCCESS:
      return {
        ...state,
        profiles: {
          ...state.profiles,
          [payload.type]: payload.data.shared
            ? [...state.profiles[payload.type], payload.data]
            : state.profiles[payload.type],
        },
      }
    case DELETE_CONNECTION_PROFILE_SUCCESS:
      return {
        ...state,
        profiles: {
          ...state.profiles,
          [payload.type]: state.profiles[payload.type].filter(
            profile => profile.profile_id !== payload.profileId,
          ),
        },
      }
    case UPDATE_CONNECTION_PROFILE_SUCCESS:
      return {
        ...state,
        profiles: {
          ...state.profiles,
          [payload.type]: state.profiles[payload.type].map(profile => {
            if (payload.profileId === profile.profile_id) {
              return payload.data
            }
            return profile
          }),
        },
        [`${payload.type}SelectedProfile`]: payload.data,
      }
    case GET_ACCESS_POINTS_SUCCESS:
      return {
        ...state,
        accessPoints: payload,
      }
    case GET_ACCESS_POINTS_FAILURE:
      return {
        ...state,
        accessPoints: [],
      }
    default:
      return state
  }
}

export default connectionProfiles
