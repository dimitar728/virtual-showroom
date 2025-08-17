import { useEffect, useRef, useState } from "react";
import * as THREE from "three";
import { GLTFLoader } from "three/examples/jsm/loaders/GLTFLoader.js";
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls.js";
import { PointerLockControls } from "three/examples/jsm/controls/PointerLockControls.js";

export default function ModelViewer({
  modelUrl,
  initialControl = "orbit", // "orbit" | "first"
  background = "#101114",
}) {
  const containerRef = useRef(null);
  const [mode, setMode] = useState(initialControl);

  // three refs
  const rendererRef = useRef();
  const sceneRef = useRef();
  const cameraRef = useRef();
  const orbitRef = useRef();
  const fpRef = useRef(); // pointer-lock controls
  const clockRef = useRef(new THREE.Clock());

  // movement (first-person)
  const move = useRef({ f: false, b: false, l: false, r: false });
  const velocity = useRef(new THREE.Vector3());
  const direction = useRef(new THREE.Vector3());

  useEffect(() => {
    const container = containerRef.current;
    const width = container.clientWidth;
    const height = container.clientHeight;

    // renderer
    const renderer = new THREE.WebGLRenderer({ antialias: true });
    renderer.setSize(width, height);
    renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
    renderer.outputColorSpace = THREE.SRGBColorSpace;
    renderer.shadowMap.enabled = true;
    rendererRef.current = renderer;
    container.appendChild(renderer.domElement);

    // scene
    const scene = new THREE.Scene();
    scene.background = new THREE.Color(background);
    sceneRef.current = scene;

    // camera
    const camera = new THREE.PerspectiveCamera(60, width / height, 0.1, 2000);
    camera.position.set(4, 2, 6);
    cameraRef.current = camera;

    // lights
    const hemi = new THREE.HemisphereLight(0xffffff, 0x444444, 0.8);
    hemi.position.set(0, 50, 0);
    scene.add(hemi);

    const dir = new THREE.DirectionalLight(0xffffff, 0.8);
    dir.position.set(5, 10, 7.5);
    dir.castShadow = true;
    dir.shadow.camera.near = 0.1;
    dir.shadow.camera.far = 100;
    scene.add(dir);

    // helpers (grid/floor)
    const grid = new THREE.GridHelper(100, 100, 0x666666, 0x333333);
    grid.material.opacity = 0.35;
    grid.material.transparent = true;
    scene.add(grid);

    const floorGeo = new THREE.PlaneGeometry(200, 200);
    const floorMat = new THREE.MeshStandardMaterial({ color: 0x151515, metalness: 0.1, roughness: 0.9 });
    const floor = new THREE.Mesh(floorGeo, floorMat);
    floor.rotation.x = -Math.PI / 2;
    floor.receiveShadow = true;
    scene.add(floor);

    // controls
    const orbit = new OrbitControls(camera, renderer.domElement);
    orbit.enableDamping = true;
    orbitRef.current = orbit;

    const fp = new PointerLockControls(camera, renderer.domElement);
    fpRef.current = fp;

    // load model
    const loader = new GLTFLoader();
    loader.load(
      modelUrl,
      (gltf) => {
        const root = gltf.scene || gltf.scenes[0];
        root.traverse((obj) => {
          if (obj.isMesh) {
            obj.castShadow = true;
            obj.receiveShadow = true;
          }
        });
        scene.add(root);

        // auto-frame model for orbit
        const box = new THREE.Box3().setFromObject(root);
        const size = box.getSize(new THREE.Vector3()).length();
        const center = box.getCenter(new THREE.Vector3());
        const fitDist = Math.max(6, size * 1.2);
        orbit.target.copy(center);
        camera.position.copy(center).add(new THREE.Vector3(fitDist, fitDist * 0.6, fitDist));
        camera.lookAt(center);
      },
      undefined,
      (err) => console.error("GLTF load error:", err)
    );

    // resize
    const onResize = () => {
      const w = container.clientWidth;
      const h = container.clientHeight;
      renderer.setSize(w, h);
      camera.aspect = w / h;
      camera.updateProjectionMatrix();
    };
    window.addEventListener("resize", onResize);

    // keyboard (first-person)
    const onKeyDown = (e) => {
      switch (e.code) {
        case "KeyW": case "ArrowUp": move.current.f = true; break;
        case "KeyS": case "ArrowDown": move.current.b = true; break;
        case "KeyA": case "ArrowLeft": move.current.l = true; break;
        case "KeyD": case "ArrowRight": move.current.r = true; break;
      }
    };
    const onKeyUp = (e) => {
      switch (e.code) {
        case "KeyW": case "ArrowUp": move.current.f = false; break;
        case "KeyS": case "ArrowDown": move.current.b = false; break;
        case "KeyA": case "ArrowLeft": move.current.l = false; break;
        case "KeyD": case "ArrowRight": move.current.r = false; break;
      }
    };
    document.addEventListener("keydown", onKeyDown);
    document.addEventListener("keyup", onKeyUp);

    // render loop
    let raf;
    const animate = () => {
      raf = requestAnimationFrame(animate);
      const dt = Math.min(0.05, clockRef.current.getDelta());

      // orbit damping
      if (orbit.enabled) orbit.update();

      // first-person move
      if (fp.isLocked) {
        velocity.current.x -= velocity.current.x * 8.0 * dt;
        velocity.current.z -= velocity.current.z * 8.0 * dt;

        direction.current.set(
          Number(move.current.r) - Number(move.current.l),
          0,
          Number(move.current.b) - Number(move.current.f)
        ).normalize();

        const speed = 10; // units/sec
        if (move.current.f || move.current.b) velocity.current.z -= direction.current.z * speed * dt;
        if (move.current.l || move.current.r) velocity.current.x -= direction.current.x * speed * dt;

        const cam = fp.getObject();
        cam.translateX(velocity.current.x * dt);
        cam.translateZ(velocity.current.z * dt);
      }

      renderer.render(scene, camera);
    };
    animate();

    // cleanup
    return () => {
      cancelAnimationFrame(raf);
      window.removeEventListener("resize", onResize);
      document.removeEventListener("keydown", onKeyDown);
      document.removeEventListener("keyup", onKeyUp);
      orbit.dispose();
      fp.dispose();
      renderer.dispose();
      container.removeChild(renderer.domElement);
      scene.traverse((obj) => {
        if (obj.geometry) obj.geometry.dispose();
        if (obj.material) {
          if (Array.isArray(obj.material)) obj.material.forEach((m) => m.dispose());
          else obj.material.dispose();
        }
      });
    };
  }, [modelUrl, background]);

  // toggle modes
  useEffect(() => {
    const orbit = orbitRef.current;
    const fp = fpRef.current;
    if (!orbit || !fp) return;

    if (mode === "orbit") {
      // exit pointer lock if active
      if (fp.isLocked) fp.unlock();
      orbit.enabled = true;
    } else {
      orbit.enabled = false;
      // click canvas to lock
      const el = rendererRef.current?.domElement;
      const lock = () => fp.lock();
      el?.addEventListener("click", lock);
      return () => el?.removeEventListener("click", lock);
    }
  }, [mode]);

  const handlePointerLockToggle = () => {
    const fp = fpRef.current;
    if (!fp) return;
    fp.isLocked ? fp.unlock() : fp.lock();
  };

  return (
    <div className="w-full h-full relative" ref={containerRef} style={{ minHeight: 480 }}>
      <div className="absolute top-3 left-3 flex gap-2 z-10">
        <button
          onClick={() => setMode("orbit")}
          className={`px-3 py-1 rounded-md text-sm ${mode === "orbit" ? "bg-white/90" : "bg-white/60"}`}
          title="Orbit mode"
        >
          Orbit
        </button>
        <button
          onClick={() => setMode("first")}
          className={`px-3 py-1 rounded-md text-sm ${mode === "first" ? "bg-white/90" : "bg-white/60"}`}
          title="First-person (click canvas to look, WASD to move)"
        >
          First-Person
        </button>
        {mode === "first" && (
          <button
            onClick={handlePointerLockToggle}
            className="px-3 py-1 rounded-md text-sm bg-white/60"
            title="Toggle pointer lock"
          >
            Toggle Look
          </button>
        )}
      </div>

      {mode === "first" && (
        <div className="absolute bottom-3 left-3 text-xs text-white/80 z-10 select-none">
          Click canvas to look around • Move: WASD / Arrows • Esc to release
        </div>
      )}
    </div>
  );
}