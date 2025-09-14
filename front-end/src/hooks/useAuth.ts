import { useState, useEffect } from 'react';

export interface User {
  id: string;
  name: string;
  email: string;
  avatar: string;
  role: 'admin' | 'agent' | 'user';
  unidade?: string;
}

export interface AuthState {
  user: User | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  error: string | null;
}

export const useAuth = () => {
  const [state, setState] = useState<AuthState>({
    user: null,
    isAuthenticated: false,
    isLoading: true,
    error: null,
  });

  // Simulate authentication check on mount
  useEffect(() => {
    const checkAuth = async () => {
      try {
        // In a real app, this would be an API call
        const storedUser = localStorage.getItem('user');
        if (storedUser) {
          const user = JSON.parse(storedUser);
          setState({
            user,
            isAuthenticated: true,
            isLoading: false,
            error: null,
          });
        } else {
          setState({
            user: null,
            isAuthenticated: false,
            isLoading: false,
            error: null,
          });
        }
      } catch (error) {
        setState({
          user: null,
          isAuthenticated: false,
          isLoading: false,
          error: 'Authentication check failed',
        });
      }
    };

    checkAuth();
  }, []);

  const login = async (credentials: { email: string; password: string }) => {
    try {
      // In a real app, this would be an API call to /api/auth/login
      // For demo purposes, we'll simulate a successful login
      const mockUser: User = {
        id: '1',
        name: 'Admin User',
        email: credentials.email,
        avatar:
          'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=100&q=80',
        role: 'admin',
        unidade: 'Matriz',
      };

      // Simulate API call delay
      await new Promise((resolve) => setTimeout(resolve, 1000));

      localStorage.setItem('user', JSON.stringify(mockUser));
      setState({
        user: mockUser,
        isAuthenticated: true,
        isLoading: false,
        error: null,
      });

      return { success: true, user: mockUser };
    } catch (error) {
      setState({
        ...state,
        isLoading: false,
        error: 'Invalid credentials',
      });
      return { success: false, error: 'Invalid credentials' };
    }
  };

  const logout = () => {
    localStorage.removeItem('user');
    setState({
      user: null,
      isAuthenticated: false,
      isLoading: false,
      error: null,
    });
  };

  return {
    ...state,
    login,
    logout,
  };
};
