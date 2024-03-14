import React, {SyntheticEvent, useState} from "react";
import {Navigate} from "react-router-dom";


const LoginUser = () => {
  const [messages] = useState([]);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [redirect, setRedirect] = useState(false);

  const Submit = async (e: SyntheticEvent) => {
    e.preventDefault();
    await fetch("http://localhost:8080/api/login_user", {
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
              Login
            </h1>
            <form method="POST" action="#">
              <div>
                <label htmlFor="email">Email address</label>
                <div>
                  <input type="email" className="form-control" id="email" placeholder="Enter email" name="email" onChange={(e) => setEmail(e.target.value)} />
                </div>
              </div>
              <div>
                <label htmlFor="password">Password</label>
                <div>
                  <div className="input-group-prepend"></div>
                  <input
                    type="password"
                    className="form-control"
                    id="password"
                    placeholder="Password"
                    name="password"
                    onChange={(e) => setPassword(e.target.value)}
                  />
                </div>
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