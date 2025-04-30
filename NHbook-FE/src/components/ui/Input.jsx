import clsx from "clsx";
import React from "react";

const Input = ({
  size = "xs",
  className,
  type = "text",
  style = "neutral",
  placeholder = "",
  onChange,
  props,
}) => {
  const styles = {
    info: "input-info",
    success: "input-success",
    warning: "input-warning",
    error: "input-error",
    primary: "input-primary",
    secondary: "input-secondary",
    accent: "input-accent",
    neutral: "input-neutral",
  };
  const sizes = {
    xs: "input-xs",
    sm: "input-sm",
    md: "input-md",
    lg: "input-lg",
    xl: "input-xl",
  };
  return (
    <input
      type={type}
      placeholder={placeholder}
      onChange={onChange}
      className={clsx(
        "input",
        styles[style],
        sizes[size],
        className && className
      )}
      {...props}
    />
  );
};

export default Input;
