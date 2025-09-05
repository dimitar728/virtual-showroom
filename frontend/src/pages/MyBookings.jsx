import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { fetchAllBookings, cancelBooking } from '../store/bookingSlice';

const MyBookings = () => {
  const dispatch = useDispatch();
  const { myBookings, loading, error } = useSelector((state) => state.booking);

  useEffect(() => {
    dispatch(fetchAllBookings());
  }, [dispatch]);

  if (loading) return <div>Loading your bookings...</div>;
  if (error) return <div className="text-red-500">{error}</div>;

  return (
    <div className="max-w-2xl mx-auto mt-8">
      <h2 className="text-xl font-bold mb-4">My Bookings</h2>
      <ul>
        {myBookings.map((booking) => (
          <li key={booking.id} className="border rounded p-3 mb-2 flex justify-between items-center">
            <div>
              <div><b>Showroom:</b> {booking.showroom_name}</div>
              <div><b>Date:</b> {new Date(booking.date).toLocaleString()}</div>
              <div><b>Status:</b> {booking.status}</div>
            </div>
            {booking.status !== 'cancelled' && (
              <button
                className="bg-red-500 text-white px-3 py-1 rounded"
                onClick={() => dispatch(cancelBooking(booking.id))}
              >Cancel</button>
            )}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default MyBookings;
