import React from 'react';

const CategoriesPage = () => {
  // Dummy categories
  const categories = [
    { id: 1, name: 'Khoa học' },
    { id: 2, name: 'Văn học' },
    { id: 3, name: 'Thiếu nhi' },
  ];
  return (
    <div className='container mx-auto py-8'>
      <h1 className='text-2xl font-bold mb-6'>Thể loại sách</h1>
      <ul className='space-y-4'>
        {categories.map((cat) => (
          <li
            key={cat.id}
            className='p-4 bg-gray-100 rounded shadow flex justify-between items-center'
          >
            <span>{cat.name}</span>
            <button className='px-3 py-1 bg-blue-500 text-white rounded'>
              Xem sách
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
};
export default CategoriesPage;
