import ToggleTheme from "@/components/common/ToggleTheme";
import { Button } from "@/components/ui/button";
import { User } from "lucide-react";
import Link from "next/link";
import React from "react";

const NavLinks = [
  { name: "Trang chủ", href: "/" },
  { name: "Sách", href: "/books" },
  { name: "Giới thiệu", href: "/about" },
];

const Header = () => {
  const isLoggedIn = true; // Replace with actual authentication logic
  return (
    <div className="shadow-sm h-16 flex items-center justify-between px-6">
      {/* LOGO */}
      <div className="text-lg font-bold text-blue-600">NHBook</div>
      {/* Navigation */}
      <div className="flex gap-6">
        {NavLinks.map((link) => (
          <Link
            key={link.name}
            href={link.href}
            className="hover:text-secondary transition-colors"
          >
            {link.name}
          </Link>
        ))}
      </div>
      {/* User Profile / Login */}
      <div className="flex items-center gap-4">
        <ToggleTheme />
        {isLoggedIn ? (
          <Link
            href="/profile"
            className="hover:text-secondary transition-colors"
          >
            <User />
          </Link>
        ) : (
          <div>
            <Link href="/login">
              <Button variant="default" className="cursor-pointer">
                Đăng nhập
              </Button>
            </Link>
            <Link href="/register" className="ml-2">
              <Button>Đăng ký</Button>
            </Link>
          </div>
        )}
      </div>
    </div>
  );
};

export default Header;
