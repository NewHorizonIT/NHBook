import React from "react";
import { demoBooks } from "@/mock/demo_book";
import Image from "next/image";
import { Button } from "@/components/ui/button";

export default function BookDetailPage({
  params,
}: {
  params: Promise<{ slug: string }>;
}) {
  const book = demoBooks[0];

  if (!book)
    return (
      <div className="text-center py-12 text-gray-500">
        Không tìm thấy sách.
      </div>
    );

  return (
    <div className="max-w-5xl mx-auto py-10 px-4 grid grid-cols-1 md:grid-cols-2 gap-10">
      {/* Ảnh sách */}
      <div className="flex flex-col items-center">
        <div className="relative w-60 h-80 rounded-base overflow-hidden bg-gray-100 shadow">
          <Image
            src={book.image}
            alt={book.title}
            fill
            className="object-cover object-center"
          />
        </div>
      </div>
      {/* Thông tin sách */}
      <div className="flex flex-col gap-4">
        <h1 className="text-3xl font-bold text-main mb-2">{book.title}</h1>
        <div className="text-lg text-gray-600 mb-2">
          Tác giả:{" "}
          <span className="font-semibold text-main">{book.author}</span>
        </div>
        <div className="flex items-center gap-3 mb-2">
          {book.salePrice ? (
            <>
              <span className="text-2xl font-bold text-red-500">
                {book.salePrice.toLocaleString()}₫
              </span>
              <span className="text-base line-through text-gray-400">
                {book.price.toLocaleString()}₫
              </span>
            </>
          ) : (
            <span className="text-2xl font-bold text-main">
              {book.price.toLocaleString()}₫
            </span>
          )}
        </div>
        <div className="text-gray-700 mb-4">
          <p>
            Đây là mô tả ngắn về sách. Bạn có thể cập nhật mô tả chi tiết, số
            trang, nhà xuất bản, năm xuất bản, v.v. ở đây để khách hàng hiểu rõ
            hơn về sản phẩm.
          </p>
        </div>
        <Button className="w-full md:w-fit">Thêm vào giỏ hàng</Button>
      </div>
    </div>
  );
}
