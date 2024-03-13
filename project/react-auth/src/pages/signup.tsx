import React, {SyntheticEvent, useState} from "react";
import { Navigate } from 'react-router-dom';

const  Signup = () => {

    const [messages] = useState([]);
    const [email, setEmail] = useState("");
    const [firstName, setFirstName] = useState("");
    const [password1, setPassword1] = useState("");
    const [password2, setPassword2] = useState("");
    const [company, setCompany] = useState("");
    const [redirect, setRedirect] = useState(false);

    const Submit = async (e: SyntheticEvent) => {
      console.log("submit");
        e.preventDefault();
        await fetch("http://localhost:8000/api/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                email,
                firstName,
                password1,
                password2,
                company,
            }),
            });
            setRedirect(true);

    }
    if (redirect)
        return <Navigate to="/login_admin" />;

    return (
      <main>
        {messages.map((message, index) => (
          <div key={index} className={`alert alert-${message}`}>
            {message}
          </div>
        ))}
        <div className="container my-5">
          <div className="row d-flex justify-content-center">
            <div className="col-md-6">
              <h1 className="text-center mb-4">
                <i className="fas fa-lock"></i>
                Sign Up
              </h1>
              <form onSubmit={Submit}>
                <div className="form-group">
                  <label htmlFor="email">Email Address</label>
                  <input
                    type="email"
                    className="form-control"
                    id="email"
                    name="email"
                    placeholder="Enter email"
                    required
                    onChange={(e) => setEmail(e.target.value)}
                  />
                </div>
                <div className="form-group">
                  <label htmlFor="firstName">First Name</label>
                  <input
                    type="text"
                    className="form-control"
                    id="firstName"
                    name="firstName"
                    placeholder="Enter first name"
                    required
                    onChange={(e) => setFirstName(e.target.value)}
                  />
                </div>
                <div className="form-group">
                  <label htmlFor="password1">Password</label>
                  <input
                    type="password"
                    className="form-control"
                    id="password1"
                    name="password1"
                    placeholder="Enter password"
                    required
                    onChange={(e) => setPassword1(e.target.value)}
                  />
                </div>
                <div className="form-group">
                  <label htmlFor="password2">Password (Confirm)</label>
                  <input
                    type="password"
                    className="form-control"
                    id="password2"
                    name="password2"
                    placeholder="Confirm password"
                    required
                    onChange={(e) => setPassword2(e.target.value)}
                  />
                </div>
                <div className="form-group">
                  <label htmlFor="company">Please enter the name of your company</label>
                  <input
                    type="text"
                    className="form-control"
                    id="company"
                    name="company"
                    placeholder="Enter company name"
                    required
                    onChange={(e) => setCompany(e.target.value)}
                  />
                </div>
                <br />
                <button type="submit" className="btn btn-primary btn-block">
                  Submit
                </button>
              </form>
            </div>
          </div>
        </div>
      </main>
    );
  };
  
  export default Signup;