import {
  RevenueBarChart,
  OrderStatusPieChart,
  TopBooksBarChart,
  RevenueLineChart,
} from "@/components/ui/charts-bookstore";

const revenueData = [
  { month: "T1", revenue: 12000000 },
  { month: "T2", revenue: 15000000 },
  { month: "T3", revenue: 18000000 },
  { month: "T4", revenue: 14000000 },
  { month: "T5", revenue: 20000000 },
  { month: "T6", revenue: 17000000 },
];

const orderStatusData = [
  { name: "Đã giao", value: 120 },
  { name: "Đang xử lý", value: 30 },
  { name: "Đã hủy", value: 10 },
];

const topBooks = [
  { name: "JS Cơ Bản", sold: 120 },
  { name: "Next.js Nâng Cao", sold: 90 },
  { name: "Python 101", sold: 70 },
  { name: "Kỹ Năng Sống", sold: 60 },
  { name: "Văn Học Việt", sold: 50 },
];

export default function AdminPage() {
  return (
    <div className="max-w-7xl mx-auto py-8 px-4">
      <h1 className="text-3xl font-bold text-main mb-8">Thống kê tổng quan</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-8 mb-8">
        <div className="bg-white rounded-base shadow p-6 border border-border">
          <h2 className="font-bold text-main mb-4">Doanh thu theo tháng</h2>
          <RevenueBarChart data={revenueData} />
        </div>
        <div className="bg-white rounded-base shadow p-6 border border-border">
          <h2 className="font-bold text-main mb-4">Tỉ lệ trạng thái đơn hàng</h2>
          <OrderStatusPieChart data={orderStatusData} />
        </div>
      </div>
      <div className="bg-white rounded-base shadow p-6 border border-border mb-8">
        <h2 className="font-bold text-main mb-4">Top 5 sách bán chạy</h2>
        <TopBooksBarChart data={topBooks} />
      </div>
      <div className="bg-white rounded-base shadow p-6 border border-border">
        <h2 className="font-bold text-main mb-4">Tăng trưởng doanh thu</h2>
        <RevenueLineChart data={revenueData} />
      </div>
    </div>
  );
}
"use client";

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { TrendingUp } from "lucide-react";
import { CartesianGrid, LabelList, Line, LineChart, XAxis } from "recharts";

export default function AdminDashboard() {
  const chartData = [
    { month: "January", desktop: 186, mobile: 80 },
    { month: "February", desktop: 305, mobile: 200 },
    { month: "March", desktop: 237, mobile: 120 },
    { month: "April", desktop: 73, mobile: 190 },
    { month: "May", desktop: 209, mobile: 130 },
    { month: "June", desktop: 214, mobile: 140 },
  ];

  const chartConfig = {
    desktop: {
      label: "Desktop",
      color: "var(--chart-1)",
    },
    mobile: {
      label: "Mobile",
      color: "var(--chart-2)",
    },
  } satisfies ChartConfig;

  return (
    <div>
      <h1 className="text-2xl font-bold text-main mb-6">Tổng quan quản trị</h1>
      <div className="grid grid-cols-2 md:grid-cols-4 gap-6 mb-8">
        <div className="bg-white rounded-base shadow p-6 border border-border text-center">
          <div className="text-3xl font-bold text-main">120</div>
          <div className="text-gray-500 mt-2">Sách</div>
        </div>
        <div className="bg-white rounded-base shadow p-6 border border-border text-center">
          <div className="text-3xl font-bold text-main">35</div>
          <div className="text-gray-500 mt-2">Đơn hàng mới</div>
        </div>
        <div className="bg-white rounded-base shadow p-6 border border-border text-center">
          <div className="text-3xl font-bold text-main">15</div>
          <div className="text-gray-500 mt-2">Người dùng mới</div>
        </div>
        <div className="bg-white rounded-base shadow p-6 border border-border text-center">
          <div className="text-3xl font-bold text-main">8</div>
          <div className="text-gray-500 mt-2">Danh mục</div>
        </div>
      </div>
      <div className="bg-white rounded-base shadow p-6 border border-border">
        <h2 className="text-lg font-bold text-main mb-4">Thống kê nhanh</h2>
        <div className="text-gray-600">
          <Card className="bg-secondary-background text-foreground">
            <CardHeader>
              <CardTitle>Line Chart - Label</CardTitle>
              <CardDescription>January - June 2024</CardDescription>
            </CardHeader>
            <CardContent>
              <ChartContainer
                className="[&_.recharts-layer_path]:stroke-black [&_.recharts-layer_path]:dark:stroke-white"
                config={chartConfig}
              >
                <LineChart
                  accessibilityLayer
                  data={chartData}
                  margin={{
                    top: 20,
                    left: 12,
                    right: 12,
                  }}
                >
                  <CartesianGrid vertical={false} />
                  <XAxis
                    dataKey="month"
                    tickLine={false}
                    axisLine={false}
                    tickMargin={8}
                    tickFormatter={(value) => value.slice(0, 3)}
                  />
                  <ChartTooltip
                    cursor={false}
                    content={<ChartTooltipContent indicator="line" />}
                  />
                  <Line
                    dataKey="desktop"
                    type="natural"
                    stroke="var(--color-desktop)"
                    strokeWidth={2}
                    dot={{
                      fill: "var(--color-desktop)",
                    }}
                    activeDot={{
                      r: 6,
                    }}
                  >
                    <LabelList
                      position="top"
                      offset={12}
                      className="fill-foreground"
                      fontSize={12}
                    />
                  </Line>
                </LineChart>
              </ChartContainer>
            </CardContent>
            <CardFooter className="flex-col items-start gap-2 text-sm">
              <div className="flex gap-2 leading-none font-medium">
                Trending up by 5.2% this month{" "}
                <TrendingUp className="h-4 w-4" />
              </div>
              <div className="text-muted-foreground leading-none">
                Showing total visitors for the last 6 months
              </div>
            </CardFooter>
          </Card>
        </div>
      </div>
    </div>
  );
}
