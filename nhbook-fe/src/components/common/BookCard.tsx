// components/BookCard.tsx
import React from "react";
import Image from "next/image";
import { Button } from "../ui/button";
import { Card, CardContent, CardFooter, CardHeader } from "../ui/card";
import { Plus } from "lucide-react";

export interface BookCardProps {
  title: string;
  author: string;
  price: number;
  image: string;
  salePrice?: number;
  onAddToCart?: () => void;
}

const BookCard: React.FC<BookCardProps> = ({
  title,
  author,
  price,
  image,
  salePrice,
  onAddToCart,
}) => {
  return (
    <Card>
      <CardHeader className="flex justify-center items-center">
        <Image
          src={image}
          alt={title}
          height={200}
          width={200}
          className="object-cover object-center"
        />
      </CardHeader>

      <CardContent>
        <h3 className="text-xl font-semibold min-h-[70px]">{title}</h3>
        <p className="text-bas mb-1">{author}</p>

        <div className="flex items-center gap-2 mb-2">
          {salePrice ? (
            <>
              <span className="text-xl font-bold text-red-500">
                {salePrice.toLocaleString()}₫
              </span>
              <span className="text-base line-through text-gray-400">
                {price.toLocaleString()}₫
              </span>
            </>
          ) : (
            <span className="text-xl font-bold text-red-500">
              {price.toLocaleString()}₫
            </span>
          )}
        </div>
      </CardContent>
      <CardFooter>
        <Button
          onClick={onAddToCart}
          className="w-full text-xl cursor-pointer"
          variant="default"
        >
          <Plus /> Thêm vào giỏ
        </Button>
      </CardFooter>
    </Card>
  );
};

export default BookCard;
