import { create } from 'zustand';
import { persist } from 'zustand/middleware';

const useCartStore = create(
  persist(
    (set, get) => ({
      items: [],
      addItem: (item) => {
        const newItems = [...get().items, item];
        set({ items: newItems });
        // Using debounce func
      },
      clearCart: () => {
        set({ items: [] });
        // Using debounce func
      },
      removeItem: (id) => {
        const newItems = get().items.filter((item) => item.id !== id);
        set({ items: newItems });
        // Using debounce func
      },
      updateItem: (id, quantity) => {
        const newItem = get().items.map((item) =>
          item.id === id ? { ...item, quantity } : item
        );
        set({ items: newItem });
        // Using debounce func
      },
    }),
    {
      name: 'cart-storage',
    }
  )
);

export default useCartStore;
