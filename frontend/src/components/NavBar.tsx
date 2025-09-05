import React from "react";
import { Link } from "react-router-dom";
import { useSelector, useDispatch } from "react-redux";
import { selectIsAuthenticated, logout } from "../store/authSlice";
import { AppDispatch } from "../store";

const Navbar: React.FC = () => {
  const isAuthenticated = useSelector(selectIsAuthenticated);
  const dispatch = useDispatch<AppDispatch>();

  const handleLogout = () => {
    dispatch(logout());
  };

  return (
    <nav className="bg-sky-700 p-4">
      <div className="container mx-auto flex justify-between items-center">
        <Link to="/" className="text-white text-xl font-bold">
          React / Golang - Virtual Showroom
        </Link>
        <div>
          {isAuthenticated ? (
            <>
              {/* <Link to="/" className="text-white mr-4">
                View Users
              </Link> */}
              <button onClick={handleLogout} className="text-white">
                Logout
              </button>
            </>
          ) : (
            <Link to="/" className="text-white">
              Home
            </Link>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
