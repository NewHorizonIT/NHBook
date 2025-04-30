import Footer from "@/components/layout/Footer/Footer";
import { Header } from "@/components/layout/Header";
import React from "react";
import { Outlet } from "react-router-dom";

const DefaultLayout = () => {
  return (
    <div>
      <Header />
      <div className="py-32 min-h-[100vh]">
        <Outlet />
      </div>
      <Footer />
    </div>
  );
};

export default DefaultLayout;
