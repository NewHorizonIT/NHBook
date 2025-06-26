import clsx from "clsx";
import React from "react";

const Title = ({ className, children, size }) => {
  const sizes = {
    xs: "text-xs font-medium",
    sm: "text-sm font-medium",
    md: "text-md font-medium",
    lg: "text-lg font-medium",
    xl: "text-xl font-medium",
  };
  return <div className={clsx(size && sizes[size], className)}>{children}</div>;
};

export default Title;
