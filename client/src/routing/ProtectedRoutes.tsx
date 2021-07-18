import React, { useContext } from 'react';

import { Redirect, Route, RouteProps } from 'react-router-dom';
import AuthContext from '../contexts/AuthContext';

interface NewRouteProps extends RouteProps {
    component: React.ComponentType<any>;
}

const ProtectedRoute: React.FC<NewRouteProps> = ({
    component: Component,
    ...rest
}) => {
    const auth = useContext(AuthContext);

    return (
        <>
            {auth.token ? (
                <Route
                    {...rest}
                    render={(props) => <Component {...rest} {...props} />}
                />
            ) : (
                <Redirect to="/login" />
            )}
        </>
    );
};

export default ProtectedRoute;
