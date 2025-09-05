import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import * as api from '../services/api';

export interface Showroom {
  id: string;
  name: string;
  description: string;
  model_path: string;
  capacity: number;
}

export const fetchShowrooms = createAsyncThunk<Showroom[], void, { rejectValue: string }>(
  'showroom/fetchShowrooms',
  async (_, thunkAPI) => {
    try {
      const res = await api.get('/showrooms');
      return res;
    } catch (err) {
      return thunkAPI.rejectWithValue(err instanceof Error ? err.message : 'Failed to fetch showrooms');
    }
  }
);

export const fetchShowroomById = createAsyncThunk<Showroom, string, { rejectValue: string }>(
  'showroom/fetchShowroomById',
  async (id, thunkAPI) => {
    try {
      const res = await api.get(`/showrooms/${id}`);
      return res;
    } catch (err) {
      return thunkAPI.rejectWithValue(err instanceof Error ? err.message : 'Failed to fetch showroom');
    }
  }
);

export const createShowroom = createAsyncThunk<Showroom, Partial<Showroom>, { rejectValue: string }>(
  'showroom/createShowroom',
  async (data, thunkAPI) => {
    try {
      const res = await api.post('/admin/showrooms', data);
      return res;
    } catch (err) {
      return thunkAPI.rejectWithValue(err instanceof Error ? err.message : 'Failed to create showroom');
    }
  }
);

export const updateShowroom = createAsyncThunk<Showroom, Showroom, { rejectValue: string }>(
  'showroom/updateShowroom',
  async (data, thunkAPI) => {
    try {
      const res = await api.patch(`/admin/showrooms/${data.id}`, data);
      return res;
    } catch (err) {
      return thunkAPI.rejectWithValue(err instanceof Error ? err.message : 'Failed to update showroom');
    }
  }
);

export const deleteShowroom = createAsyncThunk<string, string, { rejectValue: string }>(
  'showroom/deleteShowroom',
  async (id, thunkAPI) => {
    try {
      await api.deleteRequest(`/admin/showrooms/${id}`);
      return id;
    } catch (err) {
      return thunkAPI.rejectWithValue(err instanceof Error ? err.message : 'Failed to delete showroom');
    }
  }
);

const initialState = {
  showrooms: [] as Showroom[],
  selectedShowroom: null as Showroom | null,
  loading: false,
  error: null as string | null,
};

const showroomSlice = createSlice({
  name: 'showroom',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchShowrooms.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchShowrooms.fulfilled, (state, action: PayloadAction<Showroom[]>) => {
        state.loading = false;
        state.showrooms = action.payload;
      })
      .addCase(fetchShowrooms.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload || 'Failed to fetch showrooms';
      })
      .addCase(fetchShowroomById.fulfilled, (state, action: PayloadAction<Showroom>) => {
        state.selectedShowroom = action.payload;
      })
      .addCase(createShowroom.fulfilled, (state, action: PayloadAction<Showroom>) => {
        state.showrooms.push(action.payload);
      })
      .addCase(updateShowroom.fulfilled, (state, action: PayloadAction<Showroom>) => {
        state.showrooms = state.showrooms.map((s) =>
          s.id === action.payload.id ? action.payload : s
        );
      })
      .addCase(deleteShowroom.fulfilled, (state, action: PayloadAction<string>) => {
        state.showrooms = state.showrooms.filter((s) => s.id !== action.payload);
      });
  },
});

export default showroomSlice.reducer;
