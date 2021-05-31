import { Button } from '@chakra-ui/button';
import { Center, Flex, Spacer, Text } from '@chakra-ui/layout';
import { useEffect, useState } from 'react';
import { Todo } from '../models/todos';
import { todoService } from '../services/todos-service';

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

            {todos.map((todo) => (
                <p id={todo.id}> {todo.name} </p>
            ))}
        </>
    );
}

export default TodosList;
