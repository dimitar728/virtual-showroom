export async function deleteRequest(path: string) {
  const res = await fetch(`${API_URL}${path.startsWith('/') ? path : '/' + path}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
      ...(getToken() ? { Authorization: `Bearer ${getToken()}` } : {}),
    },
  });
  if (!res.ok) throw new Error(await res.text());
  return await res.json();
}
const API_URL = "http://localhost:8080/api";

export async function register(email: string, password: string) {
  const res = await fetch(`${API_URL}/auth/register`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  });
  return res.json();
}

export async function login(email: string, password: string) {
  const res = await fetch(`${API_URL}/auth/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  });
  return res.json();
}

export async function getMe(token: string) {
  const res = await fetch(`${API_URL}/auth/me`, {
    headers: { Authorization: `Bearer ${token}` },
  });
  return res.json();
}

// Generic API helpers
function getToken() {
  return localStorage.getItem('token');
}

export async function get(path: string) {
  const res = await fetch(`${API_URL}${path.startsWith('/') ? path : '/' + path}`, {
    headers: {
      'Content-Type': 'application/json',
      ...(getToken() ? { Authorization: `Bearer ${getToken()}` } : {}),
    },
  });
  if (!res.ok) throw new Error(await res.text());
  return await res.json();
}

export async function post(path: string, body: any) {
  const res = await fetch(`${API_URL}${path.startsWith('/') ? path : '/' + path}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      ...(getToken() ? { Authorization: `Bearer ${getToken()}` } : {}),
    },
    body: JSON.stringify(body),
  });
  if (!res.ok) throw new Error(await res.text());
  return await res.json();
}

export async function patch(path: string, body: any) {
  const res = await fetch(`${API_URL}${path.startsWith('/') ? path : '/' + path}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
      ...(getToken() ? { Authorization: `Bearer ${getToken()}` } : {}),
    },
    body: JSON.stringify(body),
  });
  if (!res.ok) throw new Error(await res.text());
  return await res.json();
}