import apiInstance from '@/api/axios';
import AuthForm from '@/components/form/AuthForm';
import useAuthStore from '@/store/usAuthStore';
import React from 'react';
import { useNavigate } from 'react-router-dom';

const RegisterPage = () => {
  const naviagte = useNavigate();
  const { setDataUser } = useAuthStore();
  const handleRegister = async (data) => {
    try {
      const res = await apiInstance.post('/auth/register', data);
      console.log(data);
      setDataUser(res.data);
      naviagte('/');
    } catch (error) {
      console.log(error);
    }
  };
  return (
    <div className='w-max'>
      <AuthForm type='register' eventSubmit={handleRegister} />
    </div>
  );
};

export default RegisterPage;
