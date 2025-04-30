import apiInstance from "@/api/axios";
import AuthForm from "@/components/form/AuthForm";
import useAuthStore from "@/store/usAuthStore";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const LoginPage = () => {
  const naviagte = useNavigate();
  const { setIsAuth, setDataUser } = useAuthStore();
  const [error, setError] = useState("");
  const handleLogin = async (data) => {
    try {
      const res = await apiInstance.post("/auth/login", data);
      await setIsAuth(true);
    } catch (error) {
      
    }
  };
  return (
    <div className="w-max">
      <AuthForm type="login" eventSubmit={handleLogin} />
    </div>
  );
};

export default LoginPage;
