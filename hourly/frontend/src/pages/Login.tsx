import React, { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../AuthContext';
import {
    Container,
    Button,
    Input,
    FormLabel,
    FormControl,
} from '@chakra-ui/react';

const Login: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const { user, login } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
        await login(email, password);
        if (user) {
            navigate('/');
        }
    } catch (error) {
        setError('Invalid username or password');
        console.error('Login failed', error);
    }
  };

  // Redirect to home if user has successfully logged in 
  // Or is already logged in
  useEffect(() => {
    if (user) {
      navigate('/');
    }
  })

  return (
    <Container>
      <h1>Login</h1>
      { error && <p color={'red'}>{error}</p> }
      <form onSubmit={handleSubmit}>
        <FormControl>
            <FormLabel>Email</FormLabel>
            <Input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            />
            <FormLabel>Password</FormLabel>
            <Input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            />
            <Button type="submit">Login</Button>
        </FormControl>
      </form>
    </Container>
  )
}


export default Login