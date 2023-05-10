import { Box, Typography, useTheme } from "@mui/material";
import Friend from "../../components/Friend";
import WidgetWrapper from "../../components/WidgetWrapper";
import { useDispatch, useSelector } from "react-redux";
import { setFriends, selectCurrentUser } from "../../features/auth/authSlice";
import { useGetUserQuery} from '../../features/users/usersApiSlice'

const FriendListWidget = ({ userID }) => {
  const dispatch = useDispatch();
  const { palette } = useTheme();
  const {friends} =  useSelector(selectCurrentUser)
  const {
  data:user,
  isSuccess,
  isError,  
} = useGetUserQuery(userID)

 if (isError && !user){
    return null;
  }

if (isSuccess){  
  dispatch(setFriends(user ));
}

  return (
    <WidgetWrapper>
      <Typography
        color={palette.neutral.dark}
        variant="h5"
        fontWeight="500"
        sx={{ mb: "1.5rem" }}
      >
        Friend List
      </Typography>
      <Box display="flex" flexDirection="column" gap="1.5rem">
        {friends && ( friends.map((friend) => (
          <Friend
            key={friend.id}
            friendId={friend.id}
            name={`${friend.firstName} ${friend.lastName}`}
            subtitle={friend.occupation}
            userPicturePath={friend.picturePath}
          />
        )) )}
      </Box>
    </WidgetWrapper>
  );
};

export default FriendListWidget;