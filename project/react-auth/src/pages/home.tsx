import React from "react";
import { Link } from "react-router-dom";

const Main = () => {
  return (
    <main className="container py-5">
      <div className="row align-items-center">
        <div className="col-lg-6">
          <h1 className="display-4 fw-bold mb-4">Welcome to the Intelligent Ticketing System</h1>
          <p className="lead mb-4">
            This prototype demonstrates a comprehensive solution for streamlining ticketing processes in large organizations. By leveraging advanced natural language processing techniques, retrieval-augmented generation, and industry standards, our intelligent chatbot system provides accurate, contextually relevant, and human-like responses to user queries.
          </p>
          <div className="d-grid gap-3 d-md-flex">
            <Link to="/signup" className="btn btn-primary btn-lg px-4">
              Create an Account
            </Link>
            <Link to="/login_admin" className="btn btn-outline-secondary btn-lg px-4">
              Login as Admin
            </Link>
          </div>
        </div>
        <div className="col-lg-6">
          <img src="path/to/chatbot-illustration.svg" alt="Chatbot Illustration" className="img-fluid" />
        </div>
      </div>
      <div className="row mt-5">
        <div className="col-lg-4">
          <div className="card border-0 shadow-sm">
            <div className="card-body">
              <h5 className="card-title">Enhanced Customer Experience</h5>
              <p className="card-text">Our chatbot provides immediate responses to common inquiries, improving customer satisfaction and reducing wait times.</p>
            </div>
          </div>
        </div>
        <div className="col-lg-4">
          <div className="card border-0 shadow-sm">
            <div className="card-body">
              <h5 className="card-title">Reduced Operational Costs</h5>
              <p className="card-text">By automating a significant portion of the ticketing process, organizations can reduce human resource costs and increase efficiency.</p>
            </div>
          </div>
        </div>
        <div className="col-lg-4">
          <div className="card border-0 shadow-sm">
            <div className="card-body">
              <h5 className="card-title">Scalability and Consistency</h5>
              <p className="card-text">Our chatbot system can handle high volumes of inquiries while maintaining consistent and accurate responses across the organization.</p>
            </div>
          </div>
        </div>
      </div>
    </main>
  );
};

export default Main;