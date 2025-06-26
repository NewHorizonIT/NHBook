import apiInstance from '@/api/axios';
import AuthForm from '@/components/form/AuthForm';
import useAuthStore from '@/store/usAuthStore';
import React from 'react';
import { useNavigate } from 'react-router-dom';

const LoginPage = () => {
  const naviagte = useNavigate();
  const { setIsAuth, setDataUser } = useAuthStore();
  const handleLogin = async (data) => {
    try {
      const res = await apiInstance.post('/auth/login', data);
      console.log(res);
      setDataUser(res);
      setIsAuth(true);
      naviagte('/');
    } catch (error) {
      console.log(error);
    }
  };
  return (
    <div className='w-max'>
      <AuthForm type='login' eventSubmit={(data) => handleLogin(data)} />
    </div>
  );
};

export default LoginPage;
