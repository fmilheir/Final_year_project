import React, { useState, useEffect, ChangeEvent } from 'react';
import { SyntheticEvent } from 'react';
import '../App.css';
import { get } from 'http';

interface Props {
  userID: number;
}

const Profile: React.FC<Props> = ({ userID }) => {
  const [companyId, setCompanyId] = useState<number | null>(null);
  const [showPopup, setShowPopup] = useState(false);
  const [showEditPopup, setShowEditPopup] = useState(false);
  const [showDeletePopup, setShowDeletePopup] = useState(false);
  const [firstName, setFirstName] = useState('John');
  const [lastName, setLastName] = useState('Doe');
  const [email, setEmail] = useState('doe@gmail.com');
  const [role, setRole] = useState('Admin');
  const [passwordError, setPasswordError] = useState('');
  const [currentPassword, setCurrentPassword] = useState('');
  const [newPassword, setNewPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');


  const fetchCompanyId = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/getCompanyId', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ userID:userID.toString() }),
      });
      const data = await response.json();
      setCompanyId(data.companyID);
      fetchUserDetails(data.companyID.toString());
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
      window.location.href = '/'; // Redirect to home page
    } catch (error) {
      console.error('Error deleting user:', error);
    }
  };

  const updateUser = async (e: SyntheticEvent) => {
    e.preventDefault();

    // Check if the new password and confirm password match
    if (newPassword !== confirmPassword) {
      setPasswordError('New password and confirm password do not match');
      return;
    }

    try {
      // Check if the old password is correct
      console.log('Role:', role);
      if (role === 'User'){
            console.log('USer:', role);
            const loginResponse = await fetch('http://localhost:8080/api/login_user', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email, password: currentPassword }),
            });
            if (!loginResponse.ok) {
                setPasswordError('Current password is incorrect');
                return;
            }
        }
        if (role === 'admin'){
            console.log('Admin:', role);
            const loginResponse = await fetch('http://localhost:8080/api/login_admin', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email, password: currentPassword }),
            });
            if (!loginResponse.ok) {
                setPasswordError('Current password is incorrect');
                return;
            }
        }
        // Old password is correct, proceed with the update
        await fetch('http://localhost:8080/api/update_user', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ ID: userID.toString(), firstName, lastName, email, role, newPassword }),
        });
        console.log('User updated successfully!');
        setShowPopup(false);
        setPasswordError('');
        fetchUserDetails(userID.toString());
    }
    catch (error) {
      console.error('Error updating user:', error);
    }
  };

  const fetchUserDetails = async (companyid: string | null) => {
    try {
      const response = await fetch('http://localhost:8080/api/getUserInformation', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ userID: userID.toString() }),
      });
      const data = await response.json();
      setFirstName(data.firstName);
      setLastName(data.lastName);
      setEmail(data.email);
      setRole(data.role);
    } catch (error) {
      console.error('Error fetching user details:', error);
    }
  };



  useEffect(() => {
    fetchCompanyId();
  }, []);
return (
  <main>
    <div className="container">
      <div className="row justify-content-center">
        <div className="col-md-6">
          <h1 className="text-center mb-4">
            <i className="fas fa-user me-2"></i>Profile
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
            {showEditPopup && (
              <div className="modal d-block">
                <div className="modal-dialog modal-dialog-centered">
                  <div className="modal-content">
                    <div className="modal-header">
                      <h5 className="modal-title">Update Details</h5>
                      <button type="button" className="btn-close" onClick={() => setShowEditPopup(false)}></button>
                    </div>
                    <form onSubmit={updateUser}>
                      <div className="modal-body">
                        <div className="mb-3">
                          <label htmlFor="firstName" className="form-label">First Name:</label>
                          <input type="text" className="form-control" id="firstName" value={firstName} onChange={(e) => setFirstName(e.target.value)} />
                        </div>
                        <div className="mb-3">
                          <label htmlFor="lastName" className="form-label">Last Name:</label>
                          <input type="text" className="form-control" id="lastName" value={lastName} onChange={(e) => setLastName(e.target.value)} />
                        </div>
                        <div className="mb-3">
                          <label htmlFor="email" className="form-label">Email:</label>
                          <input type="email" className="form-control" id="email" value={email} onChange={(e) => setEmail(e.target.value)} />
                        </div>
                        <div className="mb-3">
                          <label htmlFor="currentPassword" className="form-label">Current Password:</label>
                          <input type="password" className="form-control" id="currentPassword" value={currentPassword} onChange={(e) => setCurrentPassword(e.target.value)} />
                        </div>
                        <div className="mb-3">
                          <label htmlFor="newPassword" className="form-label">New Password:</label>
                          <input type="password" className="form-control" id="newPassword" value={newPassword} onChange={(e) => setNewPassword(e.target.value)} />
                        </div>
                        <div className="mb-3">
                          <label htmlFor="confirmPassword" className="form-label">Confirm Password:</label>
                          <input type="password" className="form-control" id="confirmPassword" value={confirmPassword} onChange={(e) => setConfirmPassword(e.target.value)} />
                        </div>
                        {passwordError && <p className="text-danger">{passwordError}</p>}
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
            {showDeletePopup && (
              <div className="modal d-block">
                <div className="modal-dialog modal-dialog-centered">
                  <div className="modal-content">
                    <div className="modal-header">
                      <h5 className="modal-title">Confirm Delete</h5>
                      <button type="button" className="btn-close" onClick={() => setShowDeletePopup(false)}></button>
                    </div>
                    <div className="modal-body">
                      <p>Are you sure you want to delete your account?</p>
                    </div>
                    <div className="modal-footer">
                      <button type="button" className="btn btn-danger" onClick={deleteUser}>Delete</button>
                      <button type="button" className="btn btn-secondary" onClick={() => setShowDeletePopup(false)}>Cancel</button>
                    </div>
                  </div>
                </div>
              </div>
            )}
            <div className="card-footer d-flex justify-content-between">
              <button className="btn btn-primary" onClick={() => setShowEditPopup(true)}>
                Update Details
              </button>
              <button className="btn btn-danger" onClick={() => setShowDeletePopup(true)}>
                Delete Profile
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
);
}
export default Profile;
