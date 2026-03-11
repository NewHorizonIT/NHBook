"use client";
import { Button } from "@/components/ui/button";
import { Moon, Sun } from "lucide-react";
import { useTheme } from "next-themes";

const ToggleTheme = () => {
  const { theme, setTheme } = useTheme();
  const toggleTheme = () => {
    setTheme(theme === "light" ? "dark" : "light");
  };
  return (
    <Button onClick={toggleTheme} variant="ghost" className="rounded-full">
      {theme === "light" ? <Sun /> : <Moon />}
    </Button>
  );
};

export default ToggleTheme;
