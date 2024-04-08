import React, { SyntheticEvent, useState } from "react";
import { Navigate } from "react-router-dom";
import '../App.css';

const LoginAdmin: React.FC = () => {
  const [errorMessage, setErrorMessage] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [redirect, setRedirect] = useState<boolean>(false);

  const handleSubmit = async (e: SyntheticEvent) => {
    e.preventDefault();

    try {
      const response = await fetch("http://localhost:8080/api/login_admin", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email,
          password,
        }),
      });

      if (!response.ok) {
        throw new Error("Invalid credentials");
      }
      else { 
        // get JWT token from response headers
        const jwtToken = response.json();
        console.log(document.cookie); // Log the JWT token to the console
        console.log(response.headers.entries); // Log the JWT token to the console
        console.log(jwtToken); // Log the JWT token to the console
        if (jwtToken) {
          document.cookie = ("jwt="+jwtToken);
          //localStorage.setItem("jwt", jwtToken);
          document.cookie = "admin=true";
        }
      }
      setRedirect(true);
    } catch (error) {
      setErrorMessage("Invalid email or password. Please try again.");
    }
  };

  if (redirect) return <Navigate to="/dashboard" />;

  return (
    <main>
      {errorMessage && (
        <div className="alert alert-danger" role="alert">
          {errorMessage}
        </div>
      )}
      <div className="container">
        <div className="row d-flex justify-content-center">
          <div className="col-md-6">
            <h1 className="text-center mb-4">
              <i className="fas fa-lock"></i>
              Admin Portal
            </h1>
            <form onSubmit={handleSubmit}>
              <div className="form-group">
                <label htmlFor="email">Email Address</label>
                <input
                  type="email"
                  className="form-control"
                  id="email"
                  placeholder="Enter email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  required
                />
              </div>
              <div className="form-group">
                <label htmlFor="password">Password</label>
                <input
                  type="password"
                  className="form-control"
                  id="password"
                  placeholder="Enter password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  required
                />
              </div>
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

export default LoginAdmin;
