import { IconButton } from '@chakra-ui/button';
import { Checkbox } from '@chakra-ui/checkbox';
import { useDisclosure } from '@chakra-ui/hooks';
import { Box, Flex, Heading, Spacer, Text } from '@chakra-ui/layout';
import { FC } from 'react';
import { Todo } from '../models/Todo';
import TodoModalDialog from './TodoModalDialog';
import { EditIcon } from './svg/EditIcon';

const TodoComponent: FC<{ todo: Todo }> = (props) => {
    const { isOpen, onOpen, onClose } = useDisclosure();
    const todo: Todo = props.todo;
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
                            mx="2"
                            borderColor="teal"
                        />

                        <IconButton
                            onClick={onOpen}
                            variant="link"
                            colorScheme="teal"
                            aria-label="Edit Todo"
                            size="lg"
                            icon={<EditIcon />}
                        />
                    </Flex>
                </Flex>

                <TodoModalDialog isOpen={isOpen} onClose={onClose} />

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
