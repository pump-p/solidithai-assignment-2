import React, { useState } from 'react';
import { queryLogs } from '../api/logApi';

const LogsPage = () => {
  const [query, setQuery] = useState('');
  const [logs, setLogs] = useState([]);
  const [error, setError] = useState(null);

  const handleSearch = async () => {
    try {
      const response = await queryLogs(query);
      setLogs(response.data);
      setError(null);
    } catch (err) {
      setError('Failed to retrieve logs.');
      setLogs([]);
    }
  };

  return (
    <div className="container mx-auto mt-8">
      <h2 className="text-2xl font-bold mb-4">Logs Viewer</h2>
      <div className="flex mb-4">
        <input
          type="text"
          placeholder="Search logs..."
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          className="w-full px-4 py-2 border rounded"
        />
        <button
          onClick={handleSearch}
          className="bg-blue-500 text-white px-4 py-2 rounded ml-2"
        >
          Search
        </button>
      </div>
      {error && <p className="text-red-500">{error}</p>}
      <ul className="space-y-4">
        {logs.map((log, index) => (
          <li key={index} className="border p-4 rounded">
            <strong>{log.sender}:</strong> {log.content} <br />
            <em>{log.time}</em>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default LogsPage;
