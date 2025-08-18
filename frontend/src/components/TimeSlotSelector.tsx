import { useEffect, useState } from "react";
import { generateTimeSlots } from "../utils/timeSlots";
import type { Booking } from "../types";

type Props = {
  date: Date | null;
  existingBookings: Booking[];
  onSelect: (slot: { start: string; end: string } | null) => void;
};

export default function TimeSlotSelector({ date, existingBookings, onSelect }: Props) {
  const [selected, setSelected] = useState<string | null>(null);

  if (!date) {
    return <div className="p-3 text-gray-500">Select a date to view available slots.</div>;
  }

  const slots = generateTimeSlots();

  // mark booked slots
  const booked = existingBookings.map((b) => ({
    start: new Date(b.start_time).toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" }),
    end: new Date(b.end_time).toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" }),
  }));

  const isBooked = (slot: { start: string; end: string }) =>
    booked.some((b) => b.start === slot.start && b.end === slot.end);

  return (
    <div className="grid grid-cols-2 gap-2">
      {slots.map((slot) => {
        const disabled = isBooked(slot);
        const isActive = selected === `${slot.start}-${slot.end}`;
        return (
          <button
            key={`${slot.start}-${slot.end}`}
            onClick={() => {
              if (disabled) return;
              setSelected(`${slot.start}-${slot.end}`);
              onSelect(slot);
            }}
            disabled={disabled}
            className={`px-3 py-2 rounded-lg border text-sm ${
              disabled
                ? "bg-gray-200 text-gray-500 cursor-not-allowed"
                : isActive
                ? "bg-blue-600 text-white"
                : "bg-white hover:bg-blue-50"
            }`}
          >
            {slot.start} – {slot.end}
          </button>
        );
      })}
    </div>
  );
}