import { Link } from "react-router-dom";

export default function Login() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-2xl shadow-md w-96">
        <h2 className="text-2xl font-semibold mb-6">Login</h2>
        <form className="space-y-4">
          <input type="email" placeholder="Email"
            className="w-full px-4 py-2 border rounded-lg focus:ring focus:ring-blue-300" />
          <input type="password" placeholder="Password"
            className="w-full px-4 py-2 border rounded-lg focus:ring focus:ring-blue-300" />
          <button type="submit"
            className="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700">
            Login
          </button>
        </form>
        <p className="mt-4 text-sm">
          Don’t have an account? <Link to="/register" className="text-blue-600">Register</Link>
        </p>
      </div>
    </div>
  );
}