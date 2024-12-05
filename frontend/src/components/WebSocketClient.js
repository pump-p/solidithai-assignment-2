import React, { useState, useEffect, useRef } from 'react';

const WebSocketClient = () => {
  const [messages, setMessages] = useState([]);
  const [sender, setSender] = useState('');
  const [messageContent, setMessageContent] = useState('');
  const ws = useRef(null); // Use useRef to persist the WebSocket instance

  useEffect(() => {
    // Establish WebSocket connection
    ws.current = new WebSocket('ws://localhost:8081/ws');

    ws.current.onopen = () => {
      console.log('Connected to WebSocket server');
    };

    ws.current.onmessage = (event) => {
      const newMessage = JSON.parse(event.data);
      setMessages((prevMessages) => [...prevMessages, newMessage]);
    };

    ws.current.onclose = () => {
      console.log('Disconnected from WebSocket server');
    };

    return () => {
      ws.current.close(); // Clean up the WebSocket connection on component unmount
    };
  }, []); // Empty dependency array to run only once on mount

  const sendMessage = () => {
    if (sender.trim() === '' || messageContent.trim() === '') {
      alert('Please enter both sender name and message content.');
      return;
    }

    if (ws.current) {
      const message = {
        sender,
        content: messageContent
      };
      ws.current.send(JSON.stringify(message));
      setMessageContent(''); // Clear message input after sending
    }
  };

  return (
    <div className="p-4 border rounded max-w-md mx-auto mt-8">
      <h2 className="text-xl font-bold mb-4">WebSocket Client</h2>

      <div className="mb-4">
        <label className="block mb-1 font-medium">Sender Name</label>
        <input
          type="text"
          value={sender}
          onChange={(e) => setSender(e.target.value)}
          placeholder="Enter your name"
          className="w-full px-4 py-2 border rounded focus:outline-none focus:ring focus:border-blue-300"
        />
      </div>

      <div className="mb-4">
        <label className="block mb-1 font-medium">Message Content</label>
        <input
          type="text"
          value={messageContent}
          onChange={(e) => setMessageContent(e.target.value)}
          placeholder="Enter your message"
          className="w-full px-4 py-2 border rounded focus:outline-none focus:ring focus:border-blue-300"
        />
      </div>

      <button
        onClick={sendMessage}
        className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
      >
        Send Message
      </button>

      <ul className="space-y-2 mt-8">
        {messages.map((msg, index) => (
          <li key={index} className="border p-2 rounded">
            <strong>{msg.sender}:</strong> {msg.content} ({msg.time})
          </li>
        ))}
      </ul>
    </div>
  );
};

export default WebSocketClient;
