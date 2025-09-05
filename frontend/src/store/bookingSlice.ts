import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import * as api from '../services/api';

// Types
export interface Slot {
  id: string;
  date: string;
}
export interface Booking {
  id: string;
  showroom_id: string;
  showroom_name?: string;
  user_email?: string;
  slot_time: string; // <-- use slot_time
  status: string;
  created_at?: string;
}
export interface BookingState {
  availableSlots: Slot[];
  myBookings: Booking[];
  allBookings: Booking[];
  loading: boolean;
  error: string | null;
  bookingLoading: boolean;
  bookingError: string | null;
}

// Helper for error extraction
function getErrorMessage(err: unknown): string {
  if (typeof err === 'object' && err && 'response' in err && (err as any).response?.data?.error) {
    return (err as any).response.data.error;
  }
  if (err instanceof Error) return err.message;
  return 'Unknown error';
}

// Fetch available slots for a showroom
export const fetchAvailableSlots = createAsyncThunk<Slot[], string, { rejectValue: string }>(
  'booking/fetchAvailableSlots',
  async (showroomId, thunkAPI) => {
    try {
      const res = await api.get(`/showrooms/${showroomId}/slots`);
      return res;
    } catch (err) {
      return thunkAPI.rejectWithValue(getErrorMessage(err));
    }
  }
);

// Book a slot
export const bookSlot = createAsyncThunk<
  Booking,
  { showroomId: string; slot_time: string },
  { rejectValue: string }
>(
  'booking/bookSlot',
  async ({ showroomId, slot_time }, thunkAPI) => {
    try {
      const res = await api.post('/bookings', { showroom_id: showroomId, slot_time });
      return res;
    } catch (err) {
      return thunkAPI.rejectWithValue(getErrorMessage(err));
    }
  }
);

// Fetch user's bookings
export const fetchMyBookings = createAsyncThunk<Booking[], void, { rejectValue: string }>(
  'booking/fetchMyBookings', // <-- change this line!
  async (_, thunkAPI) => {
    try {
      const res = await api.get('/bookings');
      return res.data;
    } catch (err) {
      return thunkAPI.rejectWithValue(getErrorMessage(err));
    }
  }
);

// Cancel a booking
export const cancelBooking = createAsyncThunk<string, string, { rejectValue: string }>(
  'booking/cancelBooking',
  async (bookingId, thunkAPI) => {
    try {
  await api.patch(`/bookings/${bookingId}/cancel`, {});
      return bookingId;
    } catch (err) {
      return thunkAPI.rejectWithValue(getErrorMessage(err));
    }
  }
);

// Admin: fetch all bookings
export const fetchAllBookings = createAsyncThunk<Booking[], void, { rejectValue: string }>(
  'booking/fetchAllBookings',
  async (_, thunkAPI) => {
    try {
      const res = await api.get('/admin/bookings');
      return res; // <-- FIX: return the parsed array directly
    } catch (err) {
      return thunkAPI.rejectWithValue(getErrorMessage(err));
    }
  }
);

// Admin: update booking status
export const updateBookingStatus = createAsyncThunk<
  Booking,
  { id: string; status: string },
  { rejectValue: string }
>(
  'booking/updateBookingStatus',
  async ({ id, status }, thunkAPI) => {
    try {
      const res = await api.patch(`/admin/bookings/${id}`, { status });
      return res;
    } catch (err) {
      return thunkAPI.rejectWithValue(getErrorMessage(err));
    }
  }
);

const initialState: BookingState = {
  availableSlots: [],
  myBookings: [],
  allBookings: [],
  loading: false,
  error: null,
  bookingLoading: false,
  bookingError: null,
};

const bookingSlice = createSlice({
  name: 'booking',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      // Fetch available slots
      .addCase(fetchAvailableSlots.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchAvailableSlots.fulfilled, (state, action: PayloadAction<Slot[]>) => {
        state.loading = false;
        state.availableSlots = action.payload;
      })
      .addCase(fetchAvailableSlots.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload || null;
      })
      // Book slot
      .addCase(bookSlot.pending, (state) => {
        state.bookingLoading = true;
        state.bookingError = null;
      })
      .addCase(bookSlot.fulfilled, (state, action: PayloadAction<Booking>) => {
        state.bookingLoading = false;
        state.myBookings.push(action.payload);
      })
      .addCase(bookSlot.rejected, (state, action) => {
        state.bookingLoading = false;
        state.bookingError = action.payload || null;
      })
      // Fetch my bookings
      .addCase(fetchMyBookings.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchMyBookings.fulfilled, (state, action: PayloadAction<Booking[]>) => {
        state.loading = false;
        state.myBookings = action.payload;
      })
      .addCase(fetchMyBookings.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload || null;
      })
      // Cancel booking
      .addCase(cancelBooking.fulfilled, (state, action: PayloadAction<string>) => {
        state.myBookings = state.myBookings.map((b) =>
          b.id === action.payload ? { ...b, status: 'cancelled' } : b
        );
      })
      // Admin: fetch all bookings
      .addCase(fetchAllBookings.fulfilled, (state, action: PayloadAction<Booking[]>) => {
        state.allBookings = action.payload;
      })
      // Admin: update booking status
      .addCase(updateBookingStatus.fulfilled, (state, action: PayloadAction<Booking>) => {
        state.allBookings = state.allBookings.map((b) =>
          b.id === action.payload.id ? action.payload : b
        );
      });
  },
});

export default bookingSlice.reducer;
