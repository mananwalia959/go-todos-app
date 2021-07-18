import { Container } from '@chakra-ui/layout';
import { Route, Switch } from 'react-router-dom';
import AppHeader from './components/AppHeader';
import TodosList from './components/TodosList';
import LoginCallback from './components/auth/LoginCallback';
import LoginPage from './components/auth/LoginPage';
import AuthContext from './contexts/AuthContext';
import { useState } from 'react';
import { authService } from './services/auth-service';
import StorageService from './services/StorageService';
import ProtectedRoute from './routing/ProtectedRoutes';

function App() {
    const [token, setToken] = useState(StorageService.getToken() || '');
    const [userPrincipal, setUserPrincipal] = useState(
        authService.tokenToUserPrincipal(token)
    );

    const setTokenAndUserPrincipal = (tkn: string) => {
        StorageService.setToken(tkn);
        setToken(tkn);
        setUserPrincipal(authService.tokenToUserPrincipal(tkn));
    };

    return (
        <>
            <AuthContext.Provider
                value={{
                    token,
                    userPrincipal,
                    setToken: setTokenAndUserPrincipal,
                }}
            >
                <AppHeader />
                <Container maxW="container.md" p="2">
                    <Switch>
                        <ProtectedRoute path="/" exact component={TodosList} />
                        <Route path="/login" exact component={LoginPage} />

                        <Route
                            path="/callback/googleoauth"
                            component={LoginCallback}
                        />
                    </Switch>
                </Container>
            </AuthContext.Provider>
        </>
    );
}

export default App;
