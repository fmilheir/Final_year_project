import React from 'react';
import { Link } from 'react-router-dom';
import { Navbar, Container, Nav, Button, NavDropdown } from 'react-bootstrap';
import logo from '../images/solent_logo.png';
import profileIcon from '../images/undraw_profile.svg';
import '../css/nav.css';

interface Props {
  isAuthenticated: boolean;
  firstName: string;
  role: boolean;
}

const handleLogout = async () => {
  try {
    // Make a request to the logout endpoint
    const response = await fetch('http://localhost:8080/api/logout', {
      method: 'POST',
      credentials: 'include', // Include cookies in the request
    });
    if (!response.ok) {
      throw new Error('Failed to logout');
    }
    // Redirect the user to the login page after successful logout
    window.location.href = '/login';
    window.location.reload();
  } catch (error) {
    console.error('Error logging out:', error);
  }
};

const Navigation: React.FC<Props> = ({ isAuthenticated, role }) => {
  return (
    <header>
      <Navbar bg="white" variant="white" expand="md">
        <Navbar.Brand as={Link} to="/">
          <img src={logo} alt="Company Logo" className="nav-logo" width="100" height="auto" />
        </Navbar.Brand>
        <Navbar.Toggle aria-controls="navbar-nav" />
        <Navbar.Collapse id="navbar-nav">
          <Nav className="ms-auto align-items-center">
            {isAuthenticated ? (
              <>
                {role ? (
                  <>
                    <Nav.Link as={Link} to="/admin_dashboard">Admin Dashboard</Nav.Link>
                    <Nav.Link as={Link} to="/admin_panel">Admin Panel</Nav.Link>
                  </>
                ) : (
                  <Nav.Link as={Link} to="/dashboard">Dashboard</Nav.Link>
                )}
                <Nav.Link as={Link} to="/chatbot">Chatbot</Nav.Link>
                <Nav.Link as={Link} to="/create_ticket">Create Ticket</Nav.Link>
                <NavDropdown
                  title={
                    <img src={profileIcon} alt="Profile" className="profile-icon" width="32" height="32" />
                  }
                  id="profile-dropdown"
                >
                  <NavDropdown.Item as={Link} to="/profile">Profile</NavDropdown.Item>
                  <NavDropdown.Divider />
                  <NavDropdown.Item onClick={handleLogout}>Sign out</NavDropdown.Item>
                </NavDropdown>
              </>
            ) : (
              <>
                <Nav.Link as={Link} to="/signup" className="btn btn-primary me-2">
                  Sign-up
                </Nav.Link>
                <Nav.Link as={Link} to="/login_user" className="btn btn-outline-primary me-2">
                  Login
                </Nav.Link>
                <Nav.Link as={Link} to="/login_admin" className="btn btn-outline-secondary">
                  Admin
                </Nav.Link>
              </>
            )}
          </Nav>
        </Navbar.Collapse>
      </Navbar>
    </header>
  );
};

export default Navigation;