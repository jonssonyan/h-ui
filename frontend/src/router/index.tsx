import { createBrowserRouter } from 'react-router-dom';
import Login from '../pages/login';
import Error404 from '../pages/404';
import Dashboard from '../pages/dashboard';
import Config from '../pages/config';
import AdminLayout from '../layout';
import Account from '../pages/account';
import Hysteria2 from '../pages/hysteria2';

export const router = createBrowserRouter([
    {
        path: '/',
        element: <Login />,
    },

    {
        path: '/dashboard',
        element: <AdminLayout />,
        children: [
            {
                path: 'index',
                element: <Dashboard />,
            },
        ],
    },
    {
        path: '/account',
        element: <AdminLayout />,
        children: [
            {
                path: 'index',
                element: <Account />,
            },
        ],
    },
    {
        path: '/config',
        element: <AdminLayout />,
        children: [
            {
                path: 'index',
                element: <Config />,
            },
        ],
    },
    {
        path: '/hysteria2',
        element: <AdminLayout />,
        children: [
            {
                path: 'index',
                element: <Hysteria2 />,
            },
        ],
    },

    {
        path: '*',
        element: <Error404 />,
    },
]);

export default router;