import Button from "@/components/ui/Button";
import Image from "@/components/ui/Image";
import formatPrice from "@/utils/formatPrice";
import React from "react";
import { NavLink, useNavigate } from "react-router-dom";

const CardBook = ({ book, className }) => {
  const {
    id,
    title,
    thumbnail,
    price,
    description,
    stock,
    category,
    authors,
    publishedAt,
  } = book;

  return (
    <div className="card bg-base-200 shadow-sm p-5">
      <figure>
        <Image
          src={thumbnail}
          alt="anh sach"
          fallback="https://placehold.co/600x400"
          className="max-h-[200px] h-[200px]"
        />
      </figure>
      <div className="card-body">
        <h2 className="card-title">
          {title}
          <div className="badge badge-secondary">NEW</div>
        </h2>
        <p className="text-primary text-2xl">{formatPrice(price)}</p>

        <div className="card-actions flex flex-col gap-4 items-center">
          <Button wide="wide" color="primary" size="md">
            <NavLink to={`/books/${id}`}>Xem chi tiết</NavLink>
          </Button>
          <Button wide="wide" color="primary" size="md">
            Thêm vào giỏ hàng
          </Button>
        </div>
      </div>
    </div>
  );
};

export default CardBook;
