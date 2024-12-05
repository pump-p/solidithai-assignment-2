import React, { useContext } from 'react';
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate
} from 'react-router-dom';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css'; // Import the CSS for Toastify
import Header from './components/Header';
import HomePage from './pages/HomePage';
import LoginPage from './pages/LoginPage';
import SignupPage from './pages/SignupPage';
import UsersPage from './pages/UsersPage';
import AuthContext from './context/AuthContext';
import LogsPage from './pages/LogsPage';
import WebSocketClient from './components/WebSocketClient';

const App = () => {
  const { isAuthenticated } = useContext(AuthContext);

  return (
    <Router>
      <Header />
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route
          path="/login"
          element={!isAuthenticated ? <LoginPage /> : <Navigate to="/users" />}
        />
        <Route
          path="/signup"
          element={!isAuthenticated ? <SignupPage /> : <Navigate to="/users" />}
        />
        <Route
          path="/users"
          element={isAuthenticated ? <UsersPage /> : <Navigate to="/login" />}
        />
        <Route
          path="/logs"
          element={isAuthenticated ? <LogsPage /> : <Navigate to="/login" />}
        />
        <Route
          path="/websocket"
          element={
            isAuthenticated ? <WebSocketClient /> : <Navigate to="/login" />
          }
        />
      </Routes>
      <ToastContainer position="top-right" autoClose={2000} />
    </Router>
  );
};

export default App;
