import React, { useState, ChangeEvent } from 'react';
import { TextField, Button, Box, Container, Typography, Select, MenuItem, SelectChangeEvent} from '@mui/material';
import axios from 'axios';
import '../css/nav.css';

const CreateTicketForm = () => {
  const [ticketData, setTicketData] = useState({
    name: '',
    category: '',
    priority: '',
    domain: 'RAN',
    IncidentDetail: '',
    sourceObject: [{
      id: '12345678',
      href: '/resourceInventoryManagement/v4/resource/12345678'
    }],
    rootEventId: [{
      id: '30086529',
      '@type': 'Alarm',
      href: 'exemple/'
    }]
  });

  const handleChange = (event: ChangeEvent<HTMLInputElement | { name?: string; value: unknown }>) => {
    const { name, value } = event.target;
    setTicketData({ ...ticketData, [name as string]: value });
  };

  const handleChangeDropdown = (event: SelectChangeEvent<string>) => {
    const { name, value } = event.target;
    setTicketData({ ...ticketData, [name]: value });
  };

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    try {
      await axios.post('http://localhost:8080/api/create_ticket', ticketData);
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
      domain: '',
      IncidentDetail: '',
      sourceObject: [{
        id: '',
        href: ''
      }],
      rootEventId: [{
        id: '',
        '@type': '',
        href: ''
      }]
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
        <Select
          label="Priority"
          name="priority"
          value={ticketData.priority}
          onChange={handleChangeDropdown}
          fullWidth
          margin="dense"
        >
          <MenuItem value="low">Low</MenuItem>
          <MenuItem value="medium">Medium</MenuItem>
          <MenuItem value="high">High</MenuItem>
        </Select>
        <TextField label="IncidentDetail" name="IncidentDetail" value={ticketData.IncidentDetail} onChange={handleChange} fullWidth margin="normal" />
        <Button variant="contained" color="primary" type="submit" fullWidth sx={{ marginTop: 2 }}>
          Create Ticket
        </Button>
      </Box>
    </Container>
  );
};

export default CreateTicketForm;
