import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { fetchShowrooms } from '../store/showroomSlice';
import { Link } from 'react-router-dom';

const ShowroomList = () => {
  const dispatch = useDispatch();
  const { showrooms, loading, error } = useSelector((state) => state.showroom);

  useEffect(() => {
    dispatch(fetchShowrooms());
  }, [dispatch]);

  if (loading) return <div>Loading showrooms...</div>;
  if (error) return <div className="text-red-500">{error}</div>;

  return (
    <div className="max-w-4xl mx-auto mt-8">
      <h2 className="text-2xl font-bold mb-4">Showrooms</h2>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        {showrooms.map((room) => (
          <div key={room.id} className="border rounded p-4 bg-white shadow">
            <h3 className="text-lg font-semibold mb-2">{room.name}</h3>
            <p className="mb-2">{room.description}</p>
            <Link to={`/showrooms/${room.id}`} className="text-blue-600 hover:underline">
              View 3D Explorer
            </Link>
          </div>
        ))}
      </div>
    </div>
  );
};

export default ShowroomList;
