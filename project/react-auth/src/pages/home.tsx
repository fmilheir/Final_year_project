import React from "react";
import { Link } from "react-router-dom";

const Main = () => {
  return (
    <main>
      <div>
        <div className="bg-white">
          <div className="row align-items-center">
            <div className="col-lg-6">
              <h1 className="display-4 fw-bold mb-4">Streamline Your Ticketing Process</h1>
              <p className="lead mb-4">
                Our intelligent chatbot system leverages advanced NLP techniques and industry standards to provide accurate and efficient ticket resolution.
              </p>
              <div className="d-grid gap-3 d-md-flex">
                <Link to="/signup" className="btn btn-primary btn-lg px-4">
                  Get Started
                </Link>
                <Link to="/login_admin" className="btn btn-outline-secondary btn-lg px-4">
                  Login
                </Link>
              </div>
            </div>
          </div>
        </div>
        <div className="bg-white py-5">
          <div className="row g-4">
            <div className="col-lg-4">
              <div className="card p-4">
                <div className="card-body text-center">
                  <i className="fas fa-ticket-alt fa-3x text-primary mb-3"></i>
                  <h5 className="card-title">Ticket Management</h5>
                  <p className="card-text">Efficiently manage and track tickets across departments.</p>
                </div>
              </div>
            </div>
            <div className="col-lg-4">
              <div className="card p-4">
                <div className="card-body text-center">
                  <i className="fas fa-robot fa-3x text-primary mb-3"></i>
                  <h5 className="card-title">Intelligent Chatbot</h5>
                  <p className="card-text">Automated ticket resolution with human-like responses.</p>
                </div>
              </div>
            </div>
            <div className="col-lg-4">
              <div className="card p-4">
                <div className="card-body text-center">
                  <i className="fas fa-chart-line fa-3x text-primary mb-3"></i>
                  <h5 className="card-title">Analytics & Reporting</h5>
                  <p className="card-text">Gain insights into ticket trends and performance metrics.</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  );
};

export default Main;