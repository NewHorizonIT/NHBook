import { ListBook } from '@/components/common/Books';
import { Container } from '@/components/layout/Containner';
import { Hero } from '@/components/layout/Hero';
import Title from '@/components/ui/Title';
import useCartStore from '@/store/useCartStore';
import React, { useEffect } from 'react';

const HomePage = () => {
  const { fetchCart } = useCartStore();
  useEffect(() => {
    fetchCart();
  }, [fetchCart]);
  return (
    <Container>
      <Hero />
      <div className='my-5 flex flex-col gap-4'>
        <Title size='xl'>Sách bán chạy</Title>
        <ListBook />
      </div>
    </Container>
  );
};

export default HomePage;
