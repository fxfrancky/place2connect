import { apiSlice } from "../../app/api/apiSlice";

const POSTS_URL = 'posts'
const USER_POSTS_URL = 'userposts'
const LIKE_POSTS_URL = 'userlikes'
const UPLOADS_URL = 'imageUpload'
const COMMENTS_POSTS_URL = 'usercomments'

export const postsApiSlice = apiSlice.injectEndpoints({
    endpoints: builder => ({
        getPosts: builder.query({
            query: ()=> POSTS_URL,
            keepUnusedDataFor: 0,
            providesTags: ['Post'],
        }),
        getUserPosts: builder.query({
            query: (id)=> `${USER_POSTS_URL}/${id}`,
            keepUnusedDataFor: 0,
            providesTags: ['Post'],
        }),
        getPost: builder.query({
            query: (id)=> `${POSTS_URL}/${id}`,
            keepUnusedDataFor: 0,
            providesTags: ['Post'],
        }),
        addPost: builder.mutation({
          query: (post) => ({
            url : POSTS_URL,
            method: 'POST',
            body: post
          }),          
          invalidatesTags: ['Post'],
        }),
        uploadImage: builder.mutation({
          query: (picture) => ({
            url : UPLOADS_URL,
            method: 'POST',
            body: picture
          }),          
          invalidatesTags: ['Post'],
        }),
        addCommentToPost: builder.mutation({
          query: (postComment) => ({
            url : COMMENTS_POSTS_URL,
            method: 'POST',
            body: postComment
          }),          
          invalidatesTags: ['Post'],
        }),
        addLikeToPost: builder.mutation({
          query: (postLike) => ({
            url : LIKE_POSTS_URL,
            method: 'POST',
            body: postLike
          }),          
          invalidatesTags: ['Post'],
        }),
    })
})

export const {
  useGetPostsQuery,
  useGetUserPostsQuery,
  useGetPostQuery,
  useAddPostMutation,
  useUploadImageMutation,
  useAddCommentToPostMutation,
  useAddLikeToPostMutation,
} = postsApiSlice