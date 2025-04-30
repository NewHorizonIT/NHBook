import clsx from "clsx";
import React from "react";

const Select = ({
  color,
  size = "sm",
  className,
  options = [],
  title,
  ghost = false,
}) => {
  const colors = {
    primary: "select-primary",
    secondary: "select-secondary",
    neutral: "select-neutral",
  };

  const sizes = {
    xs: "select-xs",
    sm: "select-sm",
    md: "select-md",
    lg: "select-lg",
    xl: "select-xl",
  };
  return (
    <select
      defaultValue={title}
      className={clsx(
        "select",
        !!color && colors[color],
        sizes[size],
        ghost && "select-ghost",
        !!className && className
      )}
    >
      {options.length !== 0 &&
        options.map((option) => (
          <option value={option.value}>{option.lable}</option>
        ))}
    </select>
  );
};

export default Select;
