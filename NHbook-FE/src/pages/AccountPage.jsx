import React from 'react';

const AccountPage = () => {
    // Dummy user info
    const user = {
        name: 'Nguyễn Văn A',
        email: 'nguyenvana@example.com',
        avatar: '/src/assets/logo.png',
    };
    return (
        <div className='container mx-auto py-8'>
            <h1 className='text-2xl font-bold mb-6'>Thông tin tài khoản</h1>
            <div className='flex items-center gap-6'>
                <img
                    src={user.avatar}
                    alt='avatar'
                    className='w-24 h-24 rounded-full border'
                />
                <div>
                    <div className='font-semibold'>{user.name}</div>
                    <div className='text-gray-600'>{user.email}</div>
                </div>
            </div>
            <div className='mt-8'>
                <button className='px-4 py-2 bg-blue-500 text-white rounded'>
                    Cập nhật thông tin
                </button>
            </div>
        </div>
    );
};
export default AccountPage;
