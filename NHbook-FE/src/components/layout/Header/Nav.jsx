import React from "react";
import { NavLink } from "react-router-dom";

const Nav = ({ links }) => {
  return (
    <div className="flex gap-10">
      {links.map((link, idx) => (
        <NavLink
          to={link.link}
          key={idx}
          className="text-primary-content text-base font-medium"
        >
          {link.lable}
        </NavLink>
      ))}
    </div>
  );
};

export default Nav;
