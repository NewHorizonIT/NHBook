"use client";
import React, { useState } from "react";
import { Button } from "@/components/ui/button";

// Demo data
const orders = [
  {
    id: "DH001",
    date: "2025-09-20",
    status: "Đã giao hàng",
    total: 320000,
    items: [
      { name: "Lập Trình JavaScript Cơ Bản", qty: 1 },
      { name: "React & Next.js từ cơ bản đến nâng cao", qty: 2 },
    ],
  },
  {
    id: "DH002",
    date: "2025-08-15",
    status: "Đang xử lý",
    total: 150000,
    items: [{ name: "Lập Trình JavaScript Cơ Bản", qty: 1 }],
  },
];

export default function OrderPage() {
  const [filter, setFilter] = useState("tat-ca");
  const [showDetail, setShowDetail] = useState(
    null as null | (typeof orders)[0]
  );

  const filteredOrders =
    filter === "tat-ca"
      ? orders
      : orders.filter((o) => {
          if (filter === "dang-xu-ly") return o.status === "Đang xử lý";
          if (filter === "da-giao") return o.status === "Đã giao hàng";
          return true;
        });

  const handleCancel = (id: string) => {
    // Thực tế sẽ gọi API, ở đây chỉ alert demo
    alert(`Đã gửi yêu cầu hủy đơn ${id}`);
  };

  return (
    <div className="max-w-3xl mx-auto bg-white rounded-base shadow p-8 border border-border">
      <h1 className="text-2xl font-bold text-main mb-6">Đơn hàng của tôi</h1>
      {/* Filter trạng thái */}
      <div className="flex gap-2 mb-6">
        <Button
          variant={filter === "tat-ca" ? "default" : "neutral"}
          onClick={() => setFilter("tat-ca")}
        >
          Tất cả
        </Button>
        <Button
          variant={filter === "dang-xu-ly" ? "default" : "neutral"}
          onClick={() => setFilter("dang-xu-ly")}
        >
          Đang xử lý
        </Button>
        <Button
          variant={filter === "da-giao" ? "default" : "neutral"}
          onClick={() => setFilter("da-giao")}
        >
          Đã giao hàng
        </Button>
      </div>
      {filteredOrders.length === 0 ? (
        <div className="text-gray-500 text-center py-12">
          Không có đơn hàng phù hợp.
        </div>
      ) : (
        <div className="flex flex-col gap-6">
          {filteredOrders.map((order) => (
            <div key={order.id} className="border-b last:border-b-0 pb-4">
              <div className="flex flex-wrap items-center justify-between gap-2 mb-2">
                <div className="font-semibold text-main">
                  Mã đơn: {order.id}
                </div>
                <div className="text-gray-500 text-sm">
                  Ngày đặt: {order.date}
                </div>
                <div className="text-sm font-medium px-2 py-1 rounded-base bg-main/10 text-main">
                  {order.status}
                </div>
              </div>
              <div className="text-sm mb-2">
                Sản phẩm:{" "}
                {order.items.map((i) => `${i.name} (x${i.qty})`).join(", ")}
              </div>
              <div className="flex items-center justify-between gap-2 flex-wrap">
                <div className="font-bold text-lg text-main">
                  Tổng tiền: {order.total.toLocaleString()}₫
                </div>
                <div className="flex gap-2">
                  <Button
                    variant="neutral"
                    className="h-8 px-4 text-sm"
                    onClick={() => setShowDetail(order)}
                  >
                    Xem chi tiết
                  </Button>
                  {order.status === "Đang xử lý" && (
                    <Button
                      variant="neutral"
                      className="h-8 px-4 text-sm border border-red-400 text-red-500 hover:bg-red-50"
                      onClick={() => handleCancel(order.id)}
                    >
                      Hủy đơn
                    </Button>
                  )}
                </div>
              </div>
            </div>
          ))}
        </div>
      )}

      {/* Popup chi tiết đơn */}
      {showDetail && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/40">
          <div className="bg-white rounded-base shadow-lg p-8 min-w-[320px] max-w-md w-full relative">
            <button
              className="absolute top-2 right-2 text-gray-400 hover:text-main text-xl"
              onClick={() => setShowDetail(null)}
              aria-label="Đóng"
            >
              ×
            </button>
            <h2 className="text-xl font-bold text-main mb-4">
              Chi tiết đơn hàng {showDetail.id}
            </h2>
            <div className="mb-2 text-sm text-gray-600">
              Ngày đặt: {showDetail.date}
            </div>
            <div className="mb-2 text-sm text-gray-600">
              Trạng thái: {showDetail.status}
            </div>
            <div className="mb-2 text-sm text-gray-600">
              Tổng tiền:{" "}
              <span className="font-bold text-main">
                {showDetail.total.toLocaleString()}₫
              </span>
            </div>
            <div className="mb-2 text-sm text-gray-600">Sản phẩm:</div>
            <ul className="mb-4 pl-4 list-disc text-sm">
              {showDetail.items.map((item, idx) => (
                <li key={idx}>
                  {item.name} <span className="text-gray-400">x{item.qty}</span>
                </li>
              ))}
            </ul>
            <Button
              variant="neutral"
              className="w-full"
              onClick={() => setShowDetail(null)}
            >
              Đóng
            </Button>
          </div>
        </div>
      )}
    </div>
  );
}
