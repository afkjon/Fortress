import React from 'react';
import { Outlet, Navigate } from 'react-router-dom';
import { useAuth } from '../AuthContext';

const PrivateRoute : React.FC = () => {
    const { user } = useAuth()

    return user ? <Outlet /> : <Navigate to="/login" />;
};

export default PrivateRoute;
