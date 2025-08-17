export async function apiFetch(url, options = {}) {
  const token = localStorage.getItem('token');

  const headers = {
    ...options.headers,
    'Content-Type': 'application/json',
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
  };

  const response = await fetch(url, { ...options, headers });
  return response.json();
}