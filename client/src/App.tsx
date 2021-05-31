import { Container } from '@chakra-ui/layout';
import AppHeader from './components/AppHeader';
import TodosList from './components/TodosList';

function App() {
    return (
        <>
            <AppHeader />
            <Container maxW="container.md">
                <TodosList></TodosList>
            </Container>
        </>
    );
}

export default App;
