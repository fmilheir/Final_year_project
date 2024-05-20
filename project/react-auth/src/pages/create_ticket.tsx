import React, { useState, ChangeEvent } from 'react';
import { TextField, Button, Box, Container, Typography, Select, MenuItem, SelectChangeEvent } from '@mui/material';
import { Navigate } from 'react-router-dom';
import '../css/nav.css';


interface CreateTicketFormProps {
  Role: boolean;
}

const CreateTicketForm: React.FC<CreateTicketFormProps> = ({ Role }) => {
  const [ticketData, setTicketData] = useState({
    name: '',
    category: '',
    priority: '',
    domain: 'RAN',
    IncidentDetail: '',
    ackState: "acknowledged",
    occurTime: new Date(),
    state: "raised",
    sourceObject: [{ id: '12345678', href: '/resourceInventoryManagement/v4/resource/12345678' }],
    rootEventId: [{ id: '3008652', '@type': 'Alarm', href: 'exemple/' }]
  });
  const [redirect, setRedirect] = useState('');

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
      console.log('Creating ticket:', ticketData);
      const response = await fetch('/tmf-api/Incident/v4/incident', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(ticketData),
      });
      if (!response.ok) {
        throw new Error('Failed to create ticket');
      }
      console.log('Ticket created successfully!');
      resetForm();
      setRedirect(Role === true ? '/admin_dashboard' : '/dashboard'); // Set redirect based on user's role
    } catch (error) {
      console.error('Error creating ticket:', error);
    }
  };

  const resetForm = () => {
    setTicketData({
      name: '',
      category: '',
      priority: '',
      domain: 'RAN',
      IncidentDetail: '',
      ackState: "acknowledged",
      state: "raised",
      occurTime: new Date(),
      sourceObject: [{ id: '12345678', href: '/resourceInventoryManagement/v4/resource/12345678' }],
      rootEventId: [{ id: '30086529', '@type': 'Alarm', href: 'exemple/' }]
    });
  };

  if (redirect) {
    return <Navigate to={redirect} />;
  }

  return (
    <Container maxWidth="md">
      <Box component="form" onSubmit={handleSubmit} sx={{ marginTop: 4 }}>
        <Typography variant="h5" gutterBottom>
          Create Ticket
        </Typography>
        <TextField
          label="Name"
          name="name"
          value={ticketData.name}
          onChange={handleChange}
          fullWidth
          margin="normal"
        />
        <TextField
          label="Category"
          name="category"
          value={ticketData.category}
          onChange={handleChange}
          fullWidth
          margin="normal"
        />
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
        <TextField
          label="IncidentDetail"
          name="IncidentDetail"
          value={ticketData.IncidentDetail}
          onChange={handleChange}
          fullWidth
          margin="normal"
        />
        <Button variant="contained" color="primary" type="submit" fullWidth sx={{ marginTop: 2 }}>
          Create Ticket
        </Button>
      </Box>
    </Container>
  );
};

export default CreateTicketForm;