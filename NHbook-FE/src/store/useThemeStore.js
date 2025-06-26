import { create } from "zustand";
import { persist } from "zustand/middleware";

const useThemeStore = create(
  persist(
    (set, get) => ({
      theme: "light",
      setTheme: () => {
        const newTheme = get().theme === "light" ? "dark" : "light";
        set({ theme: newTheme });
      },
    }),
    {
      name: "theme",
    }
  )
);

export default useThemeStore;
