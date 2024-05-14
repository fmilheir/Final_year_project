import React, { SyntheticEvent, useState } from "react";
import { Navigate } from "react-router-dom";
import '../App.css';

const LoginUser = () => {
  const [errorMessage, setErrorMessage] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [redirect, setRedirect] = useState<boolean>(false);

  const handleSubmit = async (e: SyntheticEvent) => {
    e.preventDefault();

    try {
      const response = await fetch("http://localhost:8080/api/login_user", {
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
        const data = await response.json();
        console.log(data);
        console.log(data.jwt); // Log the JWT token to the console
        const jwtToken = data.jwt;
        document.cookie = jwtToken;
       // localStorage.setItem("jwt", jwtToken);
        document.cookie = "user=true";
        window.location.reload();
      }

      setRedirect(true);
    } catch (error) {
      setErrorMessage("Invalid email or password. Please try again.");
    }
  };

  if (redirect) return <Navigate to="/dashboard" />;

  return (
    <main>
    <div className="container bg-white" >
        {errorMessage && (
          <div className="alert alert-danger" role="alert">
            {errorMessage}
          </div>
        )}
        <div className="row d-flex justify-content-center">
          <div className="col-md-6">
            <h1 className="text-center mb-4">
              <i className="fas fa-lock"></i>
              Login
            </h1>
            <form onSubmit={handleSubmit}>
              <div className="form-group">
                <label htmlFor="email">Email address</label>
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
              <p>
                <a href="/signup">Don't have an account? Sign up now!</a>
              </p>
              <div className="form-check d-flex align-items-center">
                <input type="checkbox" className="form-check-input" id="rememberMe" />
                <label className="form-check-label" htmlFor="rememberMe">
                  Remember me
                </label>
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

export default LoginUser;
