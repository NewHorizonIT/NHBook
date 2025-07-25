import { create } from 'zustand';
import { persist } from 'zustand/middleware';

const useAuthStore = create(
  persist(
    (set) => ({
      isAuth: false,
      dataUser: {},
      setIsAuth: (status) => set({ isAuth: status }),
      setDataUser: (data) => set({ dataUser: data }),
    }),
    {
      name: 'auth',
    }
  )
);

export default useAuthStore;
