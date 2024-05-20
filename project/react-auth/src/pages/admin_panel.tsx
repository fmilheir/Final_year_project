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
  ID?: number;
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
  const [showCreatePopup, setShowCreatePopup] = useState(false);
  const [showEditPopup, setShowEditPopup] = useState(false);
  const [showDeletePopup, setShowDeletePopup] = useState(false);
  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  const [editedUser, setEditedUser] = useState<User>({ firstName: '', lastName: '', email: '', role: '', password: '' });

  // Function to fetch company ID from the server
  const fetchCompanyId = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/getCompanyId', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ userID: userID.toString() }),
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
        body: JSON.stringify({ companyid: companyid?.toString() }),
      });
      const data = await response.json();
      if (response.status === 404) {
        console.error('No users found');
      }
      const usersData = data.map((user: User) => ({
        id: user.ID,
        FirstName: user.FirstName,
        LastName: user.LastName,
        Email: user.Email,
        Role: user.Role,
      }));
      console.log('Users:', usersData);
      setUsers(usersData);
    } catch (error) {
      console.error('Error fetching users:', error);
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
        setShowCreatePopup(false);
      } else {
        console.error('Failed to create user');
      }
    } catch (error) {
      console.error('Error creating user:', error);
    }
  };

  const handleInputChange = (event: ChangeEvent<HTMLInputElement>) => {
      const { name, value } = event.target;
      setEditedUser((prevUser) => ({
        ...prevUser,
        [name]: value,
      }));
    };
    
    const handleEditRoleChange = (event: ChangeEvent<HTMLSelectElement>) => {
      const { value } = event.target;
      setEditedUser((prevUser) => ({
        ...prevUser,
        role: value,
      }));
    };

  // Function to handle editing a user
  const handleEdit = (user: User) => {
    setSelectedUser(user);
    setEditedUser({
      firstName: user.FirstName || '',
      lastName: user.LastName || '',
      email: user.Email || '',
      role: user.Role || '',
      password: '', // Leave password empty for security reasons
    });
    setShowEditPopup(true);
  };
  

  // Function to handle deleting a user
  const handleDelete = (user: User) => {
    setSelectedUser(user);
    setShowDeletePopup(true);
  };

  const handleEditInputChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setEditedUser((prevUser) => ({
      ...prevUser,
      [name]: value,
    }));
  };

  const updateUser = async (e: SyntheticEvent) => {
    e.preventDefault();
    if (selectedUser) {
      try {
        await fetch('http://localhost:8080/api/update', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ ID: selectedUser.ID, ...editedUser }),
        });
        console.log('User updated successfully!');
        setShowEditPopup(false);
        fetchUsers(companyId?.toString() || null);
      } catch (error) {
        console.error('Error updating user:', error);
      }
    }
  };

  const deleteUser = async (e: SyntheticEvent) => {
    e.preventDefault();
    if (selectedUser) {
      try {
        await fetch('http://localhost:8080/api/delete', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ ID: selectedUser.ID }),
        });
        console.log('User deleted successfully!');
        setShowDeletePopup(false);
        fetchUsers(companyId?.toString() || null);
      } catch (error) {
        console.error('Error deleting user:', error);
      }
    }
  };

  useEffect(() => {
    fetchCompanyId();
  }, []);

  return (
    <main className="admin-panel-container">
      <div className="admin-panel-container container my-5">
        <h1 className="page-header">
          <i className="fas fa-user"></i>
          Admin Panel
        </h1>
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
              {users.map((User) => (
                <TableRow key={User.ID}>
                  <TableCell>{User.FirstName}</TableCell>
                  <TableCell>{User.LastName}</TableCell>
                  <TableCell>{User.Email}</TableCell>
                  <TableCell>{User.Role}</TableCell>
                  <TableCell>
                    <Button variant="outlined" color="primary" onClick={() => handleEdit(User)}>
                      Edit
                    </Button>
                  </TableCell>
                  <TableCell>
                    <Button variant="outlined" color="secondary" onClick={() => handleDelete(User)}>
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
          <Button variant="contained" color="primary" onClick={() => setShowCreatePopup(true)}>
            Create New User
          </Button>
          {showCreatePopup && (
            <div className="modal d-block">
              <div className="modal-dialog modal-dialog-centered">
                <div className="modal-content">
                  <div className="modal-header">
                    <h5 className="modal-title">Create New User</h5>
                    <button type="button" className="btn-close" onClick={() => setShowCreatePopup(false)}></button>
                  </div>
                  <form onSubmit={handleSubmit}>
                    <div className="modal-body">
                      <div className="mb-3">
                        <label htmlFor="firstName" className="form-label">First Name:</label>
                        <input
                          type="text"
                          id="firstName"
                          name="firstName"
                          className="form-control"
                          value={newUser.firstName}
                          onChange={handleInputChange}
                        />
                      </div>
                      <div className="mb-3">
                        <label htmlFor="lastName" className="form-label">Last Name:</label>
                        <input
                          type="text"
                          id="lastName"
                          name="lastName"
                          className="form-control"
                          value={newUser.lastName}
                          onChange={handleInputChange}
                        />
                      </div>
                      <div className="mb-3">
                        <label htmlFor="email" className="form-label">Email:</label>
                        <input
                          type="email"
                          id="email"
                          name="email"
                          className="form-control"
                          value={newUser.email}
                          onChange={handleInputChange}
                        />
                      </div>
                      <div className="mb-3">
                        <label htmlFor="role" className="form-label">Role:</label>
                        <select
                          id="role"
                          name="role"
                          className="form-select"
                          value={newUser.role}
                          onChange={handleEditRoleChange}
                        >
                          <option value="admin">Admin</option>
                          <option value="normal">Normal</option>
                        </select>
                      </div>
                    </div>
                    <div className="modal-footer">
                      <button type="submit" className="btn btn-primary">Create User</button>
                      <button type="button" className="btn btn-secondary" onClick={() => setShowCreatePopup(false)}>Cancel</button>
                    </div>
                  </form>
                </div>
              </div>
            </div>
          )}
        </section>
      </div>

      {showEditPopup && selectedUser && (
        <div className="modal d-block">
          <div className="modal-dialog modal-dialog-centered">
            <div className="modal-content">
              <div className="modal-header">
                <h5 className="modal-title">Edit User</h5>
                <button type="button" className="btn-close" onClick={() => setShowEditPopup(false)}></button>
              </div>
              <form onSubmit={updateUser}>
                <div className="modal-body">
                  <div className="mb-3">
                    <label htmlFor="editFirstName" className="form-label">First Name:</label>
                    <input
                      type="text"
                      id="editFirstName"
                      name="firstName"
                      className="form-control"
                      value={editedUser.firstName}
                      onChange={handleEditInputChange}
                    />
                  </div>
                  <div className="mb-3">
                    <label htmlFor="editLastName" className="form-label">Last Name:</label>
                    <input
                      type="text"
                      id="editLastName"
                      name="lastName"
                      className="form-control"
                      value={editedUser.lastName}
                      onChange={handleEditInputChange}
                    />
                  </div>
                  <div className="mb-3">
                    <label htmlFor="editEmail" className="form-label">Email:</label>
                    <input
                      type="email"
                      id="editEmail"
                      name="email"
                      className="form-control"
                      value={editedUser.email}
                      onChange={handleEditInputChange}
                    />
                  </div>
                  <div className="mb-3">
                  <label htmlFor="editRole" className="form-label">Role:</label>
                  <select
                    id="editRole"
                    name="role"
                    className="form-select"
                    value={editedUser.role}
                    onChange={handleEditRoleChange}
                  >
                    <option value="admin">Admin</option>
                    <option value="normal">Normal</option>
                  </select>
                </div>
              </div>
                <div className="modal-footer">
                  <button type="submit" className="btn btn-primary">Save</button>
                  <button type="button" className="btn btn-secondary" onClick={() => setShowEditPopup(false)}>Cancel</button>
                </div>
              </form>
            </div>
          </div>
        </div>
      )}

      {showDeletePopup && selectedUser && (
        <div className="modal d-block">
          <div className="modal-dialog modal-dialog-centered">
            <div className="modal-content">
              <div className="modal-header">
                <h5 className="modal-title">Confirm Delete</h5>
                <button type="button" className="btn-close" onClick={() => setShowDeletePopup(false)}></button>
              </div>
              <div className="modal-body">
                <p>Are you sure you want to delete this user?</p>
              </div>
              <div className="modal-footer">
                <button type="button" className="btn btn-danger" onClick={deleteUser}>Delete</button>
                <button type="button" className="btn btn-secondary" onClick={() => setShowDeletePopup(false)}>Cancel</button>
              </div>
            </div>
          </div>
        </div>
      )}
    </main>
  );
};

export default AdminPanel;