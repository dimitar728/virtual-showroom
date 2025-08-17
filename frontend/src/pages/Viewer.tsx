import ModelViewer from "../components/ModelViewer";

// pass a real URL from your API (e.g., /uploads/your.glb)
// or use a query param like ?model=https://.../model.glb
export default function Viewer() {
  const urlParam = new URLSearchParams(window.location.search);
  const modelUrl = urlParam.get("model") || "/models/dummy.glb";

  return (
    <div className="w-full h-full">
      <ModelViewer modelUrl={modelUrl} />
    </div>
  );
}