import apiInstance from '@/api/axios';
import CardBook from '@/components/common/Books/CardBook';
import { useQuery } from '@tanstack/react-query';
import React from 'react';
import { Loading } from '../ui';

const ListBook = () => {
  const { data, isLoading, error } = useQuery({
    queryKey: ['books'],
    queryFn: async () => {
      const res = await apiInstance.get('books/', {
        params: {
          limit: 6,
        },
      });
      return res.metadata;
    },
  });

  if (isLoading) {
    return <Loading />;
  }

  if (error || !data.data) {
    return (
      <div className='text-center text-xl font-semibold'>
        Không có cuốn sách nào!
      </div>
    );
  }

  return (
    <div className='grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-5'>
      {data.data.map((book) => (
        <CardBook key={book.id} book={book} />
      ))}
    </div>
  );
};
export default ListBook;
