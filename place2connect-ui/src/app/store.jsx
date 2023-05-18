import { configureStore} from '@reduxjs/toolkit'
import { apiSlice } from './api/apiSlice'
import authReducer from '../features/auth/authSlice'
import { combineReducers } from '@reduxjs/toolkit';
import {
  FLUSH, PAUSE,
  PERSIST, persistReducer, PURGE,
  REGISTER, REHYDRATE
} from 'redux-persist';
import storage from 'redux-persist/lib/storage';


// combine all reducers
const reducers = combineReducers({
    // [apiSlice.reducerPath]: apiSlice.reducer,
    auth:  authReducer
})


export const store = configureStore({
  reducer: persistReducer(
    {
      key: 'root',
      storage,
      version: 1
    },
    reducers
  ),
  middleware: getDefaultMiddleware => getDefaultMiddleware({
    // serializableCheck: {
    //   ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER]
    // }
    serializableCheck:false
  }).concat(apiSlice.middleware),
  devTools: true
})


// export const store = configureStore({
//   reducer: {
//     [apiSlice.reducerPath]: apiSlice.reducer,
//     auth:  authReducer
//   },
//   middleware: getDefaultMiddleware =>
//     getDefaultMiddleware().concat(apiSlice.middleware),
//     devTools: true
// })

