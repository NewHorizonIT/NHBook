"use client";
import React from "react";

export default function ContactPage() {
  return (
    <div className="max-w-2xl mx-auto py-12 px-4">
      <h1 className="text-3xl font-bold mb-6 text-main">Liên hệ với NHBook</h1>
      <p className="mb-6 text-lg text-gray-700">
        Nếu bạn có bất kỳ câu hỏi, góp ý hoặc cần hỗ trợ, vui lòng liên hệ với
        chúng tôi qua biểu mẫu dưới đây hoặc qua email/điện thoại.
      </p>
      <form className="bg-white rounded-base shadow p-6 flex flex-col gap-4 border border-border">
        <div>
          <label htmlFor="name" className="block font-medium mb-1 text-main">
            Họ và tên
          </label>
          <input
            id="name"
            name="name"
            type="text"
            required
            className="w-full px-4 py-2 border rounded-base focus:ring-2 focus:ring-main outline-none"
            placeholder="Nhập họ tên của bạn"
          />
        </div>
        <div>
          <label htmlFor="email" className="block font-medium mb-1 text-main">
            Email
          </label>
          <input
            id="email"
            name="email"
            type="email"
            required
            className="w-full px-4 py-2 border rounded-base focus:ring-2 focus:ring-main outline-none"
            placeholder="Nhập email của bạn"
          />
        </div>
        <div>
          <label htmlFor="message" className="block font-medium mb-1 text-main">
            Nội dung
          </label>
          <textarea
            id="message"
            name="message"
            rows={5}
            required
            className="w-full px-4 py-2 border rounded-base focus:ring-2 focus:ring-main outline-none"
            placeholder="Nhập nội dung liên hệ..."
          />
        </div>
        <button
          type="submit"
          className="mt-2 bg-main text-white font-bold py-2 rounded-base hover:bg-main/90 transition-colors"
        >
          Gửi liên hệ
        </button>
      </form>
      <div className="mt-8 text-gray-600">
        <div>
          Email:{" "}
          <a href="mailto:support@nhbook.vn" className="text-main underline">
            support@nhbook.vn
          </a>
        </div>
        <div>
          Điện thoại:{" "}
          <a href="tel:0123456789" className="text-main underline">
            0123 456 789
          </a>
        </div>
        <div>Địa chỉ: 123 Đường Sách, Quận 1, TP.HCM</div>
      </div>
    </div>
  );
}
