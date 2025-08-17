import { useEffect, useState } from "react";
import { fetchAllBookings } from "../services/bookingService";
import type { Booking } from "../types";

export default function BookingManager() {
  const [bookings, setBookings] = useState<Booking[]>([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    setLoading(true);
    fetchAllBookings()
      .then(setBookings)
      .catch(console.error)
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <div className="p-6">Loading all bookings…</div>;

  return (
    <div className="p-6 bg-white rounded-2xl shadow">
      <h2 className="text-lg font-semibold mb-4">All Bookings</h2>
      {bookings.length === 0 ? (
        <p>No bookings found.</p>
      ) : (
        <table className="w-full border-collapse text-sm">
          <thead>
            <tr className="bg-gray-100">
              <th className="border p-2 text-left">User</th>
              <th className="border p-2 text-left">Showroom</th>
              <th className="border p-2 text-left">Start</th>
              <th className="border p-2 text-left">End</th>
            </tr>
          </thead>
          <tbody>
            {bookings.map((b) => (
              <tr key={b.id} className="hover:bg-gray-50">
                <td className="border p-2">{b.user_id}</td>
                <td className="border p-2">{b.showroom_id}</td>
                <td className="border p-2">{new Date(b.start_time).toLocaleString()}</td>
                <td className="border p-2">{new Date(b.end_time).toLocaleString()}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}