import {createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import { setCredentials,setLogout } from '../../features/auth/authSlice'

const GO_API =  import.meta.env.VITE_REACT_APP_BACKEND;
const REFRESH_URL = 'auth/refresh'
// const LOGOUT_URL = 'auth/logout'

const baseQuery = fetchBaseQuery({
  baseUrl: `${GO_API}`,
  credentials: 'include',
  prepareHeaders: (headers, { getState, endpoint }) => {
    const UPLOAD_ENDPOINTS = ['uploadImage'];
    if (!UPLOAD_ENDPOINTS.includes(endpoint)) {
        const token = getState().auth.token
        if (token) {
          headers.set("Authorization", `Bearer ${token}`)
        }
    }
    return headers
  }
})


const baseQueryWithReauth = async (args, api, extraOptions) => {
  let result = await baseQuery(args, api, extraOptions)
  // if (result?.error?.originalStatus === 403){
  if (result?.error?.status === 403){
    // send refresh token to get new access token
    const refreshResult = await baseQuery(REFRESH_URL, api, extraOptions)
    if (refreshResult?.data){
      const user = api.getState().auth.user
      const roles = api.getState().auth.roles
      const token = refreshResult.data.access_token
     
      // store the new token
      api.dispatch(setCredentials({...token, user,roles}))
      // retry the original query with new access token
      result = await baseQuery(args, api, extraOptions)
    } else {
      api.dispatch(setLogout())
    }
  }
  return result
}

export const apiSlice = createApi({
  baseQuery: baseQueryWithReauth,
  tagTypes: ['User','Post'],
  endpoints: builder => ({})
})

