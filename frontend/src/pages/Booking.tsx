import { useState } from "react";
import { useParams } from "react-router-dom";
import BookingCalendar from "../components/BookingCalendar";
import BookingForm from "../components/BookingForm";

export default function BookingPage() {
  const { id } = useParams(); // showroomId
  const [selectedDate, setSelectedDate] = useState<Date | null>(null);
  const [refreshKey, setRefreshKey] = useState(0);

  return (
    <div className="max-w-3xl mx-auto p-6 space-y-4">
      <BookingCalendar
        key={refreshKey}
        showroomId={id!}
        onSelectDate={(d) => setSelectedDate(d)}
      />
      <BookingForm
        showroomId={id!}
        date={selectedDate}
        onBooked={() => setRefreshKey((k) => k + 1)}
      />
    </div>
  );
}