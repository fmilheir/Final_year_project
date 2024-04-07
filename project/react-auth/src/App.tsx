import React from 'react';
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
  const isAuthenticated = false;


  return (
    <body>
      <div className="App">
        <Router>
          <Navigation isAuthenticated={isAuthenticated} />
          <div className="content">
            <Routes>
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
          </div>
        </Router>
      </div>
    </body>
  );
};

export default App;
