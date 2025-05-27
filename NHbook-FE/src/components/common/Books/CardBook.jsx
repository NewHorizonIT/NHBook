import Button from '@/components/ui/Button';
import Image from '@/components/ui/Image';
import useCartStore from '@/store/useCartStore';
import formatPrice from '@/utils/formatPrice';
import React from 'react';
import { NavLink } from 'react-router-dom';

const CardBook = ({ book }) => {
  const { id, title, thumbnail, price, authors } = book;
  const { addItem } = useCartStore();

  const handleAddTocart = () => {
    const item = { id, title, thumbnail, quantity: 1, price };
    addItem(item);
  };

  return (
    <div className='card bg-base-200 shadow-sm p-5'>
      <figure>
        <Image
          src={thumbnail}
          alt='anh sach'
          fallback='https://placehold.co/600x400'
          className='max-h-[200px] h-[200px]'
        />
      </figure>
      <div className='card-body'>
        <h2 className='card-title'>
          {title}
          <div className='badge badge-secondary'>NEW</div>
        </h2>
        <div>
          {authors.map((a, idx) => (
            <span key={idx}>{a.name}</span>
          ))}
        </div>
        <p className='text-primary text-2xl'>{formatPrice(price)}</p>

        <div className='card-actions flex flex-col gap-4 items-center'>
          <Button wide='wide' color='primary' size='md'>
            <NavLink to={`/books/${id}`}>Xem chi tiết</NavLink>
          </Button>
          <Button
            wide='wide'
            color='primary'
            size='md'
            onClick={handleAddTocart}
          >
            Thêm vào giỏ hàng
          </Button>
        </div>
      </div>
    </div>
  );
};

export default CardBook;
