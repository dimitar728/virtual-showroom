import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
<<<<<<< HEAD
import Home from "./pages/Home";
=======
import Home from "./pages/Home"
>>>>>>> feature/HVSBS-110/UserManager-Suspend/delete/reactivate-users
import Login from "./pages/Login";
import Register from "./pages/Register";
import Dashboard from "./pages/Dashboard";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="*" element={<Navigate to="/" />} />
        <Route path="/dashboard" element={<ProtectedRoute><Dashboard /></ProtectedRoute>} />
        <Route path="/showrooms" element={<ProtectedRoute><Showrooms /></ProtectedRoute>} />

        <Route path="/admin"
          element={
            <RoleProtectedRoute allowedRoles={['admin']}>
              <AdminDashboard />
            </RoleProtectedRoute>
          }
        />
      </Routes>
    </BrowserRouter>
  );
}

export default App;