// src/AuthContext.js
import { createContext, useState, useContext, type ReactNode, useEffect } from 'react';
import api from './helpers/api'; // Axios instance
import axios from 'axios';

const AuthContext = createContext<AuthContextType | undefined>(undefined);

interface User {
  email: string;
}

interface AuthProviderProps {
  children: ReactNode;
}

interface AuthContextType {
  user: User | null;
  login: (email: string, password: string) => Promise<void>;
  register: (email: string, password: string) => Promise<void>;
  logout: () => void;
}


class AuthenticationError extends Error {
  constructor(response: { data: { message: string; }; }) {
    super(response.data.message);
    this.name = 'AuthenticationError';
  }
}

export const AuthProvider =  (props : AuthProviderProps ) => {
  const [user, setUser] = useState<User | null>(null);
  const [token, setToken] = useState<string | null>(null);

  const getTokenFromCookie = () => {
    const name = 'token=';
    const decodedCookie = decodeURIComponent(document.cookie);
    const ca = decodedCookie.split(';');
    for (let i = 0; i < ca.length; i++) {
      let c = ca[i];
      while (c.charAt(0) === ' ') {
        c = c.substring(1);
      }
      if (c.indexOf(name) === 0) {
        return c.substring(name.length, c.length);
      }
    }
    return '';
  };

  const isLoggedIn = () => {
    if (!document.cookie) {
      return
    }
    
    setToken(getTokenFromCookie());
    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

    /*
    try {
      const response = api.get('/auth/claims');
      setUser({ email: response.data.email });
    } catch (error) {
      console.error("Verification failed", error);
      throw error;
    }
    */
  }    

  const login = async (email: string, password: string) => {
    try {
      await api.post('/login', { email, password })
        .then(response => {
          setToken(response.data.token);
      });

      setUser({ email }); 
    } catch (error) {
      console.error("Login failed", error);
      throw new AuthenticationError({ data: { message: 'Invalid username or password' } });
    }
  };

  const register = async (email: string, password: string) => {
    try {
      await api.post('/register', { email, password });
      setUser({ email });
    } catch (error) {
      console.error("Registration failed", error);
      throw error;
    }
  };

  const logout = async () => {
    await api.post('/logout');
    setUser(null);
  };

  useEffect(() => {
    if (!token) {
      isLoggedIn();
    }
  }, [token])

  return (
    <AuthContext.Provider value={{ user, login, register, logout }}>
      {props.children}
    </AuthContext.Provider>
  );
};

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);

  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
} 
