import React, { useEffect, useState } from 'react';
import { Navigate } from 'react-router-dom';
import { Box, Button, Card, CardContent, Typography, CircularProgress } from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import axios from 'axios'; // Import axios for API requests

interface Ticket {
    id: number;
    title: string;
    priority: string;
    description: string;
}

const Dashboard = () => {
    const [tickets, setTickets] = useState<Ticket[]>([]); 
    const [isLoading, setIsLoading] = useState(true); // Change initial state to true
    const [redirect, setRedirect] = useState(false);

 

    useEffect(() => {
        const fetchTickets = async () => {
            try {
                const response = await axios.get('http://localhost:8080/api/tickets'); // Use axios for API request
                setTickets(response.data);
            } catch (error) {
                console.error('Error fetching tickets:', error);
            } finally {
                setIsLoading(false); // Update loading state after fetching
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
                            <Card key={ticket.id} sx={{ minWidth: 275 }}>
                                <CardContent>
                                    <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
                                        Ticket ID: {ticket.id}
                                    </Typography>
                                    <Typography variant="h5" component="div">
                                        {ticket.title}
                                    </Typography>
                                    <Typography sx={{ mb: 1.5 }} color="text.secondary">
                                        Priority: {ticket.priority}
                                    </Typography>
                                    <Typography variant="body2">
                                        {ticket.description}
                                    </Typography>
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
