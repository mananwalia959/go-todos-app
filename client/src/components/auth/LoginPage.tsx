import { Button, Stack, Square, Heading } from '@chakra-ui/react';
import { GoogleIcon } from './../svg/GoogleIcon';
import React from 'react';
import { authService } from '../../services/auth-service';

const LoginPage = () => {
    const onLoginClick = () => authService.getOauthUrl();
    return (
        <Stack align="center">
            <Heading as="h4" size="md" color="teal">
                Please Login To Access your todos
            </Heading>
            <Square
                size={['15em', 'sm']}
                border="2px"
                borderColor="teal"
                borderRadius="40px"
                bgColor="gray.100"
            >
                <Button
                    onClick={onLoginClick}
                    leftIcon={<GoogleIcon />}
                    colorScheme="teal"
                    variant="outline"
                >
                    Login with Google
                </Button>
            </Square>
        </Stack>
    );
};

export default LoginPage;
