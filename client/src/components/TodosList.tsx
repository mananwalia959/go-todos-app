import { Button } from '@chakra-ui/button';
import { Center, Flex, Grid, Spacer, Text } from '@chakra-ui/layout';
import { useEffect, useState } from 'react';
import { Todo } from '../models/todos';
import { todoService } from '../services/todos-service';
import TodoComponent from './TodoComponent';

function TodosList() {
    const [todos, setTodos] = useState([] as Todo[]);

    useEffect(() => {
        todoService
            .getAllTodos()
            .then((t) => setTodos(t))
            .catch((err) => console.log(err));
    }, []);

    return (
        <>
            <Flex p="3">
                <Center>
                    <Text fontSize="lg"> All Todos </Text>
                </Center>
                <Spacer />
                <Button ml="2" colorScheme="teal">
                    Add todo
                </Button>
            </Flex>
            <Grid direction="column" gridAutoRows="1fr" gap="2" mt="2">
                {todos.map((t) => (
                    <TodoComponent key={t.id} todo={t} />
                ))}
            </Grid>
        </>
    );
}

export default TodosList;
