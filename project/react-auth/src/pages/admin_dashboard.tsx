import React, { useEffect, useState } from 'react';
import { Link, Navigate } from 'react-router-dom';
import { Box, Button, Card, CardContent, Typography, CircularProgress, Modal, TextField, FormControl, InputLabel, Select, MenuItem } from '@mui/material';
import '../css/nav.css';
import axios from 'axios';
import HttpsProxyAgent from 'https-proxy-agent';
import fs from 'fs';

interface Ticket {
  ID: number;
  name: string;
  priority: string;
  incidentDetail: string;
  state: string;
  editedComment: string;
}

const Dashboard = () => {
  const [tickets, setTickets] = useState<Ticket[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [redirect, setRedirect] = useState(false);
  const [editTicket, setEditTicket] = useState<Ticket | null>(null);
  const [editedState, setEditedState] = useState('');
  const [editedComment, setEditedComment] = useState('');

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


  const getStateClass = (state:string) => {
    switch (state) {
      case 'raised':
        return 'bg-danger';
      case 'acknowledged':
        return 'bg-warning';
      case 'resolved':
        return 'bg-success';
      default:
        return '';
    }
  };
  
  const getStateBadgeClass = (state:string) => {
    switch (state) {
      case 'raised':
        return 'bg-danger text-white';
      case 'acknowledged':
        return 'bg-warning text-white';
      case 'resolved':
        return 'bg-success text-white';
      default:
        return '';
    }
  };

  const handleEditTicket = (ticket: Ticket) => {
    setEditTicket(ticket);
    setEditedState(ticket.state);
    setEditedComment('');
  };

  const handleUpdateTicket = async () => {
    if (editTicket) {
      try {
        const response = await fetch(`/tmf-api/Incident/v4/resolveIncident`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            incident: {
              ID: editTicket.ID.toString(),
              state: editedState,
              resolution: editedComment,
            },
          }),
        });
  
        if (!response.ok) {
          const errorData = await response.json();
          throw new Error(`Failed to update ticket: ${errorData.message}`);
        }
  
        setTickets((prevTickets) =>
          prevTickets.map((ticket) =>
            ticket.ID === editTicket.ID ? { ...ticket, state: editedState } : ticket
          )
        );
        setEditTicket(null);
      } catch (error) {
        console.error('Error updating ticket:', error);
      }
    }
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
            <i className="fas fa-plus me-2"></i> Create New Ticket
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
        <div className={`card h-100 ${getStateClass(ticket.state)}`}>
          <div className="card-body">
            <h5 className="card-title">{ticket.name}</h5>
            <h6 className="card-subtitle mb-2 text-muted">Ticket ID: {ticket.ID}</h6>
            <p className="card-text">{ticket.incidentDetail}</p>
            <p className="card-text">
              <small className="text-muted">Priority: {ticket.priority}</small>
            </p>
            {ticket.state === 'resolved' && (
              <p className="card-text">
                <strong>Resolution:</strong> {editedComment}
              </p>
            )}
            <div className="card-footer">
              <span className={`badge ${getStateBadgeClass(ticket.state)}`}>{ticket.state}</span>
            </div>
            <button className="btn btn-outline-primary btn-sm" onClick={() => handleEditTicket(ticket)}>
              Edit
            </button>
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

      <Modal open={Boolean(editTicket)} onClose={() => setEditTicket(null)}>
        <Box
          sx={{
            position: 'absolute',
            top: '50%',
            left: '50%',
            transform: 'translate(-50%, -50%)',
            width: 400,
            bgcolor: 'background.paper',
            border: '2px solid #000',
            boxShadow: 24,
            p: 4,
          }}
        >
          <Typography variant="h6" component="h2" gutterBottom>
            Change Ticket State
          </Typography>
          <FormControl fullWidth margin="normal">
            <InputLabel>State</InputLabel>
            <Select value={editedState} onChange={(e) => setEditedState(e.target.value as string)}>
              <MenuItem value="raised">Raised</MenuItem>
              <MenuItem value="acknowledged">Acknowledged</MenuItem>
              <MenuItem value="resolved">Resolved</MenuItem>
            </Select>
          </FormControl>
          <TextField
            fullWidth
            margin="normal"
            label="Resolution Comment"
            multiline
            rows={4}
            value={editedComment}
            onChange={(e) => setEditedComment(e.target.value)}
          />
          <Button variant="contained" onClick={handleUpdateTicket}>
            Update
          </Button>
        </Box>
      </Modal>
    </div>
  );
};

export default Dashboard;