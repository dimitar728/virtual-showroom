import React from 'react';
import { NavLink, Outlet } from 'react-router-dom';

const AdminPage = () => (
  <div className="flex min-h-screen">
    <aside className="w-64 bg-gray-800 text-white flex flex-col p-4">
      <h2 className="text-xl font-bold mb-6">Admin Panel</h2>
      <NavLink
        to="/admin/bookings"
        className={({ isActive }) =>
          `mb-4 hover:text-blue-300${isActive ? ' font-bold' : ''}`
        }
      >
        Booking
      </NavLink>
      <NavLink
        to="showrooms"
        className={({ isActive }) =>
          `mb-4 hover:text-blue-300${isActive ? ' font-bold' : ''}`
        }
      >
        Showroom
      </NavLink>
      <NavLink
        to="users"
        className={({ isActive }) =>
          `mb-4 hover:text-blue-300${isActive ? ' font-bold' : ''}`
        }
      >
        User
      </NavLink>
    </aside>
    <main className="flex-1 p-8 bg-gray-100">
      <Outlet />
    </main>
  </div>
);

export default AdminPage;