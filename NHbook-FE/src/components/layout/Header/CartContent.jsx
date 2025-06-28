import { motion } from 'motion/react';
import Button from '@/components/ui/Button';
import useCartStore from '@/store/useCartStore';
import { Close } from '@/components/ui/icons';
import Table from '@/components/ui/Table';
import { useState, useEffect } from 'react';

const CartContent = ({ onClose }) => {
  const { items, removeItem, updateItem } = useCartStore();
  console.log('ITEMS', items);
  const [quantities, setQuantities] = useState(() =>
    Object.fromEntries(items.map((item) => [item.id, item.quantity]))
  );

  const columns = [
    {
      header: 'Ảnh bìa',
      accessor: 'thumbnail',
      render: (row) => (
        <img
          src={row.thumbnail}
          alt={row.title}
          className='w-16 h-16 object-cover'
        />
      ),
    },
    {
      header: 'Tên sách',
      accessor: 'title',
    },
    {
      header: 'Số lượng',
      accessor: 'quantity',
      render: (row) => (
        <input
          type='number'
          value={quantities[row.id] || row.quantity}
          min={1}
          onChange={(e) => {
            const value = parseInt(e.target.value);
            setQuantities((prev) => ({ ...prev, [row.id]: value }));
            updateItem(row.id, value);
          }}
          className='input input-bordered w-16'
        />
      ),
    },
    {
      header: 'Giá',
      accessor: 'price',
    },
    {
      header: '',
      accessor: '',
      render: (row) => (
        <Button onClick={() => removeItem(row.id)}>
          <Close />
        </Button>
      ),
    },
  ];

  useEffect(() => {
    setQuantities(
      Object.fromEntries(items.map((item) => [item.id, item.quantity]))
    );
  }, [items]);

  return (
    <div className='bg-black/40 w-screen h-screen fixed top-0 left-0 z-50'>
      <motion.div
        className='fixed top-0 right-0 max-w-[300px] md:max-w-[450px] w-full h-screen z-[60] bg-base-100'
        key='cart'
        initial={{ x: '100%' }}
        animate={{ x: 0 }}
        exit={{ x: '100%' }}
        transition={{ type: 'tween', duration: 0.3 }}
      >
        <div className='p-4 border-b border-base-300 flex items-center justify-between'>
          <h2 className='text-xl font-semibold'>Giỏ hàng</h2>
          <Button
            onClick={onClose}
            className={'shadow-none hover:bg-transparent border-0'}
            style={'ghost'}
          >
            <Close />
          </Button>
        </div>
        {items.length === 0 ? (
          <div className='p-4 text-center text-base-content'>
            Giỏ hàng của bạn đang trống!
          </div>
        ) : (
          <Table rows={items} columns={columns} />
        )}
        {/* Footer */}
        <div className='p-4'>
          <Button color={'primary'} wide={'wide'}>
            Mua hàng
          </Button>
        </div>
      </motion.div>
    </div>
  );
};

export default CartContent;
