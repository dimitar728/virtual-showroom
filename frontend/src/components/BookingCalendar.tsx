import { useEffect, useState } from "react";
import { fetchBookings } from "../services/bookingService";
import type { Booking } from "../types";
import { Calendar } from "@/components/ui/calendar";

type Props = {
  showroomId: string;
  onSelectDate: (date: Date) => void;
};

export default function BookingCalendar({ showroomId, onSelectDate }: Props) {
  const [bookings, setBookings] = useState<Booking[]>([]);
  const [selected, setSelected] = useState<Date | undefined>(undefined);

  useEffect(() => {
    fetchBookings(showroomId).then(setBookings).catch(console.error);
  }, [showroomId]);

  const bookedDates = bookings.map((b) => new Date(b.start_time).toDateString());

  return (
    <div className="w-full p-4 rounded-2xl shadow bg-white">
      <h2 className="text-lg font-semibold mb-2">Booking Calendar</h2>
      <Calendar
        mode="single"
        selected={selected}
        onSelect={(d) => {
          if (!d) return;
          setSelected(d);
          onSelectDate(d);
        }}
        modifiers={{
          booked: (date) => bookedDates.includes(date.toDateString()),
        }}
        modifiersStyles={{
          booked: { backgroundColor: "#f87171", color: "white", borderRadius: "50%" },
        }}
      />
    </div>
  );
}