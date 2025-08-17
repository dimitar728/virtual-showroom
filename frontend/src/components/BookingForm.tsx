import { useState } from "react";
import { createBooking } from "../services/bookingService";

type Props = {
  showroomId: string;
  date: Date | null;
  onBooked: () => void;
};

export default function BookingForm({ showroomId, date, onBooked }: Props) {
  const [start, setStart] = useState("10:00");
  const [end, setEnd] = useState("11:00");
  const [loading, setLoading] = useState(false);

  if (!date) {
    return <div className="p-4 text-gray-500">Select a date from the calendar</div>;
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    try {
      const startDate = new Date(date);
      const endDate = new Date(date);
      const [sh, sm] = start.split(":").map(Number);
      const [eh, em] = end.split(":").map(Number);
      startDate.setHours(sh, sm);
      endDate.setHours(eh, em);

      await createBooking(showroomId, {
        start_time: startDate.toISOString(),
        end_time: endDate.toISOString(),
      });
      onBooked();
    } catch (err) {
      console.error(err);
      alert("Failed to book slot");
    } finally {
      setLoading(false);
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="w-full p-4 mt-4 rounded-2xl shadow bg-white space-y-3"
    >
      <h2 className="text-lg font-semibold">Book Slot on {date.toDateString()}</h2>

      <div className="flex gap-2">
        <label className="flex-1">
          Start Time
          <input
            type="time"
            value={start}
            onChange={(e) => setStart(e.target.value)}
            className="w-full border rounded px-2 py-1"
          />
        </label>
        <label className="flex-1">
          End Time
          <input
            type="time"
            value={end}
            onChange={(e) => setEnd(e.target.value)}
            className="w-full border rounded px-2 py-1"
          />
        </label>
      </div>

      <button
        type="submit"
        disabled={loading}
        className="w-full py-2 rounded-xl bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50"
      >
        {loading ? "Booking..." : "Confirm Booking"}
      </button>
    </form>
  );
}