import { useEffect, useMemo, useRef } from "react";
import * as THREE from "three";

type GizmoProps = {
  camera: THREE.Camera | null;
  rendererDom: HTMLCanvasElement | null;
  worldPosition: THREE.Vector3;
  label: string;
  onClick?: () => void;
};

export default function HotspotGizmo({ camera, rendererDom, worldPosition, label, onClick }: GizmoProps) {
  const elRef = useRef<HTMLButtonElement | null>(null);

  const screenPos = useMemo(() => new THREE.Vector3(), []);
  useEffect(() => {
    if (!camera || !rendererDom || !elRef.current) return;

    const update = () => {
      screenPos.copy(worldPosition).project(camera);
      const x = (screenPos.x * 0.5 + 0.5) * rendererDom.clientWidth;
      const y = (-screenPos.y * 0.5 + 0.5) * rendererDom.clientHeight;
      const isBehind = screenPos.z > 1; // behind camera
      const el = elRef.current!;
      el.style.transform = `translate(-50%, -50%) translate(${x}px, ${y}px)`;
      el.style.opacity = isBehind ? "0" : "1";
      el.style.pointerEvents = isBehind ? "none" : "auto";
    };

    update();
    const onResize = () => update();
    const onFrame = () => update();

    const id = window.requestAnimationFrame(function raf() {
      update();
      window.requestAnimationFrame(raf);
    });

    window.addEventListener("resize", onResize);
    rendererDom.addEventListener("mousemove", onFrame);

    return () => {
      window.cancelAnimationFrame(id);
      window.removeEventListener("resize", onResize);
      rendererDom.removeEventListener("mousemove", onFrame);
    };
  }, [camera, rendererDom, screenPos, worldPosition]);

  return (
    <button
      ref={elRef}
      onClick={onClick}
      className="absolute z-20 px-2 py-1 rounded-full text-xs font-medium bg-white/90 hover:bg-white shadow"
      title={label}
      style={{ left: 0, top: 0 }}
    >
      {label}
    </button>
  );
}