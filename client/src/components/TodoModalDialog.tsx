import { Button } from '@chakra-ui/button';
import { Input } from '@chakra-ui/input';
import { VStack } from '@chakra-ui/layout';
import {
    Modal,
    ModalBody,
    ModalCloseButton,
    ModalContent,
    ModalFooter,
    ModalHeader,
    ModalOverlay,
} from '@chakra-ui/modal';
import { Textarea } from '@chakra-ui/textarea';

import React, { FC, useEffect, useState } from 'react';
import { Todo } from '../models/Todo';
import { TodoCreateRequest } from '../models/TodoCreateRequest';
import { TodoEditRequest } from '../models/TodoEditRequest';
import { todoService } from '../services/todos-service';

const TodoModalDialog: FC<{
    isOpen: boolean;
    onClose: () => void;
    todo?: Todo;
    onSave: (todo: Todo) => void;
}> = (props) => {
    const { isOpen, onClose } = props;
    const { onSave } = props;
    const [isClicked, setClicked] = useState(false);
    const [todoName, setTodoName] = useState('');
    const [todoDescription, setTodoDescription] = useState('');
    const isNewTodo = props.todo ? false : true;
    const todo: Todo = props.todo ? props.todo : ({} as Todo);

    useEffect(() => {
        if (isNewTodo) {
            setTodoName('');
            setTodoDescription('');
        } else {
            setTodoName(todo.name);
            setTodoDescription(todo.description);
        }
        setClicked(false);
    }, [isOpen, isNewTodo, todo.description, todo.name]);

    const onSaveButton = async () => {
        setClicked(true);
        if (todoName.trim() === '') return;

        if (isNewTodo) {
            const newRequest: TodoCreateRequest = {
                name: todoName,
                description: todoDescription,
            };
            const newTodo = await todoService.saveNewTodo(newRequest);
            onSave(newTodo);
        } else {
            const newRequest: TodoEditRequest = {
                name: todoName,
                description: todoDescription,
                completed: todo.completed,
            };

            const newTodo = await todoService.editTodo(todo.id, newRequest);
            onSave(newTodo);
        }
        onClose();
    };

    return (
        <>
            <Modal isOpen={isOpen} onClose={onClose}>
                <ModalOverlay />
                <ModalContent>
                    <ModalHeader>Todo</ModalHeader>
                    <ModalCloseButton />
                    <ModalBody>
                        <VStack spacing={2}>
                            <Input
                                isInvalid={isClicked && todoName.trim() === ''}
                                value={todoName}
                                onChange={(event) => {
                                    setTodoName(event.target.value);
                                }}
                                placeholder="Todo Name"
                            />
                            <Textarea
                                value={todoDescription}
                                onChange={(event) => {
                                    setTodoDescription(event.target.value);
                                }}
                                resize="none"
                                placeholder="Describe your todos here"
                            />
                        </VStack>
                    </ModalBody>

                    <ModalFooter>
                        <Button
                            variant="outline"
                            colorScheme="teal"
                            mr={3}
                            onClick={onClose}
                        >
                            Close
                        </Button>
                        <Button onClick={onSaveButton} colorScheme="teal">
                            Save
                        </Button>
                    </ModalFooter>
                </ModalContent>
            </Modal>
        </>
    );
};

export default TodoModalDialog;
