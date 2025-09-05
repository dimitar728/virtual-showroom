import React, { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { bookSlot } from '../store/bookingSlice';
import { fetchShowrooms } from '../store/showroomSlice';

const BookingForm = () => {
  const dispatch = useDispatch();
  const { showrooms } = useSelector((state) => state.showroom || { showrooms: [] });
  const { bookingLoading, bookingError } = useSelector((state) => state.booking || {});
  const [showroomId, setShowroomId] = useState('');
  const [slotTime, setSlotTime] = useState('');

  useEffect(() => {
    dispatch(fetchShowrooms());
  }, [dispatch]);

  const handleSubmit = (e) => {
    e.preventDefault();
    // Validate showroomId is a UUID (simple check)
    const uuidRegex = /^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/;
    if (!showroomId || !uuidRegex.test(showroomId)) {
      alert('Please select a valid showroom.');
      return;
    }
    if (!slotTime) return;
    // Convert slotTime to ISO string with Z (UTC)
    const isoSlotTime = new Date(slotTime).toISOString();
    dispatch(bookSlot({ showroomId, slot_time: isoSlotTime }));
  };

  return (
    <div>
      <h2 className="text-xl font-bold mb-4">Book a Showroom</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block mb-1">Select Showroom</label>
          <select
            value={showroomId}
            onChange={(e) => setShowroomId(e.target.value)}
            className="border p-2 rounded w-full"
            required
          >
            <option value="">Choose a showroom</option>
            {showrooms.map((room) => (
              <option key={room.id} value={room.id}>{room.name}</option>
            ))}
          </select>
        </div>
        <div>
          <label className="block mb-1">Date & Time</label>
          <input
            type="datetime-local"
            value={slotTime}
            onChange={(e) => setSlotTime(e.target.value)}
            className="border p-2 rounded w-full"
            required
          />
        </div>
        <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded" disabled={bookingLoading}>
          {bookingLoading ? 'Booking...' : 'Book'}
        </button>
        {bookingError && <div className="text-red-500 mt-2">{bookingError}</div>}
      </form>
    </div>
  );
};

export default BookingForm;
