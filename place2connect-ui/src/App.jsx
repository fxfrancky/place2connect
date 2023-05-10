import { ROLES } from "./constants/RolesContants"
import {Routes, Route} from 'react-router-dom'

// import 'react-toastify/dist/ReactToastify.css'
// import Header from './components/Header'
// import Footer from './components/Footer'
import Layout from './components/Layout'
import LoginPage from "./scenes/loginPage"
import HomePage from "./scenes/homePage"
import ProfilePage from "./scenes/profilePage"
import RequireAuth from "./components/RequireAuth"

// import HomePage from './pages/HomePage'
// import LoginPage from './pages/LoginPage'
// import RegisterPage from './pages/RegisterPage'
// import UsersPage from "./pages/UsersPage"
// import DashboardPage from "./pages/DashboardPage"

import Missing from "./pages/global_pages/Missing"
import Unauthorized from "./pages/global_pages/Unauthorized"
import { useMemo } from 'react'
import { useSelector } from 'react-redux'
import { CssBaseline, ThemeProvider } from '@mui/material'
import { createTheme } from '@mui/material/styles'
import { themeSettings } from './theme'
import { selectCurrentMode } from "./features/auth/authSlice";

function App() {
    // const mode = useSelector((state) => state.mode);
    // const mode = useSelector((state) => state.mode);
    const mode = useSelector(selectCurrentMode)
    const theme = useMemo(() => createTheme(themeSettings(mode)), [mode]);

 return (
    <div className="app">
      {/* <Header /> */}
        <ThemeProvider theme={theme}>
            <CssBaseline />
            <Routes>
                <Route path="/" element={<Layout />}>
                    {/* public routes */}
                    <Route index element={<LoginPage />} />
                    {/* <Route path='/home' element={<HomePage />}  />
                    <Route path='/profile:/:userId' element={<ProfilePage />}  /> */}
                    {/* <Route path='/login' element={<LoginPage />}  />
                    <Route path='/register' element={<RegisterPage />}  /> */}
                    
                    {/* Not Authorized */}
                    <Route path="/unauthorized" element={<Unauthorized />} />

                    {/* protected routes */}
                    <Route element={<RequireAuth allowedRoles={[ROLES.Customer]}/>}> 
                        <Route path='/home' element={<HomePage />}  />
                        <Route path='/profile/:userID' element={<ProfilePage />}  />
                       {/* <Route path='/userslist' element={<UsersPage />}  /> */}
                    </Route>
                    {/* <Route element={<RequireAuth allowedRoles={[ROLES.Admin]}/>}>
                       <Route path="/dashboard" element={<DashboardPage />} />
                    </Route> */}
                    {/* catch all 'Page Not Found' */}
                    <Route path="*" element={<Missing />} />
                </Route>
            </Routes> 
        </ThemeProvider>       
      {/* <Footer />   */}
  </div>
  )
}

export default App
