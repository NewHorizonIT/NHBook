"use client";
import React from "react";
import Menu from "./Menu";
import Link from "next/link";
import { Button } from "../ui/button";
import useAuthStore from "@/store/useAuth";
import DropDownMenuAccount from "../common/DropDownMenu";
import Image from "next/image";
import { ShoppingCart } from "lucide-react";
import CartSheet from "../common/CartSheet";

const Header = () => {
  const { isLogin } = useAuthStore();

  return (
    <header className="sticky top-0 z-30 w-full bg-main shadow-lg shadow-main/20">
      <div className="mx-auto flex h-[64px] max-w-7xl items-center justify-between px-4 md:px-8 relative">
        {/* Logo */}
        <Link href="/" className="flex items-center gap-2 group">
          <Image
            width={32}
            height={32}
            src="/file.svg"
            alt="NHBook Logo"
            className="h-8 w-8 drop-shadow-md group-hover:scale-110 transition-transform"
          />
          <span className="text-2xl font-bold text-white tracking-wide group-hover:text-yellow-300 transition-colors">
            NHBook
          </span>
        </Link>

        {/* Center: Menu */}
        <nav className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 hidden lg:block">
          <Menu />
        </nav>

        {/* Right: Search + Cart + Auth */}
        <div className="flex items-center gap-4 ml-4">
          {/* Cart Icon */}
          <CartSheet cart={[{ id: 1, qty: 2 }]}>
            <div className="relative">
              <ShoppingCart
                size={26}
                className="text-white group-hover:text-yellow-300 transition-colors"
              />
              {/* Cart badge (fake count for demo) */}
              <span className="absolute -top-2 -right-2 bg-yellow-400 text-xs text-main-foreground rounded-full px-1.5 py-0.5 font-bold shadow">
                2
              </span>
            </div>
          </CartSheet>
          {/* Auth */}
          {isLogin ? (
            <DropDownMenuAccount />
          ) : (
            <div className="flex gap-2">
              <Button
                asChild
                className="min-w-[90px] bg-white text-main border-main hover:bg-main hover:text-white"
              >
                <Link href={"/login"}>Đăng nhập</Link>
              </Button>
              <Button
                asChild
                className="min-w-[90px] bg-yellow-400 text-main border-yellow-400 hover:bg-main hover:text-white"
              >
                <Link href={"/register"}>Đăng ký</Link>
              </Button>
            </div>
          )}
        </div>
      </div>
    </header>
  );
};

export default Header;
