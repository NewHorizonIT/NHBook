import { Header } from "@/components/layout/Header";
import React from "react";
import { Outlet } from "react-router-dom";

const DefaultLayout = () => {
  return (
    <div>
      <Header />
      <Outlet />
      <div>Footer</div>
    </div>
  );
};

export default DefaultLayout;
