import clsx from "clsx";
import React from "react";

const Button = ({
  type,
  color,
  style,
  size,
  className,
  children,
  onClick,
  wide,
}) => {
  const colors = {
    neutral: "btn-neutral",
    primary: "btn-primary",
    secondary: "btn-secondary",
    accent: "btn-accent",
    info: "btn-info",
    success: "btn-success",
    warning: "btn-warning",
    error: "btn-error",
  };
  const styles = {
    outline: "btn-outline",
    dash: "btn-dash",
    soft: "btn-soft",
    ghost: "btn-ghost",
    link: "btn-link",
  };

  const sizes = {
    xs: "btn-xs",
    sm: "btn-sm",
    md: "btn-md",
    lg: "btn-lg",
    xl: "btn-xl",
  };

  const wides = {
    wide: "btn-wide",
    block: "btn-block",
    square: "btn-square",
    circle: "btn-circle",
  };
  return (
    <button
      type={type}
      onClick={onClick}
      className={clsx(
        "btn",
        color && colors[color],
        style && styles[style],
        wide && wides[wide],
        size && sizes[size],
        className
      )}
    >
      {children}
    </button>
  );
};

export default Button;
