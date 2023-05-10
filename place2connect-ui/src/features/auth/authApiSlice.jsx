import { apiSlice } from '../../app/api/apiSlice';

const LOGIN_URL = 'auth/login'
const LOGOUT_URL = 'auth/logout'

export const authApiSlice = apiSlice.injectEndpoints({
    endpoints: builder => ({
        login: builder.mutation({
            query: credentials => ({
                url: LOGIN_URL,
                method: 'POST',
                body: { ...credentials }
            })
        }),
        logout: builder.mutation({
            query: () => ({
                url: LOGOUT_URL,
                method: 'GET'
            })
        }),

    })
})

export const {
    useLoginMutation,
    useLogoutMutation
} = authApiSlice