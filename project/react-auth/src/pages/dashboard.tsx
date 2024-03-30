import React, { useEffect, useState } from 'react';
import { Navigate } from 'react-router-dom';
import { Box, Button, Card, CardContent, Typography, CircularProgress } from '@mui/material';
import AddIcon from '@mui/icons-material/Add';

const Dashboard = () => {
    const [tickets, setTickets] = useState([]);
    const [isLoading, setIsLoading] = useState(false);
    const [redirect, setRedirect] = useState(false);

    useEffect(() => {
        const fetchTickets = async () => {
            setIsLoading(true);
            // Placeholder for fetch call, replace with your actual API call
            const response = await fetch('http://localhost:8080/api/tickets');
            const data = await response.json();
            setTickets(data);
            setIsLoading(false);
        };

        fetchTickets();
    }, []);

    const handleCreateTicket = () => {
        // Placeholder function for handling ticket creation
        // This could set state that causes redirection to a ticket creation form
        setRedirect(true);
    };

    if (redirect) {
        // Navigate to the ticket creation page
        return <Navigate to="/create_ticket" />;
    }

    //if (isLoading) {
      //  return <CircularProgress />;
    //}

    return (
        <Box sx={{ flexGrow: 1, m: 3 }}>
            <Typography variant="h4" gutterBottom>
                Dashboard
            </Typography>
            <Button variant="contained" startIcon={<AddIcon />} onClick={handleCreateTicket}>
                Create New Ticket
            </Button>
            <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2, mt: 3 }}>
                {tickets.length ? (
                    tickets.map((ticket, index) => (
                        <Card key={index} sx={{ minWidth: 275 }}>
                            <CardContent>
                                <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
                                    Ticket ID: {/*{ticket.id} */}
                                </Typography>
                                <Typography variant="h5" component="div">
                                    {/*{ticket.title}*/}
                                </Typography>
                                <Typography sx={{ mb: 1.5 }} color="text.secondary">
                                    Priority: {/*{ticket.priority}*/}
                                </Typography>
                                <Typography variant="body2">
                                    {/*{ticket.description}*/}
                                </Typography>
                            </CardContent>
                        </Card>
                    ))
                ) : (
                    <Typography>No tickets available.</Typography>
                )}
            </Box>
        </Box>
    );
};

export default Dashboard;
