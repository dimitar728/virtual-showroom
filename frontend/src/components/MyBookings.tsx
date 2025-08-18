import { useEffect, useState } from "react";
import { fetchMyBookings, cancelBooking } from "../services/bookingService";
import type { Booking } from "../types";

type Props = {
  userId: string;
};

export default function MyBookings({ userId }: Props) {
  const [bookings, setBookings] = useState<Booking[]>([]);
  const [loading, setLoading] = useState(false);

  const loadBookings = () => {
    setLoading(true);
    fetchMyBookings(userId)
      .then(setBookings)
      .catch(console.error)
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    loadBookings();
  }, [userId]);

  const handleCancel = async (id: string) => {
    if (!window.confirm("Cancel this booking?")) return;
    try {
      await cancelBooking(id);
      loadBookings();
    } catch (err) {
      console.error(err);
      alert("Failed to cancel booking");
    }
  };

  if (loading) return <div className="p-4">Loading your bookings…</div>;
  if (bookings.length === 0) return <div className="p-4">No bookings found.</div>;

  return (
    <div className="p-4 bg-white rounded-2xl shadow">
      <h2 className="text-lg font-semibold mb-4">My Bookings</h2>
      <ul className="divide-y">
        {bookings.map((b) => (
          <li key={b.id} className="py-3 flex items-center justify-between">
            <div>
              <p className="font-medium">
                {new Date(b.start_time).toLocaleString()} –{" "}
                {new Date(b.end_time).toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" })}
              </p>
              <p className="text-sm text-gray-600">Showroom: {b.showroom_id}</p>
            </div>
            <button
              onClick={() => handleCancel(b.id)}
              className="px-3 py-1 rounded-lg bg-red-600 text-white hover:bg-red-700 text-sm"
            >
              Cancel
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}