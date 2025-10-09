"use client";
import React from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";

const sidebarLinks = [
  { href: "/account/profile", label: "Thông tin cá nhân" },
  { href: "/account/order", label: "Đơn hàng của tôi" },
  { href: "/account/change-password", label: "Đổi mật khẩu" },
  { href: "/account/logout", label: "Đăng xuất" },
];

export default function AccountLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const pathname = usePathname();
  return (
    <div className="max-w-7xl mx-auto py-8 px-4 flex gap-8 min-h-[60vh]">
      {/* Sidebar */}
      <aside className="w-64 bg-white rounded-base shadow border border-border p-6 flex flex-col gap-2 h-fit sticky top-24">
        <h2 className="text-xl font-bold text-main mb-4">Tài khoản</h2>
        <nav className="flex flex-col gap-1">
          {sidebarLinks.map((link) => (
            <Link
              key={link.href}
              href={link.href}
              className={`px-3 py-2 rounded-base font-medium transition-colors ${
                pathname === link.href
                  ? "bg-main text-white"
                  : "text-main hover:bg-main/10"
              }`}
            >
              {link.label}
            </Link>
          ))}
        </nav>
      </aside>
      {/* Content */}
      <main className="flex-1 min-w-0">{children}</main>
    </div>
  );
}
