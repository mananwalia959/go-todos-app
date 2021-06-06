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

import React, { FC, useState } from 'react';
import { Todo } from '../models/Todo';
import { TodoCreateRequest } from '../models/TodoCreateRequest';
import { todoService } from '../services/todos-service';

const TodoModalDialog: FC<{
    isOpen: boolean;
    onClose: () => void;
    todo?: Todo;
    onSave?: (todo: Todo) => void;
    isNewTodo?: boolean;
}> = (props) => {
    const { isOpen, onClose } = props;
    const { isNewTodo, onSave } = props;
    const [isClicked, setClicked] = useState(false);
    const [todoName, setTodoName] = useState('');
    const [todoDescription, setTodoDescription] = useState('');

    const onCloseDialog = () => {
        setTodoName('');
        setTodoDescription('');
        setClicked(false);
        onClose();
    };
    const onSaveButton = async () => {
        setClicked(true);
        if (todoName.trim() === '') return;

        if (isNewTodo) {
            const newRequest: TodoCreateRequest = {
                name: todoName,
                description: todoDescription,
            };
            const newTodo = await todoService.saveNewTodo(newRequest);
            onSave !== undefined && onSave(newTodo);
        }
        onCloseDialog();
    };

    return (
        <>
            <Modal isOpen={isOpen} onClose={onCloseDialog}>
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
                            onClick={onCloseDialog}
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
