import apiInstance from "@/api/axios";
import { Container } from "@/components/layout/Containner";
import Button from "@/components/ui/Button";
import { Back } from "@/components/ui/icons";
import Input from "@/components/ui/Input";
import Lable from "@/components/ui/Lable";
import formatDate from "@/utils/formatDate";
import { useQuery } from "@tanstack/react-query";
import React from "react";
import { useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

const BookDetail = () => {
  const { bookID } = useParams();
  const navigate = useNavigate();
  const fetchBook = async () => {
    try {
      const res = await apiInstance.get(`/books/${bookID}`);
      console.log(res);
      return res;
    } catch (error) {
      console.log("[ERR] :: ", error);
    }
  };
  const {
    data: book,
    loading,
    error,
  } = useQuery({
    queryKey: ["book", bookID],
    queryFn: () => fetchBook(),
  });

  const [quantity, setQuantity] = useState(1);

  const handleAddToCart = () => {
    // Xử lý thêm vào giỏ hàng
    console.log(`Thêm ${quantity} bản ${book.title} vào giỏ hàng`);
  };

  if (loading) {
    return <span className="loading loading-spinner loading-xl"></span>;
  }

  return (
    <Container>
      <div className="w-full px-6 grid grid-cols-1 md:grid-cols-2 gap-10 relative">
        <Button className="absolute top-0 right-0" onClick={() => navigate(-1)}>
          <Back />
        </Button>
        {/* Cột trái: Hình ảnh sách */}
        <div className="flex justify-center">
          <img
            src={book?.thumbnail}
            alt={book?.title}
            className="w-full max-w-sm object-cover rounded shadow"
          />
        </div>
        {/* Cột phải: Thông tin sách và hành động */}
        <div className="space-y-4">
          <h1 className="text-3xl font-bold">{book?.title}</h1>
          <p className="text-base-content">
            Tác giả:{" "}
            {book?.authors.map((author) => (
              <span className="font-medium" key={author.id}>
                {author.name}
              </span>
            ))}
          </p>
          <p className="text-base-content">
            Thể loại:
            <span className="font-medium">{book?.category?.name}</span>
          </p>
          <p className="text-red-600 font-semibold text-xl">
            Giá: {book?.price.toLocaleString()} VND
          </p>
          <p className="text-base-content">
            Ngày xuất bản: {formatDate(book?.publishedAt)}
          </p>
          <p className="text-base-content">Còn lại: {book?.stock} quyển</p>
          <p className="text-base-content">Mô tả: {book?.description}</p>
          {/* Action: Số lượng và nút thêm vào giỏ hàng */}
          <div className="mt-6 flex items-center gap-4">
            <Lable lable="Số lượng:" />
            <Input
              type="number"
              props={{ min: 0, value: quantity }}
              onChange={(e) => setQuantity(e.target.value)}
              className="max-w-20"
            />
          </div>
          <Button color="primary">Thêm vào giỏ hàng</Button>
        </div>
      </div>
    </Container>
  );
};

export default BookDetail;
