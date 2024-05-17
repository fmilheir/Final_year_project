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
        const certificateContent = `
          -----BEGIN CERTIFICATE-----
          MIIDejCCAmICCQD5hO/X9OgMezANBgkqhkiG9w0BAQsFADB/MQswCQYDVQQGEwJV
          UzESMBAGA1UECAwJWW91clN0YXRlMREwDwYDVQQHDAhZb3VyQ2l0eTEZMBcGA1UE
          CgwQWW91ck9yZ2FuaXphdGlvbjERMA8GA1UECwwIWW91clVuaXQxGzAZBgNVBAMM
          End3dy55b3VyZG9tYWluLmNvbTAeFw0yNDA0MDYyMjA4MDlaFw0yNTA0MDYyMjA4
          MDlaMH8xCzAJBgNVBAYTAlVTMRIwEAYDVQQIDAlZb3VyU3RhdGUxETAPBgNVBAcM
          CFlvdXJDaXR5MRkwFwYDVQQKDBBZb3VyT3JnYW5pemF0aW9uMREwDwYDVQQLDAhZ
          b3VyVW5pdDEbMBkGA1UEAwwSd3d3LnlvdXJkb21haW4uY29tMIIBIjANBgkqhkiG
          9w0BAQEFAAOCAQ8AMIIBCgKCAQEAx4OKy0EpbYHIKo4CfsP8pwQl24RJmy0wI362
          eHQSPQy9zeyLuV8HIvQPodilkiHmYFufA/J6gx/S9asmw424xMAgaNk3CJzStqT1
          YzwHZnd33hfVjspA4YwbdYC3nTsqk7RXrF87rE46NeuHjbcyy4mL3gY34gcZSYvz
          bLXBTax96Qt9xEzdPu75RVCXmdKJtq7OGSe9SXWPT3ctWEuN1+bUIptg09YpK+QR
          ofI5hYfS0bJYzcWWZGmcgITBhjZ1HzZ2Im0J0n2iumq3zD7P+OoddmX8sry3FP0P
          4Su3DzA5X9ytshpfJOcYhv0rzeicGIuZf/EyPQbLOBOnLwUA1QIDAQABMA0GCSqG
          SIb3DQEBCwUAA4IBAQALlBTSLOeeETMWN4ywjqoz4I/Q0xgNrDhaJdgUxj/430Bz
          ys7jssup7lUXnzWuwk4W2zxDEIIYH6KHQhaS9uSl1bcI7ikaeOjMJC2H3H/9n7qv
          SjzCG6cGE5Tgidx5yTc8rej9aXEI2g+/FPy6IKfynHK2/spkZTFdcVkHN79Fhcu5
          sqQA8WLrz+YG0miUrIcmvuLMCQ3S6cH68R1fwYvKl7EqPgtw3G/cpuh8Vj/emi5r
          GxQKCyti2er4yuEAQP2CQEUuWoongEVySDm37zTYaD3FMnpPrmb7cRSaS2wFEA/d
          EWYVTprbpMXtkCctTLv7+SnAqio3A9MKblrNk4Rv
          -----END CERTIFICATE-----
        `;

        const response = await fetch('https://localhost:3030/tmf-api/Incident/v4/incident', {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
          },
        });

        if (!response.ok) {
          throw new Error('Failed to fetch tickets');
        }

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
