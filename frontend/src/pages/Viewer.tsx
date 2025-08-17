import ModelViewer from "../components/ModelViewer";
import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { apiFetch } from "../services/apiClient";

export default function Viewer() {
  const { id } = useParams(); // showroomId from route
  const [showroom, setShowroom] = useState<any>(null);
  const isAdmin = localStorage.getItem("role") === "admin"; // or your real RBAC

  useEffect(() => {
    apiFetch(`${import.meta.env.VITE_API_URL || "http://localhost:8080"}/api/showrooms/${id}`)
      .then(setShowroom)
      .catch(console.error);
  }, [id]);

  if (!showroom) return <div className="p-6">Loading…</div>;

  return (
    <div className="w-screen h-[calc(100vh-56px)]">
      <ModelViewer
        showroomId={id!}
        modelUrl={showroom.model_path}
        initialControl="orbit"
        enableHotspotEditor={isAdmin}
      />
    </div>
  );
}