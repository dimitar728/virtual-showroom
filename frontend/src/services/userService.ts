import { apiFetch } from "./apiClient";
import type { User } from "../types";

const API = import.meta.env.VITE_API_URL || "http://localhost:8080";

export async function fetchUsers(): Promise<User[]> {
  return apiFetch(`${API}/api/users`);
}

export async function suspendUser(id: string) {
  return apiFetch(`${API}/api/users/${id}/suspend`, { method: "PATCH" });
}

export async function reactivateUser(id: string) {
  return apiFetch(`${API}/api/users/${id}/reactivate`, { method: "PATCH" });
}

export async function deleteUser(id: string) {
  return apiFetch(`${API}/api/users/${id}`, { method: "DELETE" });
}
