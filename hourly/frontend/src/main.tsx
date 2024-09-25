import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { ChakraProvider, extendTheme, type ThemeConfig } from '@chakra-ui/react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import App from './App.tsx'
import Home from './pages/Home.tsx'
import About from './pages/About.tsx'
import Login from './pages/Login.tsx'
import Register from './pages/Register.tsx'
import UserProfile from './pages/UserProfile.tsx'
import TaskDetails from './pages/TaskDetails.tsx'
import { AuthProvider } from './AuthContext'
import './index.css'


// Create a Tanstack query client
const queryClient = new QueryClient()

// Setup ChakraUI theme config
const themeConfig : ThemeConfig = {
  initialColorMode: "dark",
  useSystemColorMode: false,
}

createRoot(document.getElementById('root')!).render(
  <StrictMode>
  <ChakraProvider theme={extendTheme({themeConfig})}>
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <BrowserRouter>
          <Routes>
            <Route path= "/" element={<App />}>
              <Route index element={<Home />} />
              <Route path="about" element={<About />} />
              <Route path="login" element={<Login />} />
              <Route path="register" element={<Register />} />
              <Route path="profile" element={<UserProfile />} />
              <Route path="tasks">
                <Route path=":taskId" element={<TaskDetails />} />
              </Route>
            </Route>          
          </Routes>
        </BrowserRouter>
      </AuthProvider>
    </QueryClientProvider>
  </ChakraProvider>
  </StrictMode>,
)
