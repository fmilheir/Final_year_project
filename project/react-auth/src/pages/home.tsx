import React, {SyntheticEvent, useState} from "react";
import { Navigate } from 'react-router-dom';



const Main = () => {
  const [messages] = useState([]);

  return (
    <main>
      {messages.map((message, index) => (
        <div key={index} className={`alert alert-${message}`}>
          {message}
        </div>
      ))}
      <div className="container">
        <h1 className="display-5 fw-bold lh-1 mb-3">Welcome to My CRM</h1>
        <p className="lead">This is a prototype for a CRM system. It is built using Go and MySQL. The system
          demonstrates how to build a simple CRM system using Go and MySQL. The system allows users to manage
          their customers, products, and orders.</p>


        <div className="row flex-lg-row align-items-center g-5 py-4 mb-4">
          <div className="col-12 col-lg-6">
            <h1 className="display-7 fs-20pt fw-bold mb-3"> Create an account now to register your company!</h1>
          </div>
          <div className="col-12 col-lg-6">
            <div className="d-grid gap-2 d-md-flex justify-content-md-start">
              <a href="/signup" className="btn btn-primary btn-dark btn-lg px-4 me-md-2">Create an account today!</a>
              <a href="/login_admin" className="btn btn-outline-secondary btn-lg px-4 me-md-2">Already have an
                account? Login
                now!</a>
            </div>
          </div>
        </div>
      </div>
    </main>
  );
};

export default Main;