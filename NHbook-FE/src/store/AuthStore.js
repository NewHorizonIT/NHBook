import { create } from "zustand";

const useAuthStore = create((set) => ({
  isAuth: false,
  setIsAuth: (val) => set(() => ({ isAuth: val })),
}));

export default useAuthStore;
