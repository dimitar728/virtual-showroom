import { apiFetch } from "./apiClient";
import type { Hotspot } from "../types";

const API = import.meta.env.VITE_API_URL || "http://localhost:8080";

export async function fetchHotspots(showroomId: string): Promise<Hotspot[]> {
  return apiFetch(`${API}/api/showrooms/${showroomId}/hotspots`);
}

export async function createHotspot(showroomId: string, payload: Partial<Hotspot>) {
  return apiFetch(`${API}/api/showrooms/${showroomId}/hotspots`, {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

export async function updateHotspot(hotspotId: string, payload: Partial<Hotspot>) {
  return apiFetch(`${API}/api/hotspots/${hotspotId}`, {
    method: "PATCH",
    body: JSON.stringify(payload),
  });
}

export async function deleteHotspot(hotspotId: string) {
  return apiFetch(`${API}/api/hotspots/${hotspotId}`, { method: "DELETE" });
}