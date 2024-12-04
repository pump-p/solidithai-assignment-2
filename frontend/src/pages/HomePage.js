import React, { useContext } from 'react';
import { Link } from 'react-router-dom';
import AuthContext from '../context/AuthContext';

const HomePage = () => {
  const { isAuthenticated, logout } = useContext(AuthContext); // Call useContext at the top level

  return (
    <div className="container mx-auto mt-8 text-center">
      <h1 className="text-3xl font-bold mb-8">Welcome to Homepage</h1>
      {!isAuthenticated ? (
        // Show Login/Signup options when logged out
        <div className="space-x-4">
          <Link
            to="/login"
            className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
          >
            Login
          </Link>
          <Link
            to="/signup"
            className="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600"
          >
            Signup
          </Link>
        </div>
      ) : (
        // Show Manage Users and Logout when logged in
        <div className="space-x-4">
          <Link
            to="/users"
            className="bg-yellow-500 text-white px-4 py-2 rounded hover:bg-yellow-600"
          >
            Manage Users
          </Link>
          <button
            onClick={logout} // Use logout directly
            className="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
          >
            Logout
          </button>
        </div>
      )}
    </div>
  );
};

export default HomePage;
