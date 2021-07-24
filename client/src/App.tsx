import { Container } from '@chakra-ui/layout';
import { Route, Switch } from 'react-router-dom';
import AppHeader from './components/AppHeader';
import TodosList from './components/TodosList';
import LoginCallback from './components/auth/LoginCallback';
import LoginPage from './components/auth/LoginPage';
import AuthContext from './contexts/AuthContext';
import { useEffect, useState } from 'react';
import { authService } from './services/auth-service';
import StorageService from './services/StorageService';
import ProtectedRoute from './routing/ProtectedRoutes';
import { UserPrincipal } from './models/auth/UserPrincipal';

function App() {
    const [token, setToken] = useState('');
    const [userPrincipal, setUserPrincipal] = useState({} as UserPrincipal);
    const [isLoaded, setLoaded] = useState(false);

    const setTokenAndUserPrincipal = (tkn: string) => {
        StorageService.setToken(tkn);
        setToken(tkn);
        setUserPrincipal(authService.tokenToUserPrincipal(tkn));
    };

    useEffect(() => {
        const tkn = StorageService.getToken() || '';
        // no need to make calls if there is no token
        if (!tkn) {
            setLoaded(true);
            return;
        }
        console.log('here');

        authService
            .isTokenValidTillNextDay(tkn)
            .then(() => {
                setToken(tkn);
                setUserPrincipal(authService.tokenToUserPrincipal(tkn));
            })
            .finally(() => setLoaded(true));
    }, []);

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

                {isLoaded ? <Routes /> : ''}
            </AuthContext.Provider>
        </>
    );
}

const Routes = () => (
    <Container maxW="container.md" p="2">
        <Switch>
            <ProtectedRoute path="/" exact component={TodosList} />
            <Route path="/login" exact component={LoginPage} />

            <Route path="/callback/googleoauth" component={LoginCallback} />
        </Switch>
    </Container>
);

export default App;
