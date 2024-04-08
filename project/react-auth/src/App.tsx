import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate} from 'react-router-dom';
import Navigation from './components/nav';
import Home from './pages/home';
import LoginAdmin from './pages/login_admin';
import Signup from './pages/signup';
import LoginUser from './pages/login_user';
import Dashboard from './pages/dashboard';
import Chatbot from './pages/chatbot';
import CreateTicketForm from './pages/create_ticket';
import './App.css';

const App: React.FC = () => {
  const [userData, setUserData] = useState<any>(null); // State to store user data

  useEffect(() => {
    // Fetch user data when the component mounts
    fetchUserData();
  }, []);

  const fetchUserData = async () => {
    try {
      // Fetch user data from the backend
      const response = await fetch('http://localhost:8080/api/user', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          // Add any authentication headers if required
        },
      });
      if (!response.ok) {
        throw new Error('Failed to fetch user data');
      }
      // Parse response JSON
      const userData = await response.json();
      // Set user data to state
      setUserData(userData);
    } catch (error) {
      console.error('Error fetching user data:', error);
    }
  };

  const isAuthenticated = false ;// userData !== null; // Check if user data is fetched

  return (
    <body>
      <div className="App">
        <Router>
        <Routes>
            <Navigation isAuthenticated={isAuthenticated} />
              <Route path="/" element={<Navigate replace to="/home" />} />
              <Route path="/home" element={<Home />} />
              <Route path="/login_admin" element={<LoginAdmin />} />
              <Route path="/login_user" element={<LoginUser />} />
              <Route path="/signup" element={<Signup />} />
              {isAuthenticated ? (
                <>
                  <Route path="/dashboard" element={<Dashboard />} />
                  <Route path="/chatbot" element={<Chatbot />} />
                  <Route path="/create_ticket" element={<CreateTicketForm />} />
                </>
              ) : (
                <Route path="*" element={<Navigate replace to="/home" />} />
              )}
              <Route path="*" element={<Navigate replace to="/home" />} />
            </Routes>
        </Router>
      </div>
    </body>
  );
};

export default App;
