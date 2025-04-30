import { Container } from "@/components/layout/Containner";
import { ThemeToggle } from "@/components/layout/Header";
import Nav from "@/components/layout/Header/Nav";
import { linksNav } from "@/constants/links";

import React from "react";
import { NavLink } from "react-router-dom";

const Header = () => {
  return (
    <div className="bg-primary w-full fixed z-50">
      <Container>
        <div className="flex justify-between items-center w-full h-[65px] bg-primary">
          {/* Left Header */}
          <p className="text-2xl font-playwrite font-bold text-primary-content">
            NHBOOK
          </p>
          {/* Center Header */}
          <Nav links={linksNav} />
          {/* Right Header */}
          <div>
            <button className="button bg-primary text-primary-content min-w-32">
              <NavLink to="/auth/login">Đăng nhập</NavLink>
            </button>
            <button className="button bg-primary text-primary-content min-w-32">
              <NavLink to="/auth/register">Đăng kí</NavLink>
            </button>
          </div>
          <ThemeToggle />
        </div>
      </Container>
    </div>
  );
};

export default Header;
