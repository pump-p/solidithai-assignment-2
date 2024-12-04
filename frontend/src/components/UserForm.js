import React, { useState, useEffect } from 'react';
import { createUser, updateUser } from '../api/userApi';
import { toast } from 'react-toastify';

const UserForm = ({ userToEdit, onSave }) => {
  const [user, setUser] = useState({ name: '', email: '', password: '' });

  useEffect(() => {
    if (userToEdit) {
      setUser({ name: userToEdit.name, email: userToEdit.email, password: '' });
    }
  }, [userToEdit]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setUser({ ...user, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      if (userToEdit) {
        await updateUser(userToEdit.id, user);
        toast.success('User updated successfully!');
      } else {
        await createUser(user);
        toast.success('User created successfully!');
      }
      onSave();
      setUser({ name: '', email: '', password: '' });
    } catch (err) {
      toast.error('Failed to save user.', {
        autoClose: 1000
      });
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="max-w-md mx-auto mt-8 p-4 border rounded"
    >
      <h2 className="text-xl font-bold mb-4">
        {userToEdit ? 'Edit User' : 'Add User'}
      </h2>
      <div className="mb-4">
        <label className="block mb-1">Name</label>
        <input
          type="text"
          name="name"
          value={user.name}
          onChange={handleChange}
          className="w-full px-3 py-2 border rounded"
        />
      </div>
      <div className="mb-4">
        <label className="block mb-1">Email</label>
        <input
          type="email"
          name="email"
          value={user.email}
          onChange={handleChange}
          className="w-full px-3 py-2 border rounded"
        />
      </div>
      <div className="mb-4">
        <label className="block mb-1">Password</label>
        <input
          type="password"
          name="password"
          value={user.password}
          onChange={handleChange}
          className="w-full px-3 py-2 border rounded"
        />
      </div>
      <button
        type="submit"
        className="bg-blue-500 text-white px-4 py-2 rounded"
      >
        {userToEdit ? 'Update User' : 'Create User'}
      </button>
    </form>
  );
};

export default UserForm;
