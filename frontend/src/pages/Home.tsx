import React from "react";
import { useSelector } from "react-redux";
import { selectIsAuthenticated } from "../store/authSlice";
import Login from "./Login";
import Register from "./Register";

const Home: React.FC = () => {

  const isAuthenticated = useSelector(selectIsAuthenticated);

  if (isAuthenticated) {
    
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-900 via-blue-700 to-sky-400 flex items-center justify-center">
      <div className="max-w-4xl w-full mx-auto p-8 bg-white/90 rounded-2xl shadow-2xl">
        <h1 className="text-4xl font-extrabold mb-4 text-blue-800 text-center drop-shadow">Welcome to Virtual Showroom</h1>
        <p className="mb-8 text-lg text-blue-600 text-center">Please log in or register to continue.</p>
        <div className="flex flex-col md:flex-row md:space-x-8 space-y-8 md:space-y-0">
          <div className="w-full md:w-1/2">
            <Login />
          </div>
          <div className="w-full md:w-1/2">
            <Register />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Home;