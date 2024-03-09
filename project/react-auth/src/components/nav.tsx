import React from "react";
import 'bootstrap/dist/css/bootstrap.min.css';

const Nav = ({ IsAuthenticated }: { IsAuthenticated: boolean }) => {
  return (
    <header className="navbar navbar-expand-md navbar-light bg-light">
      <div className="container">
        <a href="/" className="navbar-brand">
          <img src="../images/solent_logo.png" width="150" height="70" alt="Logo image" />
        </a>
        <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav"
          aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
          <span className="navbar-toggler-icon"></span>
        </button>
        <div className="collapse navbar-collapse" id="navbarNav">
          <ul className="navbar-nav me-auto mb-2 mb-md-0">
            {IsAuthenticated ? (
              <>
                <li className="nav-item">
                  <a className="nav-link" href="/home">Home</a>
                </li>
                <li className="nav-item">
                  <a className="nav-link" href="/my_stocks">Predictions</a>
                </li>
                <li className="nav-item">
                  <a className="nav-link" href="/news">News</a>
                </li>
              </>
            ) : null}
          </ul>
          <div className="d-flex">
            {IsAuthenticated ? (
              <div className="dropdown">
                <button className="btn btn-secondary dropdown-toggle" type="button" id="dropdownMenuButton"
                  data-bs-toggle="dropdown" aria-expanded="false">
                  <img src="/static/images/undraw_profile.svg" alt="Profile" className="rounded-circle" width="32" height="32" />
                </button>
                <ul className="dropdown-menu dropdown-menu-end" aria-labelledby="dropdownMenuButton">
                  <li><a className="dropdown-item" href="/profile">Profile</a></li>
                  <li><hr className="dropdown-divider" /></li>
                  <li><a className="dropdown-item" href="/logout">Sign out</a></li>
                </ul>
              </div>
            ) : (
              <div className="btn-group" role="group" aria-label="Auth Links">
                <a href="/" className="btn btn-outline-primary me-2">Home</a>
                <a href="/signup" className="btn btn-primary me-2">Sign-up</a>
                <a href="/login_user" className="btn btn-primary me-2">Login</a>
                <a href="/login_admin" className="btn btn-primary me-2">Admin</a>
              </div>
            )}
          </div>
        </div>
      </div>
    </header>
  );    
};

export default Nav;
