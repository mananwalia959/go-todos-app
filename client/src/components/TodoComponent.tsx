import { IconButton } from '@chakra-ui/button';
import { Checkbox } from '@chakra-ui/checkbox';
import { Box, Flex, Heading, Spacer, Text } from '@chakra-ui/layout';
import { FC } from 'react';
import { Todo } from '../models/todos';
import { EditIcon } from './svg/EditIcon';

const TodoComponent: FC<{ todo: Todo }> = (props) => {
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
                            variant="link"
                            colorScheme="teal"
                            aria-label="Edit Todo"
                            size="lg"
                            icon={<EditIcon />}
                        />
                    </Flex>
                </Flex>

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
