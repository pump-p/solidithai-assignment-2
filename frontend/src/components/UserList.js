import React, { useState, useEffect } from 'react';
import { deleteUser } from '../api/userApi';

const UserList = ({ users, searchQuery, onEdit }) => {
  const [filteredUsers, setFilteredUsers] = useState([]);

  // Filter users based on search query
  useEffect(() => {
    setFilteredUsers(
      users.filter(
        (user) =>
          user.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
          user.email.toLowerCase().includes(searchQuery.toLowerCase())
      )
    );
  }, [searchQuery, users]);

  const handleDelete = async (id) => {
    if (!window.confirm('Are you sure you want to delete this user?')) return;

    try {
      await deleteUser(id);
      // Handle deletion at the parent level (UsersPage)
    } catch {
      console.error('Failed to delete user.');
    }
  };

  return (
    <div>
      <h2 className="text-xl font-bold mb-4">Users</h2>
      <ul className="space-y-4">
        {filteredUsers.map((user) => (
          <li
            key={user.id}
            className="flex justify-between items-center p-4 border rounded"
          >
            <div>
              <p className="font-bold">{user.name}</p>
              <p>{user.email}</p>
            </div>
            <div className="space-x-2">
              <button
                onClick={() => onEdit(user)}
                className="bg-yellow-500 text-white px-4 py-2 rounded"
              >
                Edit
              </button>
              <button
                onClick={() => handleDelete(user.id)}
                className="bg-red-500 text-white px-4 py-2 rounded"
              >
                Delete
              </button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default UserList;
