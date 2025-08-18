import { useEffect, useState } from 'react';

export default function ShowroomList() {
  const [showrooms, setShowrooms] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchShowrooms() {
      try {
        const response = await fetch('/api/showrooms');
        const data = await response.json();
        setShowrooms(data);
      } catch (error) {
        console.error('Failed to fetch showrooms:', error);
      } finally {
        setLoading(false);
      }
    }
    fetchShowrooms();
  }, []);

  if (loading) return <p>Loading showrooms...</p>;

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6 p-4">
      {showrooms.map(showroom => (
        <div key={showroom.id} className="bg-white rounded-lg shadow-md p-4">
          <img src={showroom.thumbnail} alt={showroom.name} className="w-full h-48 object-cover rounded-md" />
          <h2 className="text-xl font-semibold mt-2">{showroom.name}</h2>
          <p className="text-gray-600">{showroom.description}</p>
        </div>
      ))}
    </div>
  );
}