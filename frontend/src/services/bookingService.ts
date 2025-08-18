import { apiFetch } from "./apiClient";
import type { Booking } from "../types";

const API = import.meta.env.VITE_API_URL || "http://localhost:8080";

export async function fetchBookings(showroomId: string): Promise<Booking[]> {
  return apiFetch(`${API}/api/showrooms/${showroomId}/bookings`);
}

export async function createBooking(showroomId: string, payload: Partial<Booking>) {
  return apiFetch(`${API}/api/showrooms/${showroomId}/bookings`, {
    method: "POST",
    body: JSON.stringify(payload),
  });
}