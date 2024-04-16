import React, { useEffect, useState } from 'react';
import { Link, Navigate } from 'react-router-dom'; // Import Link from react-router-dom
import { Box, Button, Card, CardContent, Typography, CircularProgress } from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import axios from 'axios';
import '../css/nav.css';

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
                const response = await fetch('https://localhost:3030/tmf-api/Incident/v4/incident',
                    {
                        method: 'GET',
                        headers: {
                            'Content-Type': 'application/json',
                        },
                        //agent: new (require('https').Agent)({ rejectUnauthorized: false })

                    } // Add agent to ignore self-signed certificate
                );
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
        <Box sx={{ flexGrow: 1, m: 3 }}>
            <Typography variant="h4" gutterBottom>
                Dashboard
            </Typography>
            <Button variant="contained" startIcon={<AddIcon />} onClick={handleCreateTicket}>
                Create New Ticket
            </Button>
            {isLoading ? (
                <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '50vh' }}>
                    <CircularProgress />
                </Box>
            ) : (
                <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2, mt: 3 }}>
                    {tickets.length ? (
                        tickets.map((ticket) => (
                            <Card key={ticket.ID} sx={{ minWidth: 275 }}>
                                <CardContent>
                                    <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
                                        Ticket ID: {ticket.ID}
                                    </Typography>
                                    <Typography variant="h5" component="div">
                                        {ticket.name}
                                    </Typography>
                                    <Typography sx={{ mb: 1.5 }} color="text.secondary">
                                        Priority: {ticket.priority}
                                    </Typography>
                                    <Typography variant="body2">
                                        {ticket.IncidentDetail}
                                    </Typography>
                                    <Link to={`/edit_ticket/${ticket.ID}`} style={{ textDecoration: 'none' }}>
                                        <Button variant="outlined" color="primary" size="small" sx={{ mt: 2 }}>
                                            Edit
                                        </Button>
                                    </Link>
                                </CardContent>
                            </Card>
                        ))
                    ) : (
                        <Typography>No tickets available.</Typography>
                    )}
                </Box>
            )}
        </Box>
    );
};

export default Dashboard;
