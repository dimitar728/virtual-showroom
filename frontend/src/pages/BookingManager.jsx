// Example BookingManager.jsx
import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { fetchAllBookings } from '../store/bookingSlice';

const BookingManager = () => {
  const dispatch = useDispatch();
  const { bookings = [], loading, error } = useSelector((state) => state.booking || {});

  useEffect(() => {
    dispatch(fetchAllBookings());
  }, [dispatch]);

  if (loading) return <div>Loading...</div>;
  if (error) return <div className="text-red-500">{error}</div>;
  if (!bookings || bookings.length === 0) return <div>No booking found</div>;

  return (
    <div>
      <h2 className="text-2xl font-bold mb-4">Bookings</h2>
      <ul>
        {bookings.map((booking) => (
          <li key={booking.id}>
            {booking.showroom_id} - {booking.slot_time} - {booking.status}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default BookingManager;