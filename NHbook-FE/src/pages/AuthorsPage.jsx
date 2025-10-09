import React from 'react';

const AuthorsPage = () => {
  // Dummy authors
  const authors = [
    { id: 1, name: 'Tác giả A' },
    { id: 2, name: 'Tác giả B' },
    { id: 3, name: 'Tác giả C' },
  ];
  return (
    <div className='container mx-auto py-8'>
      <h1 className='text-2xl font-bold mb-6'>Danh sách tác giả</h1>
      <ul className='space-y-4'>
        {authors.map((author) => (
          <li
            key={author.id}
            className='p-4 bg-gray-100 rounded shadow flex justify-between items-center'
          >
            <span>{author.name}</span>
            <button className='px-3 py-1 bg-blue-500 text-white rounded'>
              Xem sách
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
};
export default AuthorsPage;
