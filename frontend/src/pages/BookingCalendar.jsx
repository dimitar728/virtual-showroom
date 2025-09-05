import React, { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { fetchAvailableSlots } from '../store/bookingSlice';

const BookingCalendar = ({ showroomId, onSlotSelect }) => {
  const dispatch = useDispatch();
  const { availableSlots, loading, error } = useSelector((state) => state.booking);
  const [selectedDate, setSelectedDate] = useState('');

  useEffect(() => {
    if (showroomId) dispatch(fetchAvailableSlots(showroomId));
  }, [dispatch, showroomId]);

  if (loading) return <div>Loading slots...</div>;
  if (error) return <div className="text-red-500">{error}</div>;

  return (
    <div>
      <h3 className="font-semibold mb-2">Available Slots</h3>
      <ul>
        {availableSlots.map((slot) => (
          <li key={slot.id}>
            <button
              className={`px-2 py-1 rounded ${selectedDate === slot.date ? 'bg-blue-600 text-white' : 'bg-gray-200'}`}
              onClick={() => { setSelectedDate(slot.date); onSlotSelect(slot); }}
            >
              {new Date(slot.date).toLocaleString()}
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default BookingCalendar;
