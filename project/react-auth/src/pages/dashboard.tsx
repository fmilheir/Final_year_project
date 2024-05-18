import React, { useEffect, useState } from 'react';
import { Link, Navigate } from 'react-router-dom'; // Import Link from react-router-dom
import { Box, Button, Card, CardContent, Typography, CircularProgress } from '@mui/material';
import '../css/nav.css';
import axios from 'axios';
import HttpsProxyAgent from 'https-proxy-agent';
import fs from 'fs';


interface Ticket {
    ID: number;
    name: string;
    priority: string;
    IncidentDetail: string;
    state: string;
}


const Dashboard = () => {
    const [tickets, setTickets] = useState<Ticket[]>([]); 
    const [isLoading, setIsLoading] = useState(true);
    const [redirect, setRedirect] = useState(false);

  
  
  useEffect(() => {
    const fetchTickets = async () => {
      try {
        const response = await fetch('/tmf-api/Incident/v4/incident', {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
          },
        });
        const data = await response.json();
        console.log('Tickets:', data);
        setTickets(data);
      } catch (error) {
        console.error('Error fetching tickets:', error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchTickets();
  }, []);
  
    const handleCreateTicket = () => {
        setRedirect(true);
    };

    if (redirect) {
        return <Navigate to="/create_ticket" />;
    }

    return (
        <div className="container mt-4">
          <div className="row">
            <div className="col">
              <h2 className="mb-4">Dashboard</h2>
              <button className="btn btn-primary mb-4" onClick={handleCreateTicket}>
                <i className="fas fa-plus me-2"></i>
                Create New Ticket
              </button>
            </div>
          </div>
          {isLoading ? (
            <div className="d-flex justify-content-center align-items-center" style={{ height: '50vh' }}>
              <div className="spinner-border" role="status">
                <span className="visually-hidden">Loading...</span>
              </div>
            </div>
          ) : (
            <div className="row row-cols-1 row-cols-md-2 row-cols-lg-3 g-4">
              {tickets.length ? (
                tickets.map((ticket) => (
                  <div className="col" key={ticket.ID}>
                    <div className="card h-100">
                      <div className="card-body">
                        <h5 className="card-title">{ticket.name}</h5>
                        <h6 className="card-subtitle mb-2 text-muted">Ticket ID: {ticket.ID}</h6>
                        <p className="card-text">{ticket.IncidentDetail}</p>
                        <p className="card-text">
                          <small className="text-muted">Priority: {ticket.priority}</small>
                        </p>
                        <Link to={`/edit_ticket/${ticket.ID}`} className="btn btn-outline-primary btn-sm">
                          Edit
                        </Link>
                      </div>
                    </div>
                  </div>
                ))
              ) : (
                <div className="col">
                  <p>No tickets available.</p>
                </div>
              )}
            </div>
          )}
        </div>
      );
};

export default Dashboard;
