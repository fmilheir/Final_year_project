import React, { useState, useEffect, ChangeEvent } from 'react';
import { SyntheticEvent } from 'react';
import '../css/nav.css';



interface Props {
    userID: number;
  }

const Profile : React.FC<Props> = ({ userID }) => {
    const [companyId, setCompanyId] = useState<number | null>(null);



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
    }

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
                                    <p><strong>First Name:</strong> John</p>
                                    <p><strong>Last Name:</strong> Doe</p>
                                    <p><strong>Email:</strong> doe@gmail.com</p>
                                    <p><strong>Role:</strong> Admin</p>
                                </div>
                            </div>
                            <div className="card-footer d-flex justify-content-between">
                                <button className="btn btn-primary">Edit Profile</button>
                                <button className="btn btn-danger">Delete Profile</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </main>
    );
}
export default Profile;