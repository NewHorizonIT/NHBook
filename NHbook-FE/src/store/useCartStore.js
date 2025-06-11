import apiInstance from '@/api/axios';
import _ from 'lodash';
import { create } from 'zustand';
import { persist } from 'zustand/middleware';

const updateCartOnServer = _.debounce(async (cart) => {
  try {
    await apiInstance.post('/cart/update', cart);
  } catch (err) {
    console.error('Update cart failed:', err);
  }
}, 1000); // debounce sau 1 giây

const useCartStore = create(
  persist(
    (set, get) => ({
      items: [],
      addItem: (item) => {
        const newItems = [...get().items, item];
        set({ items: newItems });
        updateCartOnServer(newItems);
      },
      clearCart: () => {
        set({ items: [] });
        updateCartOnServer([]);
      },
      removeItem: (id) => {
        const newItems = get().items.filter((item) => item.id !== id);
        set({ items: newItems });
        updateCartOnServer(newItems);
      },
      updateItem: (id, quantity) => {
        const newItem = get().items.map((item) =>
          item.id === id ? { ...item, quantity } : item
        );
        set({ items: newItem });
        updateCartOnServer(newItem);
      },
    }),
    {
      name: 'cart-storage',
    }
  )
);

export default useCartStore;
