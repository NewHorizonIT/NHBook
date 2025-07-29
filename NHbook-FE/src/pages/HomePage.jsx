import apiInstance from '@/api/axios';
import { ListBook } from '@/components/common/Books';
import { Container } from '@/components/layout/Containner';
import { Hero } from '@/components/layout/Hero';
import Title from '@/components/ui/Title';
import useAuthStore from '@/store/usAuthStore';
import useCartStore from '@/store/useCartStore';
import React, { useEffect } from 'react';

const HomePage = () => {
  const { fetchCart } = useCartStore();
  useEffect(() => {
    fetchCart();
    const checkAuth = async () => {
      try {
        const res = await apiInstance.post('/auth/refresh-token');
        console.log(res);
        useAuthStore.getState().setAccessToken(res.accessToken);
        useAuthStore.getState().setIsAuth(true);
      } catch (err) {
        console.log(err);
        useAuthStore.getState().logout(); // clear token, isAuth = false
      }
    };

    checkAuth();
  }, []);
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
