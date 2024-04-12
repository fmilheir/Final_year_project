import React, { SyntheticEvent, useState } from "react";

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
                                <p className="card-text">First Name: John</p>
                                <p className="card-text">Last Name: Doe</p>
                                <p className="card-text">Email:doe@gmail .com</p>
                                <p className="card-text">Role: Admin</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </main>
    );
}

export default Profile;