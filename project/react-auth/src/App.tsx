import logo from './logo.svg';
import Nav from './components/nav';
import LoginAdmin from './pages/login_admin';
import Signup from './pages/signup';
import Home from './pages/home'; 
import LoginUser from './pages/login_user'; 
import Dashboard from './pages/dashboard';
import Chatbot from './pages/chatbot';
import { BrowserRouter, Route , Routes, Navigate } from 'react-router-dom';
import './App.css';

function App() {
  const IsAuthenticated = false;
  return (
    <div className="App">
      <BrowserRouter>
      <Nav IsAuthenticated={IsAuthenticated} />
      <main className='form-signin'>
          <Routes>
            <Route path="/" element={<Navigate replace to="/home" />} />
            <Route path="/home" element={<Home />} />
            <Route path="/login_admin" element={<LoginAdmin/>}/>
            <Route path="/login_user" element={<LoginUser/>}/>
            <Route path="/signup" element={<Signup/>}/>   
            <Route path="/dashboard" element={<Dashboard/>} />
            <Route path="/chatbot" element={<Chatbot/>} />

          </Routes>
        </main>
      </BrowserRouter>

    </div>

  );
}

export default App;
