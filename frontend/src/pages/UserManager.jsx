import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { fetchUsers, suspendUser, reactivateUser, deleteUser } from '../store/userSlice';

const UserManager = () => {
  const dispatch = useDispatch();
  const { users, loading, error } = useSelector((state) => state.user);

  useEffect(() => {
    dispatch(fetchUsers());
  }, [dispatch]);

  return (
    <div className="max-w-4xl mx-auto mt-8">
      <h2 className="text-2xl font-bold mb-4">User Manager</h2>
      {loading && <div>Loading...</div>}
      {error && <div className="text-red-500">{error}</div>}
      <ul>
        {users.map((user) => (
          <li key={user.id} className="border rounded p-3 mb-2 flex justify-between items-center">
            <div>
              <b>{user.name || user.email}</b> - {user.email} ({user.role})
            </div>
            <div>
              {user.status !== 'suspended' ? (
                <button className="bg-yellow-600 text-white px-3 py-1 rounded mr-2" onClick={() => dispatch(suspendUser(user.id))}>Suspend</button>
              ) : (
                <button className="bg-green-600 text-white px-3 py-1 rounded mr-2" onClick={() => dispatch(reactivateUser(user.id))}>Reactivate</button>
              )}
              <button className="bg-red-500 text-white px-3 py-1 rounded" onClick={() => dispatch(deleteUser(user.id))}>Delete</button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default UserManager;
