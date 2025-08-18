import { Navigate } from 'react-router-dom';
import jwtDecode from 'jwt-decode';

export default function RoleProtectedRoute({ children, allowedRoles }) {
  const token = localStorage.getItem('token');

  if (!token) {
    return <Navigate to="/login" />;
  }

  try {
    const decoded = jwtDecode(token);
    if (!allowedRoles.includes(decoded.role)) {
      return <Navigate to="/unauthorized" />; // or a custom 403 page
    }
    return children;
  } catch (error) {
    return <Navigate to="/login" />;
  }
}