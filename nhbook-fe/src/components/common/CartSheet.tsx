import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "../ui/sheet";
import { Button } from "../ui/button";
import { demoBooks } from "@/mock/demo_book";
import Image from "next/image";
import { useMemo, useState } from "react";

type CartItem = {
  id: number;
  qty: number;
};

interface CartSheetProps {
  children: React.ReactNode;
  cart: CartItem[];
}

const CartSheet = ({ children, cart = [] }: CartSheetProps) => {
  type Book = (typeof demoBooks)[number] & { qty: number };
  const cartBooks: Book[] = useMemo(
    () =>
      cart
        .map((item) => {
          const book = demoBooks.find((b) => b.id === item.id);
          return book ? { ...book, qty: item.qty } : null;
        })
        .filter(Boolean) as Book[],
    [cart]
  );

  const total = cartBooks.reduce(
    (sum, b) => sum + (b.salePrice || b.price) * b.qty,
    0
  );

  const handleRemove = (id: number) => {
    // TODO: Xử lý xóa sản phẩm khỏi giỏ (gọi props hoặc context)
    alert(`Xóa sản phẩm id ${id}`);
  };

  const [showOrderModal, setShowOrderModal] = useState(false);
  const [address, setAddress] = useState("");
  const [payment, setPayment] = useState("cod");
  const [note, setNote] = useState("");
  const [orderSuccess, setOrderSuccess] = useState(false);

  const handleOrder = () => {
    setShowOrderModal(true);
  };

  const handleConfirmOrder = (e: React.FormEvent) => {
    e.preventDefault();
    setOrderSuccess(true);
    setTimeout(() => {
      setShowOrderModal(false);
      setOrderSuccess(false);
      setAddress("");
      setNote("");
    }, 1500);
  };

  return (
    <Sheet>
      <SheetTrigger asChild>{children}</SheetTrigger>
      <SheetContent>
        <SheetHeader>
          <SheetTitle>Giỏ hàng</SheetTitle>
        </SheetHeader>
        <div className="flex flex-col gap-4 py-4 px-2 max-h-[60vh] overflow-y-auto">
          {cartBooks.length === 0 ? (
            <div className="text-gray-500 text-center py-8">
              Giỏ hàng của bạn đang trống.
            </div>
          ) : (
            cartBooks.map((book) => (
              <div
                key={book.id}
                className="flex items-center gap-3 border-b pb-3 last:border-b-0"
              >
                <Image
                  src={book.image}
                  alt={book.title}
                  width={56}
                  height={80}
                  className="w-14 h-20 object-cover rounded-base border"
                />
                <Image
                  src={book.image}
                  alt={book.title}
                  width={56}
                  height={80}
                  className="w-14 h-20 object-cover rounded-base border"
                />
                <div className="flex-1">
                  <div className="font-medium text-main line-clamp-2">
                    {book.title}
                  </div>
                  <div className="text-sm text-gray-500">SL: {book.qty}</div>
                  <div className="text-sm text-gray-700 font-bold">
                    {(book.salePrice || book.price).toLocaleString()}₫
                  </div>
                </div>
                <Button
                  variant="neutral"
                  className="h-8 px-3 text-xs"
                  onClick={() => handleRemove(book.id)}
                >
                  Xóa
                </Button>
              </div>
            ))
          )}
        </div>
        {/* Tổng tiền và đặt hàng */}
        <div className="border-t pt-4 mt-4 flex flex-col gap-3">
          <div className="flex items-center justify-between font-bold text-main text-lg">
            <span>Tổng tiền:</span>
            <span>{total.toLocaleString()}₫</span>
          </div>
          <Button
            className="w-full"
            disabled={cartBooks.length === 0}
            onClick={handleOrder}
          >
            Đặt hàng
          </Button>
        </div>
        {/* Modal xác nhận đặt hàng */}
        {showOrderModal && (
          <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/40">
            <form
              onSubmit={handleConfirmOrder}
              className="bg-white rounded-base shadow-lg p-8 min-w-[320px] max-w-md w-full relative flex flex-col gap-4"
            >
              <button
                type="button"
                className="absolute top-2 right-2 text-gray-400 hover:text-main text-xl"
                onClick={() => setShowOrderModal(false)}
                aria-label="Đóng"
              >
                ×
              </button>
              <h2 className="text-xl font-bold text-main mb-2">
                Xác nhận đặt hàng
              </h2>
              <div>
                <label className="block font-medium mb-1 text-main">
                  Địa chỉ nhận hàng
                </label>
                <input
                  required
                  className="w-full px-3 py-2 border rounded-base focus:ring-2 focus:ring-main outline-none"
                  placeholder="Nhập địa chỉ nhận hàng"
                  value={address}
                  onChange={(e) => setAddress(e.target.value)}
                />
              </div>
              <div>
                <label className="block font-medium mb-1 text-main">
                  Phương thức thanh toán
                </label>
                <select
                  className="w-full px-3 py-2 border rounded-base focus:ring-2 focus:ring-main outline-none"
                  value={payment}
                  onChange={(e) => setPayment(e.target.value)}
                >
                  <option value="cod">Thanh toán khi nhận hàng (COD)</option>
                  <option value="bank">Chuyển khoản ngân hàng</option>
                  <option value="momo">Ví MoMo</option>
                </select>
              </div>
              <div>
                <label className="block font-medium mb-1 text-main">
                  Ghi chú (tuỳ chọn)
                </label>
                <textarea
                  className="w-full px-3 py-2 border rounded-base focus:ring-2 focus:ring-main outline-none"
                  placeholder="Ghi chú cho đơn hàng..."
                  value={note}
                  onChange={(e) => setNote(e.target.value)}
                  rows={2}
                />
              </div>
              <div className="flex items-center justify-between font-bold text-main text-lg">
                <span>Tổng tiền:</span>
                <span>{total.toLocaleString()}₫</span>
              </div>
              <Button
                type="submit"
                className="w-full mt-2"
                disabled={orderSuccess}
              >
                {orderSuccess ? "Đang xử lý..." : "Xác nhận đặt hàng"}
              </Button>
              {orderSuccess && (
                <div className="text-green-600 text-center font-semibold mt-2">
                  Đặt hàng thành công!
                </div>
              )}
            </form>
          </div>
        )}
        <SheetFooter className="mt-4">
          <SheetClose asChild>
            <Button variant="neutral" className="w-full">
              Đóng
            </Button>
          </SheetClose>
        </SheetFooter>
      </SheetContent>
    </Sheet>
  );
};

export default CartSheet;
