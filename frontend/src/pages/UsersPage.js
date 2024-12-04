import React, { useState, useEffect } from 'react';
import UserList from '../components/UserList';
import UserForm from '../components/UserForm';
import { getUsers } from '../api/userApi';

const UsersPage = () => {
  const [userToEdit, setUserToEdit] = useState(null);
  const [isAddingUser, setIsAddingUser] = useState(false);
  const [users, setUsers] = useState([]);
  const [searchQuery, setSearchQuery] = useState('');

  // Fetch users from the backend
  const fetchUsers = async () => {
    try {
      const { data } = await getUsers();
      setUsers(data);
    } catch (error) {
      console.error('Failed to fetch users:', error);
    }
  };

  useEffect(() => {
    fetchUsers(); // Initial fetch on component mount
  }, []);

  const handleEdit = (user) => {
    setUserToEdit(user);
    setIsAddingUser(false);
  };

  const handleSave = () => {
    setUserToEdit(null);
    setIsAddingUser(false);
    fetchUsers(); // Refresh users after save
  };

  const handleAddUser = () => {
    setUserToEdit(null);
    setIsAddingUser(true);
  };

  return (
    <div className="container mx-auto mt-8">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">User Management</h1>
        <button
          onClick={handleAddUser}
          className="bg-green-500 text-white px-4 py-2 rounded"
        >
          Add User
        </button>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-8 mt-8">
        {isAddingUser && (
          <div className="border p-4 rounded">
            <h2 className="text-xl font-bold mb-4">Add User</h2>
            <UserForm userToEdit={null} onSave={handleSave} />
          </div>
        )}
        {userToEdit && (
          <div className="border p-4 rounded">
            <h2 className="text-xl font-bold mb-4">Edit User</h2>
            <UserForm userToEdit={userToEdit} onSave={handleSave} />
          </div>
        )}
      </div>
      <div className="mt-4">
        {/* Search Input */}
        <input
          type="text"
          placeholder="Search by name or email"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          className="w-full px-4 py-2 border rounded"
        />
      </div>
      <div className="mt-8">
        <UserList users={users} searchQuery={searchQuery} onEdit={handleEdit} />
      </div>
    </div>
  );
};

export default UsersPage;
