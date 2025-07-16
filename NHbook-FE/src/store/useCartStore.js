import apiInstance from '@/api/axios';
import _ from 'lodash';
import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import useAuthStore from './usAuthStore';

const updateCartOnServer = _.debounce(async (cart) => {
  try {
    await apiInstance.post('/carts/', cart, {
      headers: {
        Authorization: `Bearer ${useAuthStore.getState().dataUser.accessToken}`,
      },
    });
  } catch (err) {
    console.error('Update cart failed:', err);
  }
}, 1000); // debounce sau 1 giây

const useCartStore = create(
  persist(
    (set, get) => ({
      items: [],
      addItem: (item) => {
        if (get().items.some((i) => i.id === item.id)) {
          get().updateItem(item.id, item.quantity + 1);
          return;
        }
        const newItems = [...get().items, item];
        set({ items: newItems });
        for (const item of get().items) {
          updateCartOnServer(item);
        }
      },
      clearCart: () => {
        set({ items: [] });
      },
      removeItem: (id) => {
        const newItems = get().items.filter((item) => item.id !== id);
        set({ items: newItems });
        for (const item of get().items) {
          updateCartOnServer(item);
        }
      },
      updateItem: (id, quantity) => {
        const newItem = get().items.map((item) =>
          item.id === id ? { ...item, quantity } : item
        );
        set({ items: newItem });
        for (const item of get().items) {
          updateCartOnServer(item);
        }
      },
      fetchCart: async () => {
        try {
          const res = await apiInstance.get('/carts/', {
            headers: {
              Authorization: `Bearer ${
                useAuthStore.getState().dataUser.accessToken
              }`,
            },
          });
          console.log('RES:: ', res);
          if (res.status === 200) {
            set({ items: res.metadata });
          } else {
            set({ items: [] });
          }
        } catch (error) {
          console.log(error);
        }
      },
    }),
    {
      name: 'cart-storage',
    }
  )
);

export default useCartStore;
