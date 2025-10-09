"use client";

import { demoBooks } from "@/mock/demo_book";
import React from "react";
import BookCard from "./BookCard";
import PaginationCustom from "./PaginationCustom";

const ContainerBook = () => {
  const handleAddToCart = (title: string) => {
    alert(`Đã thêm sách "${title}" vào giỏ hàng`);
  };

  return (
    <div className="max-w-[1024px] py-10 mx-auto px-5">
      <div className="container mx-auto grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
        {demoBooks.map((book) => (
          <BookCard
            key={book.id}
            title={book.title}
            author={book.author}
            price={book.price}
            salePrice={book.salePrice}
            image={book.image}
            onAddToCart={() => handleAddToCart(book.title)}
          />
        ))}
      </div>
      <PaginationCustom currentPage={2} maxPage={10} />
    </div>
  );
};

export default ContainerBook;
