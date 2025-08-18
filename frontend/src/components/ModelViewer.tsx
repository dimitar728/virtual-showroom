import { useEffect, useRef, useState, useMemo } from "react";
import * as THREE from "three";
import { GLTFLoader } from "three/examples/jsm/loaders/GLTFLoader.js";
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls.js";
import { PointerLockControls } from "three/examples/jsm/controls/PointerLockControls.js";
import HotspotGizmo from "./HotspotGizmo";
import type { Hotspot } from "../types";
import { fetchHotspots, createHotspot } from "../services/hotspotService";

type Props = {
  showroomId: string;
  modelUrl: string;
  initialControl?: "orbit" | "first";
  background?: string;
  enableHotspotEditor?: boolean; // admin toggle
};

export default function ModelViewer({
  showroomId,
  modelUrl,
  initialControl = "orbit",
  background = "#101114",
  enableHotspotEditor = false,
}: Props) {
  const containerRef = useRef<HTMLDivElement | null>(null);
  const [mode, setMode] = useState<"orbit" | "first">(initialControl);

  const rendererRef = useRef<THREE.WebGLRenderer | null>(null);
  const sceneRef = useRef<THREE.Scene | null>(null);
  const cameraRef = useRef<THREE.PerspectiveCamera | null>(null);
  const orbitRef = useRef<OrbitControls | null>(null);
  const pointerRef = useRef<PointerLockControls | null>(null);

  const [hotspots, setHotspots] = useState<Hotspot[]>([]);
  const [selected, setSelected] = useState<Hotspot | null>(null);

  // init three
  useEffect(() => {
    if (!containerRef.current) return;
    const W = containerRef.current.clientWidth;
    const H = containerRef.current.clientHeight;

    const scene = new THREE.Scene();
    scene.background = new THREE.Color(background);
    sceneRef.current = scene;

    const camera = new THREE.PerspectiveCamera(60, W / H, 0.1, 1000);
    camera.position.set(2, 1.6, 3);
    cameraRef.current = camera;

    const renderer = new THREE.WebGLRenderer({ antialias: true });
    renderer.setSize(W, H);
    renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
    rendererRef.current = renderer;
    containerRef.current.appendChild(renderer.domElement);

    const light = new THREE.HemisphereLight(0xffffff, 0x222233, 1.0);
    scene.add(light);

    const controls = new OrbitControls(camera, renderer.domElement);
    controls.enableDamping = true;
    orbitRef.current = controls;

    const plc = new PointerLockControls(camera, renderer.domElement);
    pointerRef.current = plc;

    // load model
    const loader = new GLTFLoader();
    loader.load(
      modelUrl,
      (gltf) => {
        gltf.scene.traverse((o) => { o.frustumCulled = false; });
        scene.add(gltf.scene);
      },
      undefined,
      (err) => console.error("model load error", err)
    );

    const onResize = () => {
      if (!containerRef.current || !rendererRef.current || !cameraRef.current) return;
      const w = containerRef.current.clientWidth;
      const h = containerRef.current.clientHeight;
      rendererRef.current.setSize(w, h);
      cameraRef.current.aspect = w / h;
      cameraRef.current.updateProjectionMatrix();
    };
    window.addEventListener("resize", onResize);

    const clock = new THREE.Clock();
    const animate = () => {
      requestAnimationFrame(animate);
      orbitRef.current?.update();
      renderer.render(scene, camera);
    };
    animate();

    return () => {
      window.removeEventListener("resize", onResize);
      renderer.dispose();
      containerRef.current?.removeChild(renderer.domElement);
    };
  }, [modelUrl, background]);

  // fetch hotspots
  useEffect(() => {
    fetchHotspots(showroomId).then((hs) => setHotspots(hs || [])).catch(console.error);
  }, [showroomId]);

  // editor: create hotspot with Ctrl+Click on canvas
  useEffect(() => {
    if (!enableHotspotEditor) return;
    const canvas = rendererRef.current?.domElement;
    if (!canvas || !sceneRef.current || !cameraRef.current) return;

    const raycaster = new THREE.Raycaster();
    const mouse = new THREE.Vector2();

    const onClick = async (e: MouseEvent) => {
      if (!e.ctrlKey) return; // hold Ctrl to place
      const rect = canvas.getBoundingClientRect();
      mouse.x = ((e.clientX - rect.left) / rect.width) * 2 - 1;
      mouse.y = -((e.clientY - rect.top) / rect.height) * 2 + 1;

      raycaster.setFromCamera(mouse, cameraRef.current!);
      const intersects = raycaster.intersectObjects(sceneRef.current!.children, true);
      if (intersects.length === 0) return;

      const p = intersects[0].point;
      const label = prompt("Hotspot label?");
      if (!label) return;
      const description = prompt("Description (optional)") || "";

      try {
        const created = await createHotspot(showroomId, {
          label,
          description,
          pos_x: p.x,
          pos_y: p.y,
          pos_z: p.z,
        });
        setHotspots((prev) => [...prev, created]);
      } catch (err) {
        console.error(err);
        alert("Failed to create hotspot");
      }
    };

    canvas.addEventListener("click", onClick);
    return () => canvas.removeEventListener("click", onClick);
  }, [enableHotspotEditor, showroomId]);

  // control toggles
  const togglePointerLock = () => {
    if (!pointerRef.current) return;
    if (document.pointerLockElement) {
      document.exitPointerLock();
    } else {
      pointerRef.current.lock();
    }
  };

  const gizmos = useMemo(() => {
    if (!cameraRef.current || !rendererRef.current) return null;
    return hotspots.map((h) => {
      const pos = new THREE.Vector3(h.pos_x, h.pos_y, h.pos_z);
      return (
        <HotspotGizmo
          key={h.id}
          camera={cameraRef.current}
          rendererDom={rendererRef.current!.domElement}
          worldPosition={pos}
          label={h.label}
          onClick={() => setSelected(h)}
        />
      );
    });
  }, [hotspots]);

  return (
    <div ref={containerRef} className="relative w-full h-full">
      {/* top-right controls */}
      <div className="absolute right-3 top-3 z-20 flex items-center gap-2">
        <button
          onClick={() => setMode((m) => (m === "orbit" ? "first" : "orbit"))}
          className="px-3 py-1 rounded bg-white/90 hover:bg-white shadow text-sm"
          title="Toggle orbit/first person"
        >
          {mode === "orbit" ? "First-person" : "Orbit"}
        </button>
        {mode === "first" && (
          <button
            onClick={togglePointerLock}
            className="px-3 py-1 rounded bg-white/90 hover:bg-white shadow text-sm"
            title="Toggle pointer lock"
          >
            Look
          </button>
        )}
        {enableHotspotEditor && (
          <span className="text-xs text-white/90 bg-black/40 rounded px-2 py-1">
            Editor: Ctrl+Click to place hotspot
          </span>
        )}
      </div>

      {/* markers */}
      <div className="pointer-events-none absolute inset-0">{gizmos}</div>

      {/* info card */}
      {selected && (
        <div className="absolute left-3 bottom-3 z-30 max-w-sm bg-white rounded-xl shadow p-3">
          <div className="flex items-start justify-between">
            <h4 className="font-semibold">{selected.label}</h4>
            <button className="text-sm" onClick={() => setSelected(null)}>✕</button>
          </div>
          {selected.media_url && (
            <img src={selected.media_url} alt="" className="w-full h-auto rounded mt-2" />
          )}
          {selected.description && (
            <p className="text-sm text-gray-700 mt-2">{selected.description}</p>
          )}
          {selected.link_url && (
            <a className="text-sm text-blue-600 underline mt-2 inline-block" href={selected.link_url} target="_blank">
              Learn more
            </a>
          )}
        </div>
      )}

      {/* first-person hint */}
      {mode === "first" && (
        <div className="absolute bottom-3 right-3 text-xs text-white/80 z-10 select-none">
          Click canvas to look around • Move: WASD/Arrows • Esc to release
        </div>
      )}
    </div>
  );
}