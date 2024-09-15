import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import { User, LoginInput, RegisterInput } from '../../types/user';
import api from '../services/api';
import axios from 'axios'; // axiosをインポートしてエラーハンドリングを改善します

interface AuthState {
  user: User | null;
  isAuthenticated: boolean;
  status: 'idle' | 'loading' | 'succeeded' | 'failed';
  error: string | null;
}

const initialState: AuthState = {
  user: null,
  isAuthenticated: false,
  status: 'idle',
  error: null,
};

// 非同期アクションの定義
// Remove token storage
export const login = createAsyncThunk(
  'auth/login',
  async (input: LoginInput, { rejectWithValue }) => {
    try {
      const response = await api.post<User>('/auth/login', input);
      localStorage.setItem('username', response.data.username);
      // Remove token storage
      // localStorage.setItem('token', response.data.token);
      return response.data;
    } catch (error) {
      // Error handling remains the same
    }
  }
);

export const register = createAsyncThunk(
  'auth/register',
  async (input: RegisterInput, { rejectWithValue }) => {
    try {
      const response = await api.post<User>('/auth/register', input);
      return response.data;
    } catch (error) {
      if (axios.isAxiosError(error)) {
        return rejectWithValue(error.response?.data || 'Registration failed');
      }
      return rejectWithValue('Registration failed');
    }
  }
);

export const logout = createAsyncThunk('auth/logout', async () => {
  await api.post('/logout');
  localStorage.removeItem('username');
});

export const getProfileByUsername = createAsyncThunk(
  'auth/getProfileByUsername',
  async (username: string, { rejectWithValue }) => {
    try {
      const response = await api.get<User>(`/profile/${username}`);
      return response.data;
    } catch (error) {
      if (axios.isAxiosError(error)) {
        return rejectWithValue(error.response?.data || 'Failed to fetch profile');
      }
      return rejectWithValue('Failed to fetch profile');
    }
  }
);

const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(login.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(login.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.user = action.payload;
        state.isAuthenticated = true;
      })
      .addCase(login.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.payload as string;
      })
      .addCase(register.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(register.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.user = action.payload;
        state.isAuthenticated = true;
      })
      .addCase(register.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.payload as string;
      })
      .addCase(logout.fulfilled, (state) => {
        state.user = null;
        state.isAuthenticated = false;
      })
      .addCase(getProfileByUsername.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(getProfileByUsername.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.user = action.payload;
      })
      .addCase(getProfileByUsername.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.payload as string;
      });
  },
});

export default authSlice.reducer;