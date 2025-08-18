import { apiFetch } from "./apiClient";

const API = import.meta.env.VITE_API_URL || "http://localhost:8080";

export async function createShowroom(payload: { name: string; description: string; category: string }) {
  return apiFetch(`${API}/api/showrooms`, {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

export async function updateShowroom(id: string, payload: any) {
  return apiFetch(`${API}/api/showrooms/${id}`, {
    method: "PUT",
    body: JSON.stringify(payload),
  });
}

export async function uploadModel(id: string, file: File) {
  const formData = new FormData();
  formData.append("file", file);
  const res = await fetch(`${API}/api/showrooms/${id}/upload`, {
    method: "POST",
    body: formData,
  });
  if (!res.ok) throw new Error("Upload failed");
  return res.json();
}