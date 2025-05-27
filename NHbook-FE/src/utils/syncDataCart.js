import apiInstance from '@/api/axios';
import { debounce } from 'lodash';

const syncDataCart = debounce(async () => {
  try {
    await apiInstance.post('/');
  } catch (error) {
    console.log(error);
  }
});

export default syncDataCart;
