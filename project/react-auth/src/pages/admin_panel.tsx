import {
  Button,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
} from '@mui/material';
import React, { useState, useEffect, ChangeEvent, SyntheticEvent } from 'react';
import '../css/admin_panel.css';


interface User {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
  role: string;
  id?: number;
  FirstName?: string;
  LastName?: string;
  Email?: string;
  Role?: string;
}

interface Props {
  userID: number;
}

const AdminPanel: React.FC<Props> = ({ userID }) => {
  const [users, setUsers] = useState<User[]>([]);
  const [newUser, setNewUser] = useState<User>({ firstName: '', lastName: '', email: '', role: 'admin', password: '' });
  const [companyId, setCompanyId] = useState<number | null>(null);

  // Function to fetch company ID from the server
  const fetchCompanyId = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/getCompanyId', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ userID }),
      });
      const data = await response.json();
      setCompanyId(data.companyID);
      fetchUsers(data.companyID.toString());
    } catch (error) {
      console.error('Error fetching company ID:', error);
    }
  };

  // Function to fetch existing users from the server
  const fetchUsers = async (companyid: string | null) => {
    try {
      const response = await fetch('http://localhost:8080/api/get_users', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ companyid }),
      });
      const data = await response.json();
      setUsers(data);
    } catch (error) {
      console.error('Error fetching users:', error);
    }
  };


   //method for delleting user calling a post request to localhost 8080
   const deleteUser = async (e: SyntheticEvent) => {
    e.preventDefault();
    try {
        await fetch('http://localhost:8080/api/delete', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ ID: userID }),

    });
    console.log('User deleted successfully!');
    } catch (error) {

        console.error('Error deleting user:', error);
    }
};

//method for upading user details calling a post request to localhost 8080
const updateUser = async (e: SyntheticEvent) => {
    e.preventDefault();
    try {
        await fetch('http://localhost:8080/api/update', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ ID: userID }),

    });
    console.log('User updated successfully!');
    } catch (error) {

        console.error('Error updating user:', error);
    }
};


  // Function to handle form submission for creating a new user
  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    try {
      const endpoint = 'http://localhost:8080/api/register_user';
      const response = await fetch(endpoint, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ ...newUser, companyID: companyId?.toString() }),
      });
      if (response.ok) {
        const companyID = companyId?.toString() || null;
        fetchUsers(companyID);
        setNewUser({ firstName: '', lastName: '', email: '', role: '', password: '' });
        setCompanyId(null);
      } else {
        console.error('Failed to create user');
      }
    } catch (error) {
      console.error('Error creating user:', error);
    }
  };

  const handleInputChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setNewUser((prevUser) => ({
      ...prevUser,
      [name]: value,
    }));
  };

  const handleRoleChange = (event: ChangeEvent<HTMLSelectElement>) => {
    const { value } = event.target;
    setNewUser((prevUser) => ({
      ...prevUser,
      role: value,
    }));
  };


  // Function to handle editing a user
  const handleEdit = (user: User) => {
    // Logic to handle editing a user
    console.log('Editing user:', user);
  };

  // Function to handle deleting a user
  const handleDelete = (user: User) => {
    // Logic to handle deleting a user
    console.log('Deleting user:', user);
  };

  useEffect(() => {
    fetchCompanyId();
  }, []);

return (
<main className="admin-panel-container">
      <section>
        <h2 className="section-header">Existing Users</h2>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>First Name</TableCell>
              <TableCell>Last Name</TableCell>
              <TableCell>Email</TableCell>
              <TableCell>Role</TableCell>
              <TableCell>Edit</TableCell>
              <TableCell>Delete</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {users.map((user) => (
              <TableRow key={user.id}>
                <TableCell>{user.FirstName}</TableCell>
                <TableCell>{user.LastName}</TableCell>
                <TableCell>{user.Email}</TableCell>
                <TableCell>{user.Role}</TableCell>
                <TableCell>
                  <Button variant="outlined" color="primary" onClick={() => handleEdit(user)}>
                    Edit
                  </Button>
                </TableCell>
                <TableCell>
                  <Button variant="outlined" color="secondary" onClick={() => handleDelete(user)}>
                    Delete
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </section>
      <section>
        <h2 className="section-header">Create New User</h2>
        <form onSubmit={handleSubmit}>
          <div className="form-input">
            <label htmlFor="firstName" className="form-label">First Name:</label>
            <input
              type="text"
              id="firstName"
              name="firstName"
              value={newUser.firstName}
              onChange={handleInputChange}
            />
          </div>
          <div className="form-input">
            <label htmlFor="lastName" className="form-label">Last Name:</label>
            <input
              type="text"
              id="lastName"
              name="lastName"
              value={newUser.lastName}
              onChange={handleInputChange}
            />
          </div>
          <div className="form-input">
            <label htmlFor="email" className="form-label">Email:</label>
            <input
              type="email"
              id="email"
              name="email"
              value={newUser.email}
              onChange={handleInputChange}
            />
          </div>
          <div className="form-input">
            <label htmlFor="password" className="form-label">Password:</label>
            <input
              type="password"
              id="password"
              name="password"
              value={newUser.password}
              onChange={handleInputChange}
            />
          </div>
          <div className="form-input">
            <label htmlFor="role" className="form-label">Role:</label>
            <select
              id="role"
              name="role"
              value={newUser.role}
              onChange={handleRoleChange}
            >
              <option value="admin">Admin</option>
              <option value="normal">Normal</option>
            </select>
          </div>
          <button type="submit" className="btn btn-primary">Create User</button>
        </form>
      </section>
    </main>
  );
};

export default AdminPanel;
