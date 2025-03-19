import Container from "@/components/layout/Containner";
import { ThemeToggle } from "@/components/layout/Header";
import React from "react";

const Header = () => {
  return (
    <Container>
      <div className="flex justify-between items-center w-full h-full">
        {/* Left Header */}
        <p className="text-2xl font-playwrite font-bold text-primary-content">
          NHBOOK
        </p>
        {/* Center Header */}
        {/* Right Header */}
        <ThemeToggle />
      </div>
    </Container>
  );
};

export default Header;
