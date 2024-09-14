// src/AuthContext.js
import { createContext, useState, useContext, type ReactNode, useEffect } from 'react';
import api from './helpers/api'; // Axios instance

const AuthContext = createContext<AuthContextType | undefined>(undefined);

interface User {
  ID: number;
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

  const checkUser = async () => {
    try {
      const res = await api.get<User>('/auth/user', {
        withCredentials: true
      })
      
      // todo: fix this
      if (res.status === 200) {
        const u : User = res.data
        setUser(u);
      }
    } catch (error) {
      setUser(null);
      console.error("Error fetching user: ", error);
    }
  }    

  const login = async (email: string, password: string) => {
    try {
      const res = await api.post('/login', { email, password }, {
          headers: { 'Content-Type': 'application/json' },
          withCredentials: true
        }
      )
      setUser(res.data.user);
    } catch (error) {
      console.error("Login failed", error);
      throw new AuthenticationError({ data: { message: 'Invalid username or password' } });
    }
  };

  const register = async (email: string, password: string) => {
    try {
      const res = await api.post('/register', { email, password },
        { withCredentials: true }
      )
      if (res.status === 200) {
        setUser({ ID: res.data.user.id, email: res.data.user.email });
      }
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
    checkUser();
  }, [])

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
