import { configureStore } from '@reduxjs/toolkit';
import { combineReducers } from 'redux';
import userReducer from './userSlice';
import authReducer from './authSlice';
import showroomReducer from './showroomSlice';
import bookingReducer from './bookingSlice'; 



export const store = configureStore({
  reducer: {
    user: userReducer,
    auth: authReducer,
    showroom: showroomReducer,
    booking: bookingReducer,
  },
});

const rootReducer = combineReducers({
  auth: authReducer,
  user: userReducer,
  showroom: showroomReducer,
  booking: bookingReducer,
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
export default store;
export { rootReducer }; // <-- named export