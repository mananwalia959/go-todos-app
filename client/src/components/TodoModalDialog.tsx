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
import React, { FC } from 'react';
import { Todo } from '../models/todos';

const TodoModalDialog: FC<{
    isOpen: boolean;
    onClose: () => void;
    todo?: Todo;
}> = (props) => {
    const { isOpen, onClose } = props;

    return (
        <>
            <Modal isOpen={isOpen} onClose={onClose}>
                <ModalOverlay />
                <ModalContent>
                    <ModalHeader>Todo</ModalHeader>
                    <ModalCloseButton />
                    <ModalBody>
                        <VStack spacing={2}>
                            <Input placeholder="Todo Name" />
                            <Textarea
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
                        <Button colorScheme="teal">Save</Button>
                    </ModalFooter>
                </ModalContent>
            </Modal>
        </>
    );
};

export default TodoModalDialog;
