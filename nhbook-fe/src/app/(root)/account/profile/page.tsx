"use client";
import React, { useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";

export default function ProfilePage() {
  // Giả lập dữ liệu user
  const [user, setUser] = useState({
    name: "Nguyễn Văn A",
    email: "nguyenvana@gmail.com",
    phone: "0123456789",
    avatar: "https://i.pravatar.cc/150?img=3",
  });
  const [edit, setEdit] = useState(false);
  const [form, setForm] = useState(user);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSave = (e: React.FormEvent) => {
    e.preventDefault();
    setUser(form);
    setEdit(false);
  };

  return (
    <div className="max-w-xl mx-auto bg-white rounded-base shadow p-8 border border-border">
      <h1 className="text-2xl font-bold text-main mb-6">Thông tin cá nhân</h1>
      <div className="flex flex-col items-center gap-4 mb-6">
        <Avatar className="size-20">
          <AvatarImage src={user.avatar} alt={user.name} />
          <AvatarFallback>{user.name[0]}</AvatarFallback>
        </Avatar>
        <div className="text-lg font-semibold text-main">{user.name}</div>
      </div>
      {edit ? (
        <form onSubmit={handleSave} className="flex flex-col gap-4">
          <div>
            <label className="block font-medium mb-1 text-main">
              Họ và tên
            </label>
            <Input
              name="name"
              value={form.name}
              onChange={handleChange}
              required
            />
          </div>
          <div>
            <label className="block font-medium mb-1 text-main">Email</label>
            <Input
              name="email"
              value={form.email}
              onChange={handleChange}
              required
              type="email"
            />
          </div>
          <div>
            <label className="block font-medium mb-1 text-main">
              Số điện thoại
            </label>
            <Input
              name="phone"
              value={form.phone}
              onChange={handleChange}
              required
            />
          </div>
          <div className="flex gap-2 mt-2">
            <Button type="submit" className="bg-main text-white">
              Lưu
            </Button>
            <Button
              type="button"
              variant="neutral"
              onClick={() => setEdit(false)}
            >
              Huỷ
            </Button>
          </div>
        </form>
      ) : (
        <div className="flex flex-col gap-3">
          <div>
            <span className="font-medium text-main">Email:</span> {user.email}
          </div>
          <div>
            <span className="font-medium text-main">Số điện thoại:</span>{" "}
            {user.phone}
          </div>
          <Button className="mt-4 w-fit" onClick={() => setEdit(true)}>
            Chỉnh sửa
          </Button>
        </div>
      )}
    </div>
  );
}
