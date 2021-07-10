import { Container } from '@chakra-ui/layout';
import { Route, Switch } from 'react-router-dom';
import AppHeader from './components/AppHeader';
import TodosList from './components/TodosList';
import LoginCallback from './components/auth/LoginCallback';
import LoginPage from './components/auth/LoginPage';

function App() {
    return (
        <>
            <AppHeader />
            <Container maxW="container.md" p="2">
                <Switch>
                    <Route path="/" exact component={TodosList}></Route>

                    <Route path="/login" exact component={LoginPage}></Route>

                    <Route
                        path="/callback/googleoauth"
                        component={LoginCallback}
                    ></Route>
                </Switch>
            </Container>
        </>
    );
}

export default App;
