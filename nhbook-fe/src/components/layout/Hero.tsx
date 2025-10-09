"use client";

import { motion } from "framer-motion";
import React from "react";
import Image from "next/image"; // nếu bạn dùng Next.js, nếu dùng CRA thì import `<img>` bình thường
import { Button } from "@/components/ui/button"; // giả sử bạn có component Button từ shadcn
import { Input } from "../ui/input";

interface HeroBookProps {
  title: string;
  subtitle: string;
  description: string;
  imageSrc: string;
}

export function HeroBook({
  title,
  subtitle,
  description,
  imageSrc,
}: HeroBookProps) {
  return (
    <div className="relative overflow-hidden bg-secondary-background dark:bg-gray-900">
      <div className="container mx-auto px-6 lg:px-8 py-20 flex flex-col-reverse lg:flex-row items-center gap-10">
        {/* Nội dung văn bản */}
        <div className="w-full lg:w-1/2 space-y-6">
          <motion.span
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.2 }}
            className="text-sm font-medium text-blue-600 dark:text-blue-400 uppercase"
          >
            {subtitle}
          </motion.span>

          <motion.h1
            initial={{ opacity: 0, y: 30 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.4 }}
            className="text-4xl lg:text-5xl font-bold text-gray-900 dark:text-gray-100 leading-tight"
          >
            {title}
          </motion.h1>

          <motion.p
            initial={{ opacity: 0, y: 30 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.6 }}
            className="text-lg text-gray-700 dark:text-gray-300 max-w-md"
          >
            {description}
          </motion.p>

          <motion.a
            initial={{ opacity: 0, y: 30 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.8 }}
          >
            <div className="flex gap-3">
              <Input
                autoFocus
                type="text"
                placeholder="Tìm kiếm sách, tác giả..."
                className="flex-1 px-4 py-3 outline-none bg-transparent text-main placeholder:text-gray-400 text-lg"
              />
              <Button type="submit" className="cursor-poniter">
                Tìm kiếm
              </Button>
            </div>
          </motion.a>
        </div>

        {/* Hình ảnh */}
        <motion.div
          initial={{ opacity: 0, x: 50 }}
          animate={{ opacity: 1, x: 0 }}
          transition={{ delay: 1 }}
          className="w-full lg:w-1/2 flex justify-center"
        >
          <div className="max-w-sm shadow-lg">
            <Image
              src={imageSrc}
              width={400}
              height={600}
              alt="Book Cover"
              className="rounded-lg object-cover"
            />
          </div>
        </motion.div>
      </div>
    </div>
  );
}
