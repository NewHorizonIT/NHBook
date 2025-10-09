"use client";
import { Button } from "@/components/ui/button";

const categories = [
  { id: 1, name: "Lập trình" },
  { id: 2, name: "Kinh tế" },
  { id: 3, name: "Văn học" },
];

export default function AdminCategories() {
  return (
    <div>
      <h1 className="text-2xl font-bold text-main mb-6">Quản lý danh mục</h1>
      <Button className="mb-4">+ Thêm danh mục</Button>
      <div className="overflow-x-auto">
        <table className="min-w-full border text-sm">
          <thead className="bg-main text-white">
            <tr>
              <th className="px-3 py-2 border">ID</th>
              <th className="px-3 py-2 border">Tên danh mục</th>
              <th className="px-3 py-2 border">Thao tác</th>
            </tr>
          </thead>
          <tbody>
            {categories.map(c => (
              <tr key={c.id} className="even:bg-gray-50">
                <td className="px-3 py-2 border text-center">{c.id}</td>
                <td className="px-3 py-2 border">{c.name}</td>
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
