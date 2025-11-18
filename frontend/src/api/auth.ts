import { apiClient } from './client';

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
  username: string;
}

export interface User {
  id: string;
  email: string;
  username: string;
  risk_tolerance?: string;
  learning_style?: string;
  investment_experience?: string;
  trading_pattern_memo?: string;
  created_at: string;
  updated_at: string;
}

export interface AuthResponse {
  user: User;
  token: string;
}

export const authApi = {
  register: async (data: RegisterRequest): Promise<AuthResponse> => {
    const response = await apiClient.post('/api/auth/register', data);
    return response.data;
  },

  login: async (data: LoginRequest): Promise<AuthResponse> => {
    const response = await apiClient.post('/api/auth/login', data);
    return response.data;
  },

  getMe: async (): Promise<User> => {
    const response = await apiClient.get('/api/auth/me');
    return response.data;
  },

  updateProfile: async (data: Partial<User>): Promise<User> => {
    const response = await apiClient.put('/api/users/me/profile', data);
    return response.data;
  },
};
