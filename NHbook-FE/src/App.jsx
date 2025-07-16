import { useEffect } from 'react';
import useCartStore from './store/useCartStore';

function App() {
  const { fetchCart } = useCartStore();
  useEffect(() => {
    fetchCart();
  });
  return (
    <>
      <h1>Start Project</h1>
    </>
  );
}

export default App;
