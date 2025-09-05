import React from 'react';
import { NavLink, Outlet } from 'react-router-dom';

const UserPage = () => {
  return (
    <div className="flex min-h-screen">
      {/* Sidebar */}
      <aside className="w-64 bg-gray-100 p-6 border-r">
        <nav className="flex flex-col gap-4">
          <NavLink
            to="showrooms"
            className={({ isActive }) =>
              isActive
                ? 'font-bold text-blue-600'
                : 'text-gray-700 hover:text-blue-500'
            }
          >
            Showroom List
          </NavLink>
          <NavLink
            to="booking"
            className={({ isActive }) =>
              isActive
                ? 'font-bold text-blue-600'
                : 'text-gray-700 hover:text-blue-500'
            }
          >
            Booking Form
          </NavLink>
        </nav>
      </aside>
      {/* Main Content */}
      <main className="flex-1 p-8">
        <Outlet />
      </main>
    </div>
  );
};

export default UserPage;
