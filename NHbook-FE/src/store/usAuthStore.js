import { create } from 'zustand';

const useAuthStore = create((set) => ({
  isAuth: false,
  dataUser: null,
  accessToken: null,
  setIsAuth: (status) => set({ isAuth: status }),
  setAccessToken: (accessToken) => set({ accessToken: accessToken }),
  setDataUser: (data) => set({ dataUser: data }),
  logout: () => set({ isAuth: false, accessToken: null, dataUser: null }),
}));

export default useAuthStore;
