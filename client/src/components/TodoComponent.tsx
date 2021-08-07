import { IconButton } from '@chakra-ui/button';
import { Checkbox } from '@chakra-ui/checkbox';
import { useDisclosure } from '@chakra-ui/hooks';
import { Box, Flex, Heading, Spacer, Text } from '@chakra-ui/layout';
import { FC, useState } from 'react';
import { Todo } from '../models/Todo';
import TodoModalDialog from './TodoModalDialog';
import { EditIcon } from './svg/EditIcon';
import { todoService } from '../services/todos-service';
import { TodoEditRequest } from '../models/TodoEditRequest';
import { DeleteIcon } from './svg/DeleteIcon';

const TodoComponent: FC<{
    todo: Todo;
    deleteTodo: (todoId: string) => void;
}> = (props) => {
    const { isOpen, onOpen, onClose } = useDisclosure();
    const [todo, setTodo] = useState(props.todo);
    const editTodo = (newTodo: Todo) => {
        setTodo({ ...newTodo });
    };

    const deleteTodo = props.deleteTodo;

    const onDelete = async () => {
        const todoId = todo.id;
        await todoService.deleteTodo(todoId);
        deleteTodo(todoId);
    };

    const onCheckbox = async (isChecked: boolean) => {
        todo.completed = isChecked;

        // do this a bit early in case of latency so checkbox is checked before response
        // setchecked(todo.completed);
        setTodo({ ...todo, completed: isChecked });

        const newRequest: TodoEditRequest = {
            name: todo.name,
            description: todo.description,
            completed: isChecked,
        };

        const newTodo = await todoService.editTodo(todo.id, newRequest);
        setTodo({ ...newTodo });
    };

    return (
        <>
            <Box
                borderWidth="1px"
                borderRadius="lg"
                overflow="hidden"
                borderColor="teal"
                p="3"
            >
                <Flex>
                    <Heading isTruncated size="sm">
                        {todo.name}
                    </Heading>
                    <Spacer />
                    <Flex>
                        <Checkbox
                            colorScheme="teal"
                            size="lg"
                            mx="4"
                            borderColor="teal"
                            isChecked={todo.completed}
                            onChange={(e) => onCheckbox(e.target.checked)}
                        />

                        <IconButton
                            onClick={onOpen}
                            variant="link"
                            colorScheme="teal"
                            aria-label="Edit Todo"
                            size="lg"
                            icon={<EditIcon />}
                        />

                        <IconButton
                            onClick={onDelete}
                            variant="link"
                            colorScheme="teal"
                            aria-label="Delete Todo"
                            size="lg"
                            icon={<DeleteIcon />}
                        />
                    </Flex>
                </Flex>

                <TodoModalDialog
                    isOpen={isOpen}
                    onClose={onClose}
                    todo={todo}
                    onSave={editTodo}
                />

                {todo.description ? (
                    <Text isTruncated>{todo.description} </Text>
                ) : (
                    <Text fontStyle="italic" color="gray.600">
                        No description for this todo
                    </Text>
                )}
            </Box>
        </>
    );
};

export default TodoComponent;
