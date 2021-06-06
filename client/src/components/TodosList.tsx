import { Button } from '@chakra-ui/button';
import { useDisclosure } from '@chakra-ui/hooks';
import { Center, Flex, Grid, Spacer, Text } from '@chakra-ui/layout';
import { useEffect, useState } from 'react';
import { Todo } from '../models/Todo';
import { todoService } from '../services/todos-service';
import TodoModalDialog from './TodoModalDialog';
import TodoComponent from './TodoComponent';

function TodosList() {
    const [todos, setTodos] = useState([] as Todo[]);
    const { isOpen, onOpen, onClose } = useDisclosure();

    const onNewTodo = (todo: Todo) => {
        const newTodosList = [todo, ...todos];
        setTodos(newTodosList);
    };

    useEffect(() => {
        todoService
            .getAllTodos()
            .then((t) => setTodos(t))
            .catch((err) => console.log(err));
    }, []);

    return (
        <>
            <TodoModalDialog
                isOpen={isOpen}
                onClose={onClose}
                onSave={onNewTodo}
                isNewTodo={true}
            />

            <Flex p="3">
                <Center>
                    <Text fontSize="lg"> All Todos </Text>
                </Center>
                <Spacer />
                <Button onClick={onOpen} ml="2" colorScheme="teal">
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
