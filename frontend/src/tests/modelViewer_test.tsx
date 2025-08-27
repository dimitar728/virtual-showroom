import { render, screen } from '@testing-library/react';
import ModelViewer from '../components/ModelViewer';

describe('ModelViewer', () => {
  it('renders without crashing', () => {
    render(<ModelViewer showroomId="1" modelUrl="/models/dummy.glb" />);
    expect(screen.getByText(/loading/i)).toBeInTheDocument();
  });

  it('accepts props and renders container', () => {
    render(<ModelViewer showroomId="2" modelUrl="/models/other.glb" background="#fff" />);
    // You can expand this test to check for canvas or three.js elements
  });
});
