"use client";
import { demoBooks } from "@/mock/demo_book";
import { Button } from "@/components/ui/button";

export default function AdminBooks() {
  return (
    <div>
      <h1 className="text-2xl font-bold text-main mb-6">Quản lý sách</h1>
      <Button className="mb-4">+ Thêm sách mới</Button>
      <div className="overflow-x-auto">
        <table className="min-w-full border text-sm">
          <thead className="bg-main text-white">
            <tr>
              <th className="px-3 py-2 border">ID</th>
              <th className="px-3 py-2 border">Tên sách</th>
              <th className="px-3 py-2 border">Tác giả</th>
              <th className="px-3 py-2 border">Giá</th>
              <th className="px-3 py-2 border">Giá KM</th>
              <th className="px-3 py-2 border">Thao tác</th>
            </tr>
          </thead>
          <tbody>
            {demoBooks.map(book => (
              <tr key={book.id} className="even:bg-gray-50">
                <td className="px-3 py-2 border text-center">{book.id}</td>
                <td className="px-3 py-2 border">{book.title}</td>
                <td className="px-3 py-2 border">{book.author}</td>
                <td className="px-3 py-2 border">{book.price.toLocaleString()}₫</td>
                <td className="px-3 py-2 border">{book.salePrice ? book.salePrice.toLocaleString() + '₫' : '-'}</td>
                <td className="px-3 py-2 border text-center">
                  <Button size="sm" className="mr-2">Sửa</Button>
                  <Button size="sm" variant="neutral">Xóa</Button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
