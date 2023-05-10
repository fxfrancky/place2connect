import { PersonAddOutlined, PersonRemoveOutlined } from "@mui/icons-material";
import { Box, IconButton, Typography, useTheme } from "@mui/material";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";
import FlexBetween from "./FlexBetween";
import UserImage from "./UserImage";
import { selectCurrentUser , setFriends } from "../features/auth/authSlice";
import { useRemoveFriendFromUserMutation, useAddFriendToUserMutation} from '../features/users/usersApiSlice'
import { useEffect, useState } from "react";

const Friend = ({ friendId, name, subtitle, userPicturePath }) => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const {id} = useSelector(selectCurrentUser)
  const {friends} =  useSelector(selectCurrentUser)

  const [removeFriendFromUser, {isSuccess: isSuccessRemoveFriend}] = useRemoveFriendFromUserMutation()
  const [addFriendToUser, {isSuccess: isSuccessAddFriend}] = useAddFriendToUserMutation()

  const { palette } = useTheme();
  const primaryLight = palette.primary.light;
  const primaryDark = palette.primary.dark;
  const main = palette.neutral.main;
  const medium = palette.neutral.medium;
  const [isFriend, setIsFriend] = useState(false)



  useEffect(()=>{
    if(friends?.length > 0){
      setIsFriend(friends.find((friend) => friend.id === friendId))
    }
    // patchFriend()    
  },[isFriend, friendId,friends])


const patchFriend = async () =>{

  const userData = {
    userID: id,
    friendID: friendId
  }

  if(isFriend){
    const dataFriendRemove = await removeFriendFromUser(userData)
    if (isSuccessRemoveFriend){
      dispatch(setFriends({ ...dataFriendRemove })) 
    }
  } else if (!isFriend){
    const dataFriendAdd = await addFriendToUser(userData)
    if (isSuccessAddFriend){
      dispatch(setFriends({ ...dataFriendAdd })) 
    }
  }
}


  return (
    <FlexBetween>
      <FlexBetween gap="1rem">
        <UserImage image={userPicturePath} size="55px" />
        <Box
          onClick={() => {
            navigate(`/profile/${friendId}`);
            navigate(0);
          }}
        >
          <Typography
            color={main}
            variant="h5"
            fontWeight="500"
            sx={{
              "&:hover": {
                color: palette.primary.light,
                cursor: "pointer",
              },
            }}
          >
            {name}
          </Typography>
          <Typography color={medium} fontSize="0.75rem">
            {subtitle}
          </Typography>
        </Box>
      </FlexBetween>
      <IconButton
        onClick={() => patchFriend()}
        sx={{ backgroundColor: primaryLight, p: "0.6rem" }}
      >
        {isFriend ? (
          <PersonRemoveOutlined sx={{ color: primaryDark }} />
        ) : (
          <PersonAddOutlined sx={{ color: primaryDark }} />
        )}
      </IconButton>
    </FlexBetween>
  );
};

export default Friend;