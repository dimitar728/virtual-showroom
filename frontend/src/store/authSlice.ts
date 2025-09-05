import { createSlice, createAsyncThunk, PayloadAction } from "@reduxjs/toolkit";
import * as api from '../services/api';

interface AuthState {
  token: string | null;
  user: any;
  loading: boolean;
  error: string | null;
}

const initialState: AuthState = {
  token: localStorage.getItem('token'),
  user: null,
  loading: false,
  error: null,
};
// Async thunks
export const loginUser = createAsyncThunk<
  any,
  { email: string; password: string },
  { rejectValue: string }
>(
  'auth/loginUser',
  async ({ email, password }, thunkAPI) => {
    try {
      const res = await api.post('/auth/login', { email, password });
      localStorage.setItem('token', res.token);
      // Fetch user info after login
      const user = await api.getMe(res.token);
      return { token: res.token, user };
    } catch (err) {
      return thunkAPI.rejectWithValue(err instanceof Error ? err.message : 'Login failed');
    }
  }
);

export const registerUser = createAsyncThunk<
  any,
  { name: string; email: string; password: string },
  { rejectValue: string }
>(
  'auth/registerUser',
  async ({ name, email, password }, thunkAPI) => {
    try {
      await api.post('/auth/register', { name, email, password });
      return {};
    } catch (err) {
      return thunkAPI.rejectWithValue(err instanceof Error ? err.message : 'Registration failed');
    }
  }
);

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    setToken(state, action: PayloadAction<string>) {
      state.token = action.payload;
      localStorage.setItem('token', action.payload);
    },
    setUser(state, action: PayloadAction<any>) {
      state.user = action.payload;
    },
    logout(state) {
      state.token = null;
      state.user = null;
      localStorage.removeItem('token');
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(loginUser.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(loginUser.fulfilled, (state, action) => {
        state.loading = false;
        state.token = action.payload.token;
        state.user = action.payload.user;
      })
      .addCase(loginUser.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload || 'Login failed';
      })
      .addCase(registerUser.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(registerUser.fulfilled, (state) => {
        state.loading = false;
      })
      .addCase(registerUser.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload || 'Registration failed';
      });
  },
});

export const { setToken, setUser, logout } = authSlice.actions;
export default authSlice.reducer;
export const selectIsAuthenticated = (state: any) => !!state.auth.token;