import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { fetchAllBookings } from '../store/bookingSlice';

const AdminBookings = () => {
  const dispatch = useDispatch();
  const { allBookings = [], loading, error } = useSelector((state) => state.booking || {});

  useEffect(() => {
    dispatch(fetchAllBookings());
  }, [dispatch]);

  if (loading) return <div className="flex justify-center items-center h-64"><span className="animate-spin rounded-full h-8 w-8 border-t-2 border-b-2 border-blue-500"></span></div>;
  if (error) return <div className="text-red-500 text-center mt-8">{error}</div>;
  if (!allBookings || allBookings.length === 0) return <div className="text-gray-500 text-center mt-8">No booking found</div>;

  return (
    <div className="max-w-4xl mx-auto mt-10 p-6 bg-white rounded-xl shadow-lg">
      <h2 className="text-3xl font-bold mb-6 text-blue-700">All Bookings</h2>
      <ul className="space-y-4">
        {allBookings.map((booking) => (
          <li key={booking.id} className="bg-gray-50 rounded-lg shadow p-4 hover:bg-blue-50 transition">
            <div className="flex flex-col md:flex-row md:justify-between md:items-center">
              <div>
                <span className="font-semibold text-gray-700">User:</span> <span className="text-blue-600">{booking.user_id}</span>
              </div>
              <div>
                <span className="font-semibold text-gray-700">Showroom:</span> <span className="text-blue-600">{booking.showroom_id}</span>
              </div>
              <div>
                <span className="font-semibold text-gray-700">Time:</span> <span className="text-green-600">{new Date(booking.slot_time).toLocaleString()}</span>
              </div>
              <div>
                <span className="font-semibold text-gray-700">Status:</span> <span className={`px-2 py-1 rounded ${booking.status === 'pending' ? 'bg-yellow-100 text-yellow-800' : 'bg-green-100 text-green-800'}`}>{booking.status}</span>
              </div>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default AdminBookings;