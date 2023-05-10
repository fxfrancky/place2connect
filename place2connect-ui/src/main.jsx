import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { Provider } from 'react-redux'
import { store } from './app/store'

import { PersistGate } from 'redux-persist/integration/react'
import { persistStore } from "redux-persist";




ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistStore(store)}>
      <Router>       
        <Routes>
            <Route path='/*' element={<App />} />
        </Routes>
      </Router>
      </PersistGate>
    </Provider>    
  </React.StrictMode>,
)
