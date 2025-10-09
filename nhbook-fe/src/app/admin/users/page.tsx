"use client";
import { Button } from "@/components/ui/button";

const users = [
  { id: 1, name: "Nguyễn Văn A", email: "a@gmail.com", role: "user" },
  { id: 2, name: "Trần Thị B", email: "b@gmail.com", role: "admin" },
];

export default function AdminUsers() {
  return (
    <div>
      <h1 className="text-2xl font-bold text-main mb-6">Quản lý người dùng</h1>
      <div className="overflow-x-auto">
        <table className="min-w-full border text-sm">
          <thead className="bg-main text-white">
            <tr>
              <th className="px-3 py-2 border">ID</th>
              <th className="px-3 py-2 border">Tên</th>
              <th className="px-3 py-2 border">Email</th>
              <th className="px-3 py-2 border">Quyền</th>
              <th className="px-3 py-2 border">Thao tác</th>
            </tr>
          </thead>
          <tbody>
            {users.map(user => (
              <tr key={user.id} className="even:bg-gray-50">
                <td className="px-3 py-2 border text-center">{user.id}</td>
                <td className="px-3 py-2 border">{user.name}</td>
                <td className="px-3 py-2 border">{user.email}</td>
                <td className="px-3 py-2 border">{user.role}</td>
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
