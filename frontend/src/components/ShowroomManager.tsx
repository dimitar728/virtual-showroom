import { useState } from "react";
import { createShowroom, updateShowroom, uploadModel } from "../services/showroomService";

type Props = {
  showroomId?: string; // optional: edit existing
};

export default function ShowroomManager({ showroomId }: Props) {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [category, setCategory] = useState("");
  const [file, setFile] = useState<File | null>(null);
  const [loading, setLoading] = useState(false);

  const handleSave = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    try {
      let showroom: any;
      if (showroomId) {
        showroom = await updateShowroom(showroomId, { name, description, category });
      } else {
        showroom = await createShowroom({ name, description, category });
      }

      if (file) {
        await uploadModel(showroom.id || showroomId!, file);
      }

      alert("Showroom saved!");
    } catch (err) {
      console.error(err);
      alert("Failed to save showroom");
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSave} className="p-6 bg-white shadow rounded-2xl space-y-4">
      <h2 className="text-lg font-semibold">{showroomId ? "Edit Showroom" : "Create Showroom"}</h2>

      <div>
        <label className="block text-sm font-medium">Name</label>
        <input
          value={name}
          onChange={(e) => setName(e.target.value)}
          className="w-full border rounded px-2 py-1"
        />
      </div>

      <div>
        <label className="block text-sm font-medium">Description</label>
        <textarea
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          className="w-full border rounded px-2 py-1"
        />
      </div>

      <div>
        <label className="block text-sm font-medium">Category</label>
        <input
          value={category}
          onChange={(e) => setCategory(e.target.value)}
          className="w-full border rounded px-2 py-1"
        />
      </div>

      <div>
        <label className="block text-sm font-medium">3D Model File (.glb/.gltf)</label>
        <input
          type="file"
          accept=".glb,.gltf"
          onChange={(e) => setFile(e.target.files?.[0] || null)}
        />
      </div>

      <button
        type="submit"
        disabled={loading}
        className="w-full py-2 rounded-xl bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50"
      >
        {loading ? "Saving…" : "Save Showroom"}
      </button>
    </form>
  );
}