import React from "react";
import { Outlet } from "react-router-dom";

const AuthLayout = () => {
  return (
    <div className="flex justify-center items-center px-3 h-screen">
      <Outlet />
    </div>
  );
};

export default AuthLayout;
