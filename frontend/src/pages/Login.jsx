import React, { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { loginUser } from '../store/authSlice';
import { useNavigate } from 'react-router-dom';

const Login = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const { loading, error, user } = useSelector((state) => state.auth);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [errorMsg, setErrorMsg] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const res = await dispatch(loginUser({ email, password }));
      if (res.token) {
        localStorage.setItem('token', res.token);
        navigate('/user');
      } else {
        setErrorMsg(res.error || 'Login failed');
      }
  };

  useEffect(() => {
    if (user && user.role === 'admin') {
      navigate('/admin');
    } else if (user) {
      navigate('/user');
    }
  }, [user, navigate]);

    return (
      <div className="max-w-md mx-auto mt-16 p-8 bg-white rounded-xl shadow-lg">
        <h2 className="text-2xl font-bold mb-6 text-blue-700 text-center">Login</h2>
        <form onSubmit={handleSubmit} className="space-y-6">
          <div>
            <label className="block mb-2 text-gray-700 font-semibold">Email</label>
            <input
              type="email"
              className="w-full border border-gray-300 px-4 py-2 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              value={email}
              onChange={e => setEmail(e.target.value)}
              placeholder="Email"
              required
            />
          </div>
          <div>
            <label className="block mb-2 text-gray-700 font-semibold">Password</label>
            <input
              type="password"
              className="w-full border border-gray-300 px-4 py-2 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              value={password}
              onChange={e => setPassword(e.target.value)}
              placeholder="Password"
              required
            />
          </div>
    {(error || errorMsg) && <div className="text-red-500 text-center">{error || errorMsg}</div>}
          <button
            type="submit"
            className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 font-semibold transition"
            disabled={loading}
          >
            {loading ? 'Logging in...' : 'Login'}
          </button>
        </form>
      </div>
    );
};

export default Login;
