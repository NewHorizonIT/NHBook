import ContainerBook from "@/components/common/ContainerBook";
import Footer from "@/components/layout/Footer";
import Header from "@/components/layout/Header";
import { HeroBook } from "@/components/layout/Hero";
import React from "react";

const MainLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="bg-secondary-background">
      <Header />
      <HeroBook
        title="Khám Phá Thế Giới Tri Thức"
        subtitle="Sách Hay Mỗi Ngày"
        description="Chọn từ hàng ngàn đầu sách với ưu đãi hấp dẫn. Đọc, học, và truyền cảm hứng mỗi ngày."
        imageSrc="/book_cover.webp" // đường dẫn ảnh bìa sách
      />
      {children}
      <ContainerBook />
      <Footer />
    </div>
  );
};

export default MainLayout;
