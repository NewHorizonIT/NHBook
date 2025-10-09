"use client";
import React from "react";
import {
  ChartContainer,
  ChartTooltip,
  ChartLegend,
} from "@/components/ui/chart";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  PieChart,
  Pie,
  Cell,
  LineChart,
  Line,
} from "recharts";

const COLORS = ["#4f46e5", "#facc15", "#ef4444", "#10b981", "#6366f1"];

// Chart: Doanh thu theo tháng
export function RevenueBarChart({
  data,
}: {
  data: { month: string; revenue: number }[];
}) {
  return (
    <ChartContainer
      config={{ revenue: { color: COLORS[0], label: "Doanh thu" } }}
    >
      <BarChart data={data} height={260}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="month" />
        <YAxis />
        <Tooltip content={<ChartTooltip />} />
        <Legend content={<ChartLegend />} />
        <Bar dataKey="revenue" fill={COLORS[0]} name="Doanh thu" />
      </BarChart>
    </ChartContainer>
  );
}

// Chart: Tỉ lệ trạng thái đơn hàng
export function OrderStatusPieChart({
  data,
}: {
  data: { name: string; value: number }[];
}) {
  return (
    <ChartContainer
      config={{
        giao: { color: COLORS[0], label: "Đã giao" },
        xuly: { color: COLORS[1], label: "Đang xử lý" },
        huy: { color: COLORS[2], label: "Đã hủy" },
      }}
    >
      <PieChart height={260}>
        <Pie
          data={data}
          dataKey="value"
          nameKey="name"
          cx="50%"
          cy="50%"
          outerRadius={80}
          label
        >
          {data.map((entry, idx) => (
            <Cell key={`cell-${idx}`} fill={COLORS[idx % COLORS.length]} />
          ))}
        </Pie>
        <Tooltip content={<ChartTooltip />} />
        <Legend content={<ChartLegend />} />
      </PieChart>
    </ChartContainer>
  );
}

// Chart: Top sách bán chạy
export function TopBooksBarChart({
  data,
}: {
  data: { name: string; sold: number }[];
}) {
  return (
    <ChartContainer
      config={{ sold: { color: COLORS[1], label: "Số lượng bán" } }}
    >
      <BarChart data={data} layout="vertical" height={260}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis type="number" />
        <YAxis dataKey="name" type="category" />
        <Tooltip content={<ChartTooltip />} />
        <Bar dataKey="sold" fill={COLORS[1]} name="Số lượng bán" />
      </BarChart>
    </ChartContainer>
  );
}

// Chart: Đường tăng trưởng doanh thu
export function RevenueLineChart({
  data,
}: {
  data: { month: string; revenue: number }[];
}) {
  return (
    <ChartContainer
      config={{ revenue: { color: COLORS[3], label: "Doanh thu" } }}
    >
      <LineChart data={data} height={260}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="month" />
        <YAxis />
        <Tooltip content={<ChartTooltip />} />
        <Legend content={<ChartLegend />} />
        <Line
          type="monotone"
          dataKey="revenue"
          stroke={COLORS[3]}
          name="Doanh thu"
        />
      </LineChart>
    </ChartContainer>
  );
}
