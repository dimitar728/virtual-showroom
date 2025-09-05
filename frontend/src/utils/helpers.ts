export function saveToken(token: string) {
  localStorage.setItem("token", token);
}

export function getToken() {
  return localStorage.getItem("token");
}

export function formatDate(dateStr: string) {
  return new Date(dateStr).toLocaleString();
}