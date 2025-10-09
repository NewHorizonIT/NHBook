import { create } from "zustand";

type AuthState = {
  isLogin: boolean;
  token: string | null;
  user: Record<string, number | string> | null;
};

type AuthAction = {
  setToken: (token: string) => void;
  setUser: (user: Record<string, number> | null) => void;
  setIsLogin: () => void;
  logout: () => void;
};

const useAuthStore = create<AuthState & AuthAction>((set) => ({
  isLogin: true,
  token: null,
  user: null,
  setIsLogin: () => set({ isLogin: true }),
  setToken: (token) => set({ token }),
  setUser: (user) => set({ user }),
  logout: () => set({ isLogin: false, token: null, user: null }),
}));

export default useAuthStore;
