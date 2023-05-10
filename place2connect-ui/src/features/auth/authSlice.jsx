import { createSlice } from "@reduxjs/toolkit";


const initialState = {user: null, token: null, roles: null,mode:'light',posts:[]}
const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    setMode: (state) => {
      state.mode = state.mode === "light" ? "dark" : "light";
    },
    setCredentials: (state,action) => {
      const {user, access_token, roles} = action.payload
      state.user = user
      state.token = access_token
      state.roles = roles
    },
    setLogout: (state) => {
      state.user = null
      state.token = null
      state.roles = null
    },
    setFriends: (state, action) =>{
      if (state.user){
        state.user.friends = action.payload.friends
      } else {
        console.error("user friends non-existent :(")
      }
    },
    setPosts: (state, action) =>{
      state.posts = action.payload.posts;
    },
    setPost: (state, action) =>{
      const updatedPosts = state.posts.map((post) => {
        console.log(action.payload)
        if (post.id === action.payload.postID) return action.payload.post;
        return post;
      });
      state.posts = updatedPosts
    }
  },
})

export const {setMode, setCredentials, setLogout, setFriends, setPosts, setPost} = authSlice.actions

export default authSlice.reducer

export const selectCurrentUser = (state) => state.auth.user 
export const selectCurrentToken = (state) => state.auth.token
export const selectCurrentRoles = (state) => state.auth.roles
export const selectCurrentMode = (state) => state.auth.mode
export const selectCurrentPosts = (state) => state.auth.posts