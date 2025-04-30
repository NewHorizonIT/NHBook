import { ListBook } from "@/components/common/Books";
import { Container } from "@/components/layout/Containner";
import { SideBar } from "@/components/layout/SideBar";
import React from "react";

const BooksPage = () => {
  return (
    <Container>
      <div className="flex gap-5">
        <SideBar />
        <div className="flex-1">
          <ListBook />
        </div>
      </div>
    </Container>
  );
};

export default BooksPage;
