import { create } from "zustand";

const useBookStore = create((set) => ({
  listBook: [],
  setListBook: (books) => set({ listBook: books }),
}));

export default useBookStore;
