import React, { SyntheticEvent, useState } from "react";
import '../css/nav.css';

const Profile = () => {
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