import Button from "@/components/ui/Button";
import { Close } from "@/components/ui/icons";
import clsx from "clsx";
import React, { useEffect } from "react";

const Modal = ({ size = "sm", className, children, onClose }) => {
  const sizes = {
    sm: "w-1/4 p-4",
    md: "w-1/2 p-6",
    lg: "w-3/4 p-8",
    xl: "w-[90%] p-8",
  };

  useEffect(() => {
    document.querySelector("body").classList.add("lock");

    return () => {
      document.querySelector("body").classList.remove("lock");
    };
  }, []);
  return (
    <div className="overlay flex justify-center items-center">
      <div className={clsx("relative bg-base-300 rounded-lg", sizes[size])}>
        <div className="absolute top-2 right-2">
          <Button
            size="sm"
            style="ghost"
            wide="square"
            className="hover:border-0 bg-transparent"
            onClick={onClose}
          >
            <Close />
          </Button>
        </div>
        <div className="">{children}</div>
        <div className="modal-action"></div>
      </div>
    </div>
  );
};

export default Modal;
