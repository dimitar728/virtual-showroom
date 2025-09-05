import { createSlice, createAsyncThunk, PayloadAction } from "@reduxjs/toolkit";
import * as api from '../services/api';
import { User } from "../types/user";

export const fetchUsers = createAsyncThunk<User[], void, { rejectValue: string }>(
  'user/fetchUsers',
  async (_, thunkAPI) => {
    try {
      const res = await api.get('/admin/users');
      return res;
    } catch (err) {
      return thunkAPI.rejectWithValue(err instanceof Error ? err.message : 'Failed to fetch users');
    }
  }
);

export const suspendUser = createAsyncThunk<User, string, { rejectValue: string }>(
  'user/suspendUser',
  async (id, thunkAPI) => {
    try {
      const res = await api.patch(`/admin/users/${id}/suspend`, {});
      return res;
    } catch (err) {
      return thunkAPI.rejectWithValue(err instanceof Error ? err.message : 'Failed to suspend user');
    }
  }
);

export const reactivateUser = createAsyncThunk<User, string, { rejectValue: string }>(
  'user/reactivateUser',
  async (id, thunkAPI) => {
    try {
      const res = await api.patch(`/admin/users/${id}/reactivate`, {});
      return res;
    } catch (err) {
      return thunkAPI.rejectWithValue(err instanceof Error ? err.message : 'Failed to reactivate user');
    }
  }
);

export const deleteUser = createAsyncThunk<string, string, { rejectValue: string }>(
  'user/deleteUser',
  async (id, thunkAPI) => {
    try {
      await api.deleteRequest(`/admin/users/${id}`);
      return id;
    } catch (err) {
      return thunkAPI.rejectWithValue(err instanceof Error ? err.message : 'Failed to delete user');
    }
  }
);

const initialState: { users: User[]; loading?: boolean; error?: string | null } = {
  users: [],
  loading: false,
  error: null,
};

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    setUsers(state, action: PayloadAction<User[]>) {
      state.users = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchUsers.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchUsers.fulfilled, (state, action: PayloadAction<User[]>) => {
        state.loading = false;
        state.users = action.payload;
      })
      .addCase(fetchUsers.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload || 'Failed to fetch users';
      })
      .addCase(suspendUser.fulfilled, (state, action: PayloadAction<User>) => {
        state.users = state.users.map((u) =>
          u.id === action.payload.id ? action.payload : u
        );
      })
      .addCase(reactivateUser.fulfilled, (state, action: PayloadAction<User>) => {
        state.users = state.users.map((u) =>
          u.id === action.payload.id ? action.payload : u
        );
      })
      .addCase(deleteUser.fulfilled, (state, action: PayloadAction<string>) => {
        state.users = state.users.filter((u) => u.id !== action.payload);
      });
  },
});

export const { setUsers } = userSlice.actions;
export default userSlice.reducer;
