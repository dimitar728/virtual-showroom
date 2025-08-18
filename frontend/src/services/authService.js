export async function loginUser(credentials) {
  const response = await fetch('/api/auth/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(credentials),
  });

  if (!response.ok) {
    throw new Error('Login failed');
  }

  const data = await response.json();

  // Store JWT in localStorage
  localStorage.setItem('token', data.token);

  // Or store JWT in cookie (optional)
  document.cookie = `token=${data.token}; path=/; secure; samesite=strict`;

  return data;
}