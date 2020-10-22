 
import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreateBill from './components/RecordBill';
import ShowBill from './components/RecordBillTable';
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/RecordBill', CreateBill);
    router.registerRoute('/RecordBillTable', ShowBill);
  },
});
