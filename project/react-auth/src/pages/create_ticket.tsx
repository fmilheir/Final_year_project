import React, { useState, ChangeEvent } from 'react';
import { TextField, Button, Box, Container } from '@mui/material';
import axios from 'axios';

const CreateTicketForm = () => {
  const [ticketData, setTicketData] = useState({
    name: '',
    category: '',
    priority: '',
    severity: '',
    state: '',
    ackState: '',
    occurTime: '',
    domain: '',
    sourceObject: '',
    rootEventId: '',
  });

  
  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setTicketData({ ...ticketData, [name]: value });
  };
  
  // ... rest of your component
  

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    try {
      await axios.post('http://localhost:8080/tmf-api/Incident/v4/incident', ticketData);
      console.log('Ticket created successfully!');
      setTicketData({
        name: '',
        category: '',
        priority: '',
        severity: '',
        state: '',
        ackState: '',
        occurTime: '',
        domain: '',
        sourceObject: '',
        rootEventId: '',
      });
    } catch (error) {
      console.error('Error creating ticket:', error);
    }
  };

  return (
    <Container maxWidth="sm">
      <Box
        component="form"
        onSubmit={handleSubmit}
        noValidate
        autoComplete="off"
        sx={{
          '& .MuiTextField-root': { m: 1, width: '100%' },
          marginTop: 3,
        }}
      >
        <TextField label="Name" name="name" value={ticketData.name} onChange={handleChange} />
        <TextField label="Category" name="category" value={ticketData.category} onChange={handleChange} />
        <TextField label="Priority" name="priority" value={ticketData.priority} onChange={handleChange} />
        <TextField label="Severity" name="severity" value={ticketData.severity} onChange={handleChange} />
        <TextField label="State" name="state" value={ticketData.state} onChange={handleChange} />
        <TextField label="Acknowledgement State" name="ackState" value={ticketData.ackState} onChange={handleChange} />
        <TextField label="Occur Time" name="occurTime" value={ticketData.occurTime} onChange={handleChange} />
        <TextField label="Domain" name="domain" value={ticketData.domain} onChange={handleChange} />
        <TextField label="Source Object" name="sourceObject" value={ticketData.sourceObject} onChange={handleChange} />
        <TextField label="Root Event ID" name="rootEventId" value={ticketData.rootEventId} onChange={handleChange} />
        <Button variant="contained" color="primary" type="submit">
          Create Ticket
        </Button>
      </Box>
    </Container>
  );
};

export default CreateTicketForm;
