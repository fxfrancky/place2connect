import { apiSlice } from "../../app/api/apiSlice";

const REGISTER_URL = 'auth/register'
const USERS_URL = 'users'
const USER_FRIEND_URL = 'userfriends'

export const usersApiSlice = apiSlice.injectEndpoints({
    endpoints: builder => ({
        getUsers: builder.query({
            query: ()=> USERS_URL,
            keepUnusedDataFor: 0,
            providesTags: () => ["User"],
            // providesTags: ['User'],
        }),
        getUser: builder.query({
            query: (id)=> `${USERS_URL}/${id}`,
            keepUnusedDataFor: 0,
            providesTags: () => ["User"],
            // providesTags: ['User'],
        }),
        getUserFriends: builder.query({
            query: (id)=> `${USER_FRIEND_URL}/${id}`,
            keepUnusedDataFor: 0,
            // providesTags: ['User'],
            providesTags: () => ["User"],
        }),
        addUser: builder.mutation({
          query: (user) => ({
            url : REGISTER_URL,
            method: 'POST',
            body: user
          }),
          invalidatesTags: () => ["User"],
          // invalidatesTags: ['User'],
        }),
        updateUser: builder.mutation({
          query: (user) => ({
            url : `${USERS_URL}/${user.id}`,
            method: 'PUT',
            body: user
          }),
          invalidatesTags: () => ["User"],
          // invalidatesTags: ['User'],
        }),
        deleteUser: builder.mutation({
          query: ({id}) => ({
            url : `${USERS_URL}/${id}`,
            method: 'DELETE',
            body: id            
          }),
          invalidatesTags: () => ["User"],
          // invalidatesTags: ['User'],
        }),
        addFriendToUser: builder.mutation({
          query: (userfriendData) => ({
            url : USER_FRIEND_URL,
            method: 'POST',
            body: userfriendData
          }), 
          invalidatesTags: () => ["User"],         
          // invalidatesTags: ['User'],
        }),
        removeFriendFromUser: builder.mutation({
          query: (userfriendData) => ({
            url : USER_FRIEND_URL,
            method: 'DELETE',
            body: userfriendData
          }), 
          invalidatesTags: () => ["User"],         
          // invalidatesTags: ['User','Post'],
        }),
    })
})

export const {
  useGetUsersQuery,
  useGetUserQuery,
  useGetUserFriendsQuery,
  useAddUserMutation,
  useUpdateUserMutation,
  useDeleteUserMutation,
  useAddFriendToUserMutation,
  useRemoveFriendFromUserMutation,
} = usersApiSlice