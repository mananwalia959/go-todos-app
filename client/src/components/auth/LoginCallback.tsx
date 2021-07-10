import { FC, useEffect } from 'react';
import { authService } from '../../services/auth-service';
import { Text, Stack, Square } from '@chakra-ui/react';

type Location = {
    search: string;
};

const LoginCallback: FC<{ location: Location }> = (props) => {
    const queryParams = new URLSearchParams(props.location.search);
    const code: string = queryParams.get('code') || ''; //return empty if not present
    useEffect(() => {
        authService
            .getToken(code)
            .then((res) => {
                console.log(res.jwtToken);
            })
            .catch((err) => {
                console.log(err);
            });
    }, [code]);
    return (
        <>
            <Stack align="center">
                <Square
                    size="xs"
                    border="2px"
                    borderColor="teal"
                    borderRadius="40px"
                    bgColor="gray.100"
                >
                    <Text>Please wait while we are logging you in</Text>
                </Square>
            </Stack>
        </>
    );
};

export default LoginCallback;
