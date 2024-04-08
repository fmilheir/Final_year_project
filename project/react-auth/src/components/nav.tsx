import React from 'react';
import { Link } from 'react-router-dom';
import { Navbar, Container, Nav, NavDropdown } from 'react-bootstrap';
import logo from '../images/solent_logo.png';
import profileIcon from '../images/undraw_profile.svg';
import '../css/nav.css'; // Import custom CSS styles

interface Props {
  isAuthenticated: boolean;
}

const Navigation: React.FC<Props> = ({ isAuthenticated }) => {
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
                <Nav.Link as={Link} to="/my_stocks" className="btn btn-primary me-2">Predictions</Nav.Link>
                <Nav.Link as={Link} to="/news" className="btn btn-primary me-2">News</Nav.Link>
              </>
            )}
          </Nav>
          <Nav>
            {isAuthenticated ? (
              <NavDropdown title={<img src={profileIcon} alt="Profile" className="profile-icon" width="32" height="32" />} id="profile-dropdown">
                <NavDropdown.Item as={Link} to="/profile">Profile</NavDropdown.Item>
                <NavDropdown.Divider />
                <NavDropdown.Item as={Link} to="/logout">Sign out</NavDropdown.Item>
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
