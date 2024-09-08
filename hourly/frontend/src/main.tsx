import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { ChakraProvider, extendTheme, type ThemeConfig } from '@chakra-ui/react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import App from './App.tsx'
import Home from './pages/Home.tsx'
import About from './pages/About.tsx'
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
      <BrowserRouter>
        <Routes>
          <Route path= "/" element={<App />}>
            <Route index element={<Home />} />
            <Route path="about" element={<About />} />
          </Route>          
        </Routes>
      </BrowserRouter>
    </QueryClientProvider>
  </ChakraProvider>
  </StrictMode>,
)
