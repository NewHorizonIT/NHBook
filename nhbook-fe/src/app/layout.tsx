import type React from "react";
import type { Metadata, Viewport } from "next";
import { Inter, Crimson_Pro } from "next/font/google";
import "./globals.css";
import { ThemeProvider } from "@/components/provider/ThemeProvider";

const inter = Inter({ subsets: ["latin"] });
const crimsonPro = Crimson_Pro({
  subsets: ["latin"],
  variable: "--font-serif",
});

export const metadata: Metadata = {
  title: "Bookstore - Discover Your Next Great Read",
  description: "Browse our curated collection of bestselling and indie books",
};

export const viewport: Viewport = {
  themeColor: "#000000",
  userScalable: true,
  width: "device-width",
  initialScale: 1,
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body className={`${inter.className} ${crimsonPro.variable}`}>
        <ThemeProvider
          attribute={"class"}
          defaultTheme="system"
          enableSystem
          disableTransitionOnChange
        >
          {children}
        </ThemeProvider>
      </body>
    </html>
  );
}
