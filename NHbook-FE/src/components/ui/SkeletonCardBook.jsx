import React from 'react';

const SkeletonCardBook = () => {
  return (
    <div className='animate-pulse flex flex-col bg-white rounded-xl shadow p-4'>
      <div className='bg-gray-200 h-48 w-full rounded-md mb-4'></div>
      <div className='h-4 bg-gray-200 rounded w-3/4 mb-2'></div>
      <div className='h-3 bg-gray-200 rounded w-1/2 mb-4'></div>
      <div className='h-8 bg-gray-200 rounded w-full'></div>
    </div>
  );
};

export default SkeletonCardBook;
