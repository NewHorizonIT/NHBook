import { Card } from "@/components/ui/card";
import Image from "next/image";
import React from "react";

const AuthLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="w-screen h-screen flex items-center justify-center relative">
      <Image
        src="/bg-auth.jpg"
        alt="Auth Background"
        layout="fill"
        objectFit="cover"
        className="absolute z-0"
      />
      <Card className="z-10 w-[400px] bg-black/50 backdrop-blur-sm border-0 text-white px-10 py-12">
        {children}
      </Card>
    </div>
  );
};

export default AuthLayout;
