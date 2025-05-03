import clsx from "clsx";
import React from "react";

const Lable = ({ type = "normal", lable, children, htmlFor, className }) => {
  const types = {
    floating: "floating-label",
    normal: "lable",
  };
  return (
    <label className={clsx(types[type], className)} htmlFor={htmlFor}>
      <span>{lable}</span>
      {children}
    </label>
  );
};

export default Lable;
