import { FC, useContext, useEffect, useState } from 'react';
import { authService } from '../../services/auth-service';
import { Text, Stack, Square } from '@chakra-ui/react';
import AuthContext from '../../contexts/AuthContext';

type Location = {
    search: string;
};

const LoginCallback: FC<{ location: Location }> = (props) => {
    const queryParams = new URLSearchParams(props.location.search);
    const code: string = queryParams.get('code') || ''; //return empty if not present
    const authContext = useContext(AuthContext);
    // temporary workaround
    const [renderedOnce, setRenderOnce] = useState(false);

    useEffect(() => {
        if (!renderedOnce) {
            authService
                .getToken(code)
                .then((res) => {
                    setRenderOnce(true);
                    const token = res.jwtToken;
                    console.log(token);
                    authContext.setToken(token);
                })
                .catch((err) => {
                    setRenderOnce(true);
                    console.log(err);
                });
        }
    }, [code, renderedOnce, authContext]);
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
