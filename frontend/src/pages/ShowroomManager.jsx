import React, { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { fetchShowrooms, createShowroom, updateShowroom, deleteShowroom } from '../store/showroomSlice';

const ShowroomManager = () => {
  const dispatch = useDispatch();
  const { showrooms, loading, error } = useSelector((state) => state.showroom);
  const [form, setForm] = useState({ name: '', description: '', model_path: '', capacity: 0 });
  const [editingId, setEditingId] = useState(null);

  useEffect(() => {
    dispatch(fetchShowrooms());
  }, [dispatch]);

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (editingId) {
      dispatch(updateShowroom({ ...form, id: editingId }));
    } else {
      dispatch(createShowroom({
        name: form.name,
        description: form.description,
        model_path: form.model_path,
        capacity: Number(form.capacity)
  }));
    }
    setForm({ name: '', description: '', model_path: '', capacity: 0 });
    setEditingId(null);
  };

  const handleEdit = (room) => {
    setForm(room);
    setEditingId(room.id);
  };

  const handleDelete = (id) => {
    dispatch(deleteShowroom(id));
  };

  return (
    <div className="max-w-4xl mx-auto mt-8">
      <h2 className="text-2xl font-bold mb-4">Showroom Manager</h2>
      <form onSubmit={handleSubmit} className="mb-6 p-4 bg-white rounded shadow">
        <input name="name" value={form.name} onChange={handleChange} placeholder="Name" className="border p-2 mr-2" required />
        <input name="description" value={form.description} onChange={handleChange} placeholder="Description" className="border p-2 mr-2" required />
        <input name="model_path" value={form.model_path} onChange={handleChange} placeholder="Model URL" className="border p-2 mr-2" required />
        <input name="capacity" type="number" value={form.capacity} onChange={handleChange} placeholder="Capacity" className="border p-2 mr-2" required />
        <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded">{editingId ? 'Update' : 'Create'}</button>
      </form>
      {loading && <div>Loading...</div>}
      {error && <div className="text-red-500">{error}</div>}
      <ul>
        {showrooms.map((room) => (
          <li key={room.id} className="border rounded p-3 mb-2 flex justify-between items-center">
            <div>
              <b>{room.name}</b> - {room.description} (Capacity: {room.capacity})
            </div>
            <div>
              <button className="bg-yellow-500 text-white px-3 py-1 rounded mr-2" onClick={() => handleEdit(room)}>Edit</button>
              <button className="bg-red-500 text-white px-3 py-1 rounded" onClick={() => handleDelete(room.id)}>Delete</button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ShowroomManager;
