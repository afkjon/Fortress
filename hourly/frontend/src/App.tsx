import React from 'react'
import './App.css'

import Navbar from './components/Navbar'
import Footer from './components/Footer'

import { Outlet } from 'react-router-dom'
import { Container } from '@chakra-ui/react'

const App: React.FC = () => {

  return (
    <div>
      <Navbar />

      <Container maxW='100%' my='30px' backgroundColor={'inherit'}>
        <Outlet />
      </Container>

      <Footer />
    </div>
  )
}

export default App
