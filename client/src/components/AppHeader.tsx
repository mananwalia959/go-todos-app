import { Container, Flex, Heading } from '@chakra-ui/layout';

function AppHeader() {
    return (
        <>
            <Flex p = "2" background = "teal">
                <Container maxW="container.md">
                    <Heading color="gray.100"> TODOS-APP </Heading>
                </Container>
            </Flex>
        </>
    );
}

export default AppHeader;
