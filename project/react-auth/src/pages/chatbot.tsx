import React, { useState } from 'react';

interface Message {
  role: string;
  content: string;
}

const Chatbot: React.FC = () => {
  const [messages, setMessages] = useState<Message[]>([]);
  const [inputValue, setInputValue] = useState('');
  const [isLoading, setIsLoading] = useState(false); // Added loading state

  const sendMessage = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (inputValue.trim() !== '') {
      const newMessage: Message = {
        role: 'user',
        content: inputValue,
      };
      setMessages((prevMessages) => [...prevMessages, newMessage]);
      setIsLoading(true); // Set loading to true

      fetch('http://localhost:8080/api/chatbot', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ message: inputValue }),
      })
        .then((response) => response.json())
        .then((data) => {
          const botResponse: Message = {
            role: 'assistant',
            content: data.text,
          };
          setMessages((prevMessages) => [...prevMessages, botResponse]);
          setIsLoading(false); // Set loading to false after receiving response
        });

      setInputValue('');
    }
  };

  return (
    <div className="chatbot-container">
      <div className="chat-history">
        {messages.map((message, index) => (
          <div key={index} className={`message ${message.role}`}>
            {message.content}
          </div>
        ))}
        {isLoading && <div className="message loading">Waiting for response...</div>} {/* Display waiting message */}
      </div>
      <form onSubmit={sendMessage} className="input-container">
        <input
          type="text"
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
          placeholder="Type your message here..."
          className="input-field"
        />
        <button type="submit" className="send-button">Send</button>
      </form>
    </div>
  );
};

export default Chatbot;
