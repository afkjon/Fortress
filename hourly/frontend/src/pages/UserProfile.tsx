import React from 'react';
import { 
    Container, 
    Box,
    Button,
    Divider
} from '@chakra-ui/react';
import { useAuth } from '../AuthContext';

const UserProfile : React.FC = () => {
    const { user } = useAuth();

    return (
        <Container>
            <h1>User Profile</h1>
            <Box>
                <p>Email: {user?.email}</p>
                <Divider margin="10px" />
                <Button 
                    bg="pink"
                    margin="10px">
                    Delete Account
                </Button>
            </Box>
        </Container>
    )
}

export default UserProfile;