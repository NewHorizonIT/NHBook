import React from "react";
import { NavLink } from "react-router-dom";

const Navigation = () => {
  return (
    <div>
      <NavLink>Home</NavLink>
      <NavLink>About</NavLink>
      <NavLink>Category</NavLink>
    </div>
  );
};

export default Navigation;
