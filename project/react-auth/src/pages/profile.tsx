import React, { useState, useEffect, ChangeEvent } from 'react';
import { SyntheticEvent } from 'react';
import '../css/nav.css';

interface Props {
  userID: number;
}

const Profile: React.FC<Props> = ({ userID }) => {
  const [companyId, setCompanyId] = useState<number | null>(null);
  const [showPopup, setShowPopup] = useState(false);
  const [firstName, setFirstName] = useState('John');
  const [lastName, setLastName] = useState('Doe');
  const [email, setEmail] = useState('doe@gmail.com');
  const [role, setRole] = useState('Admin');

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
    } catch (error) {
      console.error('Error fetching company ID:', error);
    }
  };

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

  const updateUser = async (e: SyntheticEvent) => {
    e.preventDefault();
    try {
      await fetch('http://localhost:8080/api/update', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ ID: userID, firstName, lastName, email, role }),
      });
      console.log('User updated successfully!');
      setShowPopup(false);
    } catch (error) {
      console.error('Error updating user:', error);
    }
  };

  useEffect(() => {
    fetchCompanyId();
  }, []);

  return (
    <main>
      <div className="container">
        <div className="row d-flex justify-content-center">
          <div className="col-md-6">
            <h1 className="text-center mb-4">
              <i className="fas fa-user"></i>
              Profile
            </h1>
            <div className="card">
              <div className="card-body">
                <h5 className="card-title">User Information</h5>
                <div className="user-info">
                  <p><strong>First Name:</strong> {firstName}</p>
                  <p><strong>Last Name:</strong> {lastName}</p>
                  <p><strong>Email:</strong> {email}</p>
                  <p><strong>Role:</strong> {role}</p>
                </div>
              </div>
              {showPopup && (
                <div className="popup">
                  <div className="popup-content">
                    <h3>Update Details</h3>
                    <form onSubmit={updateUser}>
                      <div className="form-group">
                        <label htmlFor="firstName">First Name:</label>
                        <input
                          type="text"
                          id="firstName"
                          value={firstName}
                          onChange={(e) => setFirstName(e.target.value)}
                        />
                      </div>
                      <div className="form-group">
                        <label htmlFor="lastName">Last Name:</label>
                        <input
                          type="text"
                          id="lastName"
                          value={lastName}
                          onChange={(e) => setLastName(e.target.value)}
                        />
                      </div>
                      <div className="form-group">
                        <label htmlFor="email">Email:</label>
                        <input
                          type="email"
                          id="email"
                          value={email}
                          onChange={(e) => setEmail(e.target.value)}
                        />
                      </div>
                      <div className="form-group">
                        <label htmlFor="role">Role:</label>
                        <input
                          type="text"
                          id="role"
                          value={role}
                          onChange={(e) => setRole(e.target.value)}
                        />
                      </div>
                      <button type="submit">Save</button>
                      <button type="button" onClick={() => setShowPopup(false)}>
                        Cancel
                      </button>
                    </form>
                  </div>
                </div>
              )}
              <div className="card-footer d-flex justify-content-between">
                <button className="btn btn-primary" onClick={() => setShowPopup(true)}>
                  Update Details
                </button>
                <button className="btn btn-danger" onClick={deleteUser}>
                  Delete Profile
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  );
};

export default Profile;