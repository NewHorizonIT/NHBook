import { Container } from '@/components/layout/Containner';
import { CartContent, ThemeToggle } from '@/components/layout/Header';
import Nav from '@/components/layout/Header/Nav';
import Avatar from '@/components/ui/Avatar';
import Button from '@/components/ui/Button';
import { Cart } from '@/components/ui/icons';
import { linksNav } from '@/constants/links';
import useAuthStore from '@/store/usAuthStore';
import { AnimatePresence } from 'motion/react';

import React, { useState } from 'react';
import { NavLink } from 'react-router-dom';

const Header = () => {
  const { isAuth, setIsAuth, setDataUser } = useAuthStore();
  const [openCart, setOpenCart] = useState(false);

  const handleShowCart = () => {
    setOpenCart(true);
  };

  const handleLogout = () => {
    setIsAuth(false);
    setDataUser({});
  };

  return (
    <div className='bg-primary w-full fixed z-50'>
      <Container>
        <div className='flex justify-between items-center w-full h-[65px] bg-primary'>
          {/* Left Header */}
          <p className='text-2xl font-playwrite font-bold text-primary-content'>
            NHBOOK
          </p>
          {/* Center Header */}
          <Nav links={linksNav} />
          {/* Right Header */}
          <div className='flex justify-end items-center gap-4'>
            <ThemeToggle />
            {isAuth ? (
              <>
                <Button
                  style='ghost'
                  wide='square'
                  wides='block'
                  onClick={handleShowCart}
                  className='hover:bg-transparent border-0 shadow-none'
                >
                  <Cart />
                </Button>
                <div className='dropdown dropdown-hover'>
                  <Avatar />
                  <ul
                    tabIndex={0}
                    className='dropdown-content menu bg-base-100 rounded-box z-1 w-52 p-2 shadow-sm'
                  >
                    <li>
                      <NavLink to={'/user/profile'}>Thông tin cá nhân</NavLink>
                    </li>
                    <li onClick={() => handleLogout()}>
                      <p>Đăng xuất</p>
                    </li>
                  </ul>
                </div>
              </>
            ) : (
              <>
                <Button>
                  <NavLink to='/auth/login' className='min-w-[80px]'>
                    Đăng nhập
                  </NavLink>
                </Button>
                <Button>
                  <NavLink to='/auth/register' className='min-w-[80px]'>
                    Đăng kí
                  </NavLink>
                </Button>
              </>
            )}
          </div>
        </div>
      </Container>

      {/* Cart content */}
      <AnimatePresence mode='wait'>
        {openCart && <CartContent onClose={() => setOpenCart(false)} />}
      </AnimatePresence>
    </div>
  );
};

export default Header;
