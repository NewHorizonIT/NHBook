import React from "react";

const Container = ({ children }) => {
  return (
    <div className="w-full px-3 py-2">
      <div className="w-full md:max-w-[1080px] mx-auto bg-transparent">
        {children}
      </div>
    </div>
  );
};

export default Container;
