import React from "react";
import {
  BrowserRouter as Router,
  Route,
  Routes,
} from "react-router-dom";
import Home from "./pages/Home";
import Navbar from "./components/NavBar";
import AdminPage from './pages/AdminPage';
import BookingManager from './pages/BookingManager';
import ShowroomManager from './pages/ShowroomManager';
import UserManager from './pages/UserManager';
import UserPage from './pages/UserPage';
import ShowroomList from './pages/ShowroomList';
import BookingForm from './pages/BookingForm';
import Room3D from './pages/Room3D';
import AdminBookings from './pages/AdminBookings';

const App: React.FC = () => {
  return (
    <Router>
      <div className="min-h-screen bg-black">
        <Navbar />
        <main className="container mx-auto px-4 py-8">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/admin" element={<AdminPage />}>
              <Route path="bookings" element={<AdminBookings />} />
              <Route path="showrooms" element={<ShowroomManager />} />
              <Route path="users" element={<UserManager />} />
              <Route index element={<BookingManager />} /> {/* Default admin page */}
            </Route>
              <Route path="/user" element={<UserPage />}>
                <Route path="showrooms" element={<ShowroomList />} />
                <Route path="booking" element={<BookingForm />} />
                <Route index element={<ShowroomList />} /> {/* Default user page */}
              </Route>
              <Route path="/showrooms/:id" element={<Room3D />} />
              {/* <Route path="/admin/bookings" element={<AdminBookings />} /> */}
          </Routes>
        </main>
      </div>
    </Router>
  );
};

export default App;