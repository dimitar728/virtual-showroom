import { useState } from "react";
import { createBooking } from "../services/bookingService";
import TimeSlotSelector from "./TimeSlotSelector";
import type { Booking } from "../types";

type Props = {
  showroomId: string;
  date: Date | null;
  existingBookings: Booking[];
  onBooked: () => void;
};

export default function BookingForm({ showroomId, date, existingBookings, onBooked }: Props) {
  const [slot, setSlot] = useState<{ start: string; end: string } | null>(null);
  const [loading, setLoading] = useState(false);

  if (!date) {
    return <div className="p-4 text-gray-500">Select a date from the calendar</div>;
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!slot) {
      alert("Please select a time slot.");
      return;
    }
    setLoading(true);
    try {
      const startDate = new Date(date);
      const endDate = new Date(date);

      const [sh, sm] = slot.start.split(":").map(Number);
      const [eh, em] = slot.end.split(":").map(Number);

      startDate.setHours(sh, sm);
      endDate.setHours(eh, em);

      await createBooking(showroomId, {
        start_time: startDate.toISOString(),
        end_time: endDate.toISOString(),
      });
      onBooked();
      setSlot(null);
    } catch (err) {
      console.error(err);
      alert("Failed to book slot");
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="w-full p-4 mt-4 rounded-2xl shadow bg-white space-y-3">
      <h2 className="text-lg font-semibold">Book Slot on {date.toDateString()}</h2>

      <TimeSlotSelector date={date} existingBookings={existingBookings} onSelect={setSlot} />

      <button
        type="submit"
        disabled={loading || !slot}
        className="w-full py-2 rounded-xl bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50"
      >
        {loading ? "Booking..." : "Confirm Booking"}
      </button>
    </form>
  );
}