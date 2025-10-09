import React from 'react';

const PurchaseForm = ({
  cartItems,
  address,
  paymentMethod,
  onConfirm,
  onClose,
}) => {
  const totalPrice = cartItems.reduce(
    (total, item) => total + item.price * item.quantity,
    0
  );
  return (
    <div>
      <h2 className='text-2xl font-semibold mb-4 text-center'>
        Xác nhận đơn hàng
      </h2>

      {/* Danh sách sách */}
      <div className='mb-4 max-h-60 overflow-y-auto'>
        {cartItems.map((item, index) => (
          <div
            key={index}
            className='flex justify-between items-center py-2 border-b text-sm'
          >
            <div className='font-medium'>{item.title}</div>
            <div className='text-gray-600'>x{item.quantity}</div>
            <div className='text-right text-gray-800'>
              {item.price.toLocaleString()}₫
            </div>
          </div>
        ))}
      </div>

      {/* Tổng tiền */}
      <div className='flex justify-between font-semibold mb-4'>
        <span>Tổng cộng:</span>
        <span className='text-red-600'>{totalPrice.toLocaleString()}₫</span>
      </div>

      {/* Địa chỉ */}
      <div className='mb-2'>
        <h4 className='font-medium'>Địa chỉ nhận hàng:</h4>
        <p className='text-sm text-gray-700'>{address}</p>
      </div>

      {/* Phương thức thanh toán */}
      <div className='mb-6'>
        <h4 className='font-medium'>Phương thức thanh toán:</h4>
        <p className='text-sm text-gray-700'>{paymentMethod}</p>
      </div>

      {/* Hành động */}
      <div className='flex justify-end gap-3'>
        <button
          onClick={onClose}
          className='px-4 py-2 border rounded-md hover:bg-gray-100'
        >
          Hủy
        </button>
        <button
          onClick={onConfirm}
          className='px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700'
        >
          Xác nhận mua
        </button>
      </div>
    </div>
  );
};

export default PurchaseForm;
