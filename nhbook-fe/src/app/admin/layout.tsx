import React from "react";
import { AppSidebar } from "@/components/layout/SideBar";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
export default function AdminLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <SidebarProvider>
      {/* Sidebar */}
      <AppSidebar />
      {/* Content */}
      <SidebarInset>
        <main className="flex-1 p-8">{children}</main>
      </SidebarInset>
    </SidebarProvider>
  );
}
