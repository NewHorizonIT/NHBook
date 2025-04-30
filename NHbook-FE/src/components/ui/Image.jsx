import React from "react";

const Image = ({ src, alt, fallback, className }) => {
  const handleError = (e) => {
    e.target.src = fallback;
  };

  return (
    <img src={src} alt={alt} onError={handleError} className={className} />
  );
};

export default Image;
