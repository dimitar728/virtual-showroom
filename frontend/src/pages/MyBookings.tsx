import MyBookings from "../components/MyBookings";

export default function MyBookingsPage() {
  const userId = localStorage.getItem("userId")!; // adjust based on auth
  return (
    <div className="max-w-3xl mx-auto p-6">
      <MyBookings userId={userId} />
    </div>
  );
}
