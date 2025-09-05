import React, { useEffect, useRef } from 'react';
import { useParams } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { fetchShowroomById } from '../store/showroomSlice';
import * as THREE from 'three';
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls';

const ShowroomViewer = () => {
  const { id } = useParams();
  const dispatch = useDispatch();
  const { selectedShowroom, loading, error } = useSelector((state) => state.showroom);
  const mountRef = useRef();

  useEffect(() => {
    dispatch(fetchShowroomById(id));
  }, [dispatch, id]);

  useEffect(() => {
    if (!selectedShowroom || !selectedShowroom.model_path) return;
    let renderer, scene, camera, controls, loader, animationId;
    const mountNode = mountRef.current;
    const width = mountNode.clientWidth;
    const height = mountNode.clientHeight;
    renderer = new THREE.WebGLRenderer({ antialias: true });
    renderer.setSize(width, height);
    mountNode.appendChild(renderer.domElement);
    scene = new THREE.Scene();
    camera = new THREE.PerspectiveCamera(75, width / height, 0.1, 1000);
    camera.position.set(0, 2, 5);
    controls = new OrbitControls(camera, renderer.domElement);
    controls.enableDamping = true;
    loader = new GLTFLoader();
    loader.load(selectedShowroom.model_path, (gltf) => {
      scene.add(gltf.scene);
    });
    const light = new THREE.AmbientLight(0xffffff, 1);
    scene.add(light);
    const animate = () => {
      controls.update();
      renderer.render(scene, camera);
      animationId = requestAnimationFrame(animate);
    };
    animate();
    return () => {
      cancelAnimationFrame(animationId);
      renderer.dispose();
      if (mountNode && renderer.domElement.parentNode === mountNode) {
        mountNode.removeChild(renderer.domElement);
      }
    };
  }, [selectedShowroom]);

  if (loading) return <div>Loading 3D showroom...</div>;
  if (error) return <div className="text-red-500">{error}</div>;
  if (!selectedShowroom) return <div>No showroom found.</div>;

  return (
    <div className="max-w-4xl mx-auto mt-8">
      <h2 className="text-2xl font-bold mb-4">{selectedShowroom.name}</h2>
      <div className="mb-4">{selectedShowroom.description}</div>
      <div ref={mountRef} style={{ width: '100%', height: '500px', background: '#222' }} />
    </div>
  );
};

export default ShowroomViewer;
