import React, { useRef, useEffect } from "react";
import { useParams } from "react-router-dom";
import * as THREE from "three";

const Room3D = () => {
  const mountRef = useRef(null);
  const rendererRef = useRef(null);
  const { id } = useParams();

  useEffect(() => {
    const mountNode = mountRef.current;
    if (!mountNode) return;

    const scene = new THREE.Scene();
    scene.background = new THREE.Color(0xf0f0f0);

    const camera = new THREE.PerspectiveCamera(
      75,
      mountNode.clientWidth / mountNode.clientHeight,
      0.1,
      1000
    );
    camera.position.z = 5;

    const renderer = new THREE.WebGLRenderer();
    renderer.setSize(mountNode.clientWidth, mountNode.clientHeight);
    mountNode.appendChild(renderer.domElement);
    rendererRef.current = renderer;

    const geometry = new THREE.BoxGeometry(4, 2, 4);
    const material = new THREE.MeshBasicMaterial({ color: 0x8ecae6, wireframe: true });
    const room = new THREE.Mesh(geometry, material);
    scene.add(room);

    const animate = () => {
      room.rotation.y += 0.01;
      renderer.render(scene, camera);
      requestAnimationFrame(animate);
    };
    animate();

    return () => {
      if (mountNode && rendererRef.current) {
        mountNode.removeChild(rendererRef.current.domElement);
      }
    };
  }, [id]);

  return (
    <div>
      <h2 className="text-xl font-bold mb-2">3D Explorer for Showroom {id}</h2>
      <div
        ref={mountRef}
        style={{ width: "100vw", height: "80vh", border: "1px solid #ccc" }}
      />
    </div>
  );
};

export default Room3D;