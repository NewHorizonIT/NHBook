"use client";
import { Button } from "@/components/ui/button";

const orders = [
  { id: "DH001", user: "Nguyễn Văn A", total: 320000, status: "Đã giao hàng" },
  { id: "DH002", user: "Trần Thị B", total: 150000, status: "Đang xử lý" },
];

export default function AdminOrders() {
  return (
    <div>
      <h1 className="text-2xl font-bold text-main mb-6">Quản lý đơn hàng</h1>
      <div className="overflow-x-auto">
        <table className="min-w-full border text-sm">
          <thead className="bg-main text-white">
            <tr>
              <th className="px-3 py-2 border">Mã đơn</th>
              <th className="px-3 py-2 border">Khách hàng</th>
              <th className="px-3 py-2 border">Tổng tiền</th>
              <th className="px-3 py-2 border">Trạng thái</th>
              <th className="px-3 py-2 border">Thao tác</th>
            </tr>
          </thead>
          <tbody>
            {orders.map(order => (
              <tr key={order.id} className="even:bg-gray-50">
                <td className="px-3 py-2 border text-center">{order.id}</td>
                <td className="px-3 py-2 border">{order.user}</td>
                <td className="px-3 py-2 border">{order.total.toLocaleString()}₫</td>
                <td className="px-3 py-2 border">{order.status}</td>
                <td className="px-3 py-2 border text-center">
                  <Button size="sm" className="mr-2">Xem</Button>
                  <Button size="sm" variant="neutral">Cập nhật</Button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
