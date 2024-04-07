import React, { useState, ChangeEvent } from 'react';
import { TextField, Button, Box, Container, Typography } from '@mui/material';
import axios from 'axios';

const CreateTicketForm = () => {
  const [ticketData, setTicketData] = useState({
    name: '',
    category: '',
    priority: '',
    severity: '',
    domain: '',
    sourceObject: '',
    rootEventId: '',
  });

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setTicketData({ ...ticketData, [name]: value });
  };

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    try {
      await axios.post('http://localhost:8080/tmf-api/Incident/v4/incident', ticketData);
      console.log('Ticket created successfully!');
      resetForm(); // Reset form after successful submission
    } catch (error) {
      console.error('Error creating ticket:', error);
    }
  };

  const resetForm = () => {
    setTicketData({
      name: '',
      category: '',
      priority: '',
      severity: '',
      domain: '',
      sourceObject: '',
      rootEventId: '',
    });
  };

  return (
    <Container maxWidth="md">
      <Box
        component="form"
        onSubmit={handleSubmit}
        sx={{ marginTop: 4 }}
      >
        <Typography variant="h5" gutterBottom>
          Create Ticket
        </Typography>
        <TextField label="Name" name="name" value={ticketData.name} onChange={handleChange} fullWidth margin="normal" />
        <TextField label="Category" name="category" value={ticketData.category} onChange={handleChange} fullWidth margin="normal" />
        <TextField label="Priority" name="priority" value={ticketData.priority} onChange={handleChange} fullWidth margin="normal" />
        <TextField label="Severity" name="severity" value={ticketData.severity} onChange={handleChange} fullWidth margin="normal" />
        <TextField label="Domain" name="domain" value={ticketData.domain} onChange={handleChange} fullWidth margin="normal" />
        <TextField label="Source Object" name="sourceObject" value={ticketData.sourceObject} onChange={handleChange} fullWidth margin="normal" />
        <TextField label="Root Event ID" name="rootEventId" value={ticketData.rootEventId} onChange={handleChange} fullWidth margin="normal" />
        <Button variant="contained" color="primary" type="submit" fullWidth sx={{ marginTop: 2 }}>
          Create Ticket
        </Button>
      </Box>
    </Container>
  );
};

export default CreateTicketForm;
