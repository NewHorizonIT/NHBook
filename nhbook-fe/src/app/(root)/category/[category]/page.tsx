"use client";
import React, { useState } from "react";
import { useParams } from "next/navigation";
import BookCard from "@/components/common/BookCard";
import { demoBooks } from "@/mock/demo_book";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

const sortOptions = [
  { value: "moi-nhat", label: "Mới nhất" },
  { value: "gia-tang", label: "Giá tăng dần" },
  { value: "gia-giam", label: "Giá giảm dần" },
];

export default function CategoryDetailPage() {
  const params = useParams();
  const rawCategory = params.category;
  const category = decodeURIComponent(
    Array.isArray(rawCategory) ? rawCategory[0] : rawCategory || ""
  );
  const [search, setSearch] = useState("");
  const [sort, setSort] = useState("moi-nhat");

  // Lọc sách theo tìm kiếm và danh mục động
  const filteredBooks = demoBooks.filter((book) => {
    const matchTitle = book.title.toLowerCase().includes(search.toLowerCase());
    // Giả lập: kiểm tra category có trong title (có thể thay bằng field category thực tế)
    const matchCategory =
      category === "tat-ca" ||
      book.title.toLowerCase().includes(category.toLowerCase());
    return matchTitle && matchCategory;
  });

  // Sắp xếp sách
  const sortedBooks = [...filteredBooks].sort((a, b) => {
    if (sort === "gia-tang")
      return (a.salePrice || a.price) - (b.salePrice || b.price);
    if (sort === "gia-giam")
      return (b.salePrice || b.price) - (a.salePrice || a.price);
    return b.id - a.id; // mới nhất
  });

  return (
    <div className="max-w-7xl mx-auto py-8 px-4">
      <h1 className="text-3xl font-bold text-main mb-6 capitalize">
        Danh mục: {category.replace(/-/g, " ")}
      </h1>
      {/* Tìm kiếm nâng cao */}
      <div className="bg-white rounded-base shadow p-4 mb-8 border border-border flex flex-col md:flex-row md:items-end gap-4">
        <div className="flex-1">
          <label className="block font-medium mb-1 text-main">Tìm kiếm</label>
          <Input
            placeholder="Nhập tên sách, tác giả..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="w-full"
          />
        </div>
        <div>
          <label className="block font-medium mb-1 text-main">Sắp xếp</label>
          <select
            className="w-full px-3 py-2 border rounded-base focus:ring-2 focus:ring-main outline-none"
            value={sort}
            onChange={(e) => setSort(e.target.value)}
          >
            {sortOptions.map((opt) => (
              <option key={opt.value} value={opt.value}>
                {opt.label}
              </option>
            ))}
          </select>
        </div>
        <Button className="mt-6 md:mt-0 h-10">Tìm kiếm</Button>
      </div>
      {/* Danh sách sách */}
      <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-6">
        {sortedBooks.length === 0 ? (
          <div className="col-span-full text-center text-gray-500 py-12">
            Không tìm thấy sách phù hợp.
          </div>
        ) : (
          sortedBooks.map((book) => <BookCard key={book.id} {...book} />)
        )}
      </div>
    </div>
  );
}
