"use client";

import {
  Layers2,
  LayoutDashboard,
  LibraryBig,
  ListOrdered,
  Users,
} from "lucide-react";

import * as React from "react";

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarRail,
} from "@/components/ui/sidebar";
import Link from "next/link";
import { Button } from "../ui/button";

// This is sample data.
const adminLinks = [
  { href: "/admin/dashboard", label: "Tổng quan", icon: <LayoutDashboard /> },
  { href: "/admin/books", label: "Quản lý sách", icon: <LibraryBig /> },
  { href: "/admin/orders", label: "Quản lý đơn hàng", icon: <ListOrdered /> },
  { href: "/admin/users", label: "Quản lý người dùng", icon: <Users /> },
  { href: "/admin/categories", label: "Quản lý danh mục", icon: <Layers2 /> },
];

export function AppSidebar() {
  return (
    <Sidebar collapsible="icon">
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <div className="text-2xl text-center">NHBooks</div>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent className="p-4 list-none">
        {adminLinks.map((link) => (
          <SidebarMenuItem key={link.href} className="mb-4 cursor-pointer">
            <Link href={link.href}>
              <Button className="w-full flex gap-5 justify-start cursor-pointer">
                {link.icon}
                {link.label}
              </Button>
            </Link>
          </SidebarMenuItem>
        ))}
      </SidebarContent>
      <SidebarFooter>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton className="h-max">
              <Avatar>
                <AvatarImage
                  src="https://github.com/shadcn.png"
                  alt="@shadcn"
                />
                <AvatarFallback>CN</AvatarFallback>
              </Avatar>
              <div>
                <h1>Anh Quan</h1>
                <p>aquan@gmai.com</p>
              </div>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  );
}
