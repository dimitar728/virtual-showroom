import { useEffect, useState } from "react";
import { fetchUsers, suspendUser, reactivateUser, deleteUser } from "../services/userService";
import type { User } from "../types";

export default function UserManager() {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(false);

  const load = () => {
    setLoading(true);
    fetchUsers()
      .then(setUsers)
      .catch(console.error)
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    load();
  }, []);

  const handleAction = async (id: string, action: "suspend" | "reactivate" | "delete") => {
    try {
      if (action === "suspend") await suspendUser(id);
      if (action === "reactivate") await reactivateUser(id);
      if (action === "delete") {
        if (!window.confirm("Are you sure you want to delete this user?")) return;
        await deleteUser(id);
      }
      load();
    } catch (err) {
      console.error(err);
      alert("Failed to perform action");
    }
  };

  if (loading) return <div className="p-6">Loading users…</div>;

  return (
    <div className="p-6 bg-white rounded-2xl shadow">
      <h2 className="text-lg font-semibold mb-4">User Management</h2>
      <table className="w-full border-collapse text-sm">
        <thead>
          <tr className="bg-gray-100">
            <th className="border p-2 text-left">Email</th>
            <th className="border p-2 text-left">Role</th>
            <th className="border p-2 text-left">Status</th>
            <th className="border p-2 text-left">Actions</th>
          </tr>
        </thead>
        <tbody>
          {users.map((u) => (
            <tr key={u.id} className="hover:bg-gray-50">
              <td className="border p-2">{u.email}</td>
              <td className="border p-2">{u.role}</td>
              <td className="border p-2">{u.status}</td>
              <td className="border p-2 space-x-2">
                {u.status === "active" && (
                  <button
                    onClick={() => handleAction(u.id, "suspend")}
                    className="px-2 py-1 bg-yellow-500 text-white rounded hover:bg-yellow-600"
                  >
                    Suspend
                  </button>
                )}
                {u.status === "suspended" && (
                  <button
                    onClick={() => handleAction(u.id, "reactivate")}
                    className="px-2 py-1 bg-green-600 text-white rounded hover:bg-green-700"
                  >
                    Reactivate
                  </button>
                )}
                <button
                  onClick={() => handleAction(u.id, "delete")}
                  className="px-2 py-1 bg-red-600 text-white rounded hover:bg-red-700"
                >
                  Delete
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
