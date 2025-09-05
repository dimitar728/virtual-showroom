export interface User {
  id: string; // UUID
  email: string;
  role: string;
  created_at: string;
}

export interface UserCredentials {
  email: string;
  password: string;
}