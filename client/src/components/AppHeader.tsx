import { Container, Flex, Heading } from '@chakra-ui/layout';
import { Button, Center, Image } from '@chakra-ui/react';
import { useContext } from 'react';
import AuthContext from '../contexts/AuthContext';

function AppHeader() {
    const authContext = useContext(AuthContext);

    return (
        <>
            <Flex p="2" background="teal">
                <Container maxW="container.md">
                    <Flex
                        direction={['column', 'row']}
                        justifyContent={'space-between'}
                    >
                        <Heading color="gray.100"> TODOS-APP </Heading>

                        {authContext.token ? <LogOutButton /> : ''}
                    </Flex>
                </Container>
            </Flex>
        </>
    );
}

const LogOutButton = () => {
    const authContext = useContext(AuthContext);
    const Logout = () => authContext.setToken('');
    return (
        <>
            <Flex>
                <Button
                    onClick={Logout}
                    color="gray.100"
                    variant={'link'}
                    paddingRight={'1'}
                >
                    Logout as {authContext.userPrincipal?.name}
                </Button>
                <Center>
                    <Image
                        boxSize={'30px'}
                        // boxSize={}
                        borderRadius="full"
                        src={authContext.userPrincipal?.picture}
                    />
                </Center>
            </Flex>
        </>
    );
};

export default AppHeader;
