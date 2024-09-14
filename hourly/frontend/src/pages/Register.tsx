import React, { useState } from 'react';
import { useAuth } from '../AuthContext';
import { useNavigate } from 'react-router-dom';
import { 
    Container,
    Heading,
    FormLabel,
    FormControl,
    Input,
    Button
 } from '@chakra-ui/react';

const Register : React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const { register } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e : React.FormEvent) => {
    e.preventDefault();
    try {
      await register(email, password);
      navigate('/'); // Redirect to dashboard upon success
    } catch (error) {
      setError('Registration failed. Try again.');
      console.log('Registration failed', error);
    }
  };

  return (
    <Container>
      <Heading>Register</Heading>
      {error && <p style={{ color: 'red' }}>{error}</p>}
        <form onSubmit={handleSubmit}>
        <FormControl
            maxWidth={'400px'}
            >
            <FormLabel>Email</FormLabel>
            <Input
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
            />
            <FormLabel>Password</FormLabel>
            <Input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
            />
            <Button type="submit">Register</Button>
        </FormControl>
      </form>
    </Container>
  );
};

export default Register;
