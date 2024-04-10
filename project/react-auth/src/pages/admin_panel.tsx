import React, { useState, useEffect } from 'react';

interface User {
  id: number;
  username: string;
  email: string;
  role: string;
}

const AdminPanel: React.FC = () => {
  const [users, setUsers] = useState<User[]>([]);
  const [newUser, setNewUser] = useState<User>({ id: 0, username: '', email: '', role: '' });

  // Function to fetch existing users from the server
  const fetchUsers = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/users');
      const data = await response.json();
      setUsers(data);
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
        body: JSON.stringify(newUser),
      });
      if (response.ok) {
        // Refresh user list after successfully adding new user
        fetchUsers();
        // Clear input fields
        setNewUser({ id: 0, username: '', email: '', role: '' });
      } else {
        console.error('Failed to create user');
      }
    } catch (error) {
      console.error('Error creating user:', error);
    }
  };

  // Function to handle input changes for new user form
  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = event.target;
    setNewUser((prevUser) => ({
      ...prevUser,
      [name]: value,
    }));
  };

  useEffect(() => {
    // Fetch users when component mounts
    fetchUsers();
  }, []);

  return (
    <div>
      <h2>Admin Panel</h2>
      <div>
        <h3>Existing Users</h3>
        <ul>
          {users.map((user) => (
            <li key={user.id}>
              Username: {user.username}, Email: {user.email}, Role: {user.role}
            </li>
          ))}
        </ul>
      </div>
      <div>
        <h3>Create New User</h3>
        <form onSubmit={handleSubmit}>
          <input type="text" name="username" value={newUser.username} placeholder="Username" onChange={handleInputChange} />
          <input type="email" name="email" value={newUser.email} placeholder="Email" onChange={handleInputChange} />
          <select name="role" value={newUser.role} onChange={handleInputChange}>
            <option value="admin">Admin</option>
            <option value="normal">Normal</option>
          </select>
          <button type="submit">Create User</button>
        </form>
      </div>
    </div>
  );
};

export default AdminPanel;
