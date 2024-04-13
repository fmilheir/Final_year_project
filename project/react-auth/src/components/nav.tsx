import React from 'react';
import { Link } from 'react-router-dom';
import { Navbar, Container, Nav, NavDropdown } from 'react-bootstrap';
import logo from '../images/solent_logo.png';
import profileIcon from '../images/undraw_profile.svg';
import '../css/nav.css'; // Import custom CSS styles

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
    <Navbar expand="lg" bg="light" variant="light">
      <Container>
        <Navbar.Brand as={Link} to="/">
          <img src={logo} alt="Company Logo" className="nav-logo" width="150" height="auto" />
        </Navbar.Brand>
        <Navbar.Toggle aria-controls="navbarNav" />
        <Navbar.Collapse id="navbarNav">
          <Nav className="me-auto">
            {isAuthenticated && (
              <>
                <Nav.Link as={Link} to="/home" className="btn btn-primary me-2">Home</Nav.Link>
                <Nav.Link as={Link} to="/dashboard" className="btn btn-primary me-2">Dashboard</Nav.Link>
                <Nav.Link as={Link} to="/chatbot" className="btn btn-primary me-2">chatbot</Nav.Link>
                <Nav.Link as={Link} to="/create_ticket" className="btn btn-primary me-2">Create Ticket</Nav.Link>
                <Nav.Link as={Link} to="/admin_panel" className="btn btn-primary me-2">Admin Panel</Nav.Link>
              </>
            )}
          </Nav>
          <Nav>
          {isAuthenticated ? (
              <NavDropdown title={<img src={profileIcon} alt="Profile" className="profile-icon" width="32" height="32" />} id="profile-dropdown">
                <NavDropdown.Item as={Link} to="/profile">Profile</NavDropdown.Item>
                <NavDropdown.Divider />
                {/* Call handleLogout when "Sign out" is clicked */}
                <NavDropdown.Item onClick={handleLogout}>Sign out</NavDropdown.Item>
              </NavDropdown>
            ) : (
              <div className="d-flex">
                <Link to="/signup" className="btn btn-primary me-2">Sign-up</Link>
                <Link to="/login_user" className="btn btn-primary me-2">Login</Link>
                <Link to="/login_admin" className="btn btn-primary me-2">Admin</Link>
              </div>
            )}
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
    </header>
  );
};

export default Navigation;
