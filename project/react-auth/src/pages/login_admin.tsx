import React, { SyntheticEvent, useState } from "react";
import { Navigate } from "react-router-dom";

const LoginAdmin: React.FC = () => {
  const [messages] = useState<string[]>([]);
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [redirect, setRedirect] = useState<boolean>(false);

  const handleSubmit = async (e: SyntheticEvent) => {
    e.preventDefault();
    await fetch("http://localhost:8080/api/login_admin", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email,
        password,
      }),
    });
    setRedirect(true);
  };

  if (redirect) return <Navigate to="/home" />;

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
              Admin Portal
            </h1>
            <form onSubmit={handleSubmit}>
              <div>
                <label htmlFor="email">Email address</label>
                <div>
                  <input
                    type="email"
                    className="form-control"
                    id="email"
                    placeholder="Enter email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                  />
                </div>
              </div>
              <div>
                <label htmlFor="password">Password</label>
                <div>
                  <input
                    type="password"
                    className="form-control"
                    id="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                  />
                </div>
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
