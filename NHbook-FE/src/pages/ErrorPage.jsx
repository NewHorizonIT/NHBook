import Button from "@/components/ui/Button";
import { Back } from "@/components/ui/icons";
import React from "react";
import { useNavigate } from "react-router-dom";

const ErrorPage = () => {
  const navigate = useNavigate();
  const handleGoToBack = () => {
    navigate(-1);
  };
  return (
    <div className="w-screen h-screen bg-base-100 flex justify-center items-center">
      <Button color="primary" size="md" onClick={handleGoToBack}>
        <Back />
        <span>Về trang trước đó</span>
      </Button>
    </div>
  );
};

export default ErrorPage;
