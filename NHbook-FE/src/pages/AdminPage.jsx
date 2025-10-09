import React from 'react';

const AdminPage = () => {
  return (
    <div className='container mx-auto py-8'>
      <h1 className='text-2xl font-bold mb-6'>Trang quản trị</h1>
      <div className='bg-white p-6 rounded shadow'>
        <p>Chào mừng bạn đến với trang quản trị hệ thống NHBook.</p>
        <ul className='mt-4 list-disc list-inside'>
          <li>Quản lý sách</li>
          <li>Quản lý đơn hàng</li>
          <li>Quản lý người dùng</li>
          <li>Quản lý thể loại</li>
        </ul>
      </div>
    </div>
  );
};
export default AdminPage;
