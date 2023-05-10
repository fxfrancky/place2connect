import {
  ChatBubbleOutlineOutlined,
  FavoriteBorderOutlined,
  FavoriteOutlined,
  ShareOutlined,
} from "@mui/icons-material";
import { Box, Divider, IconButton, Typography, useTheme } from "@mui/material";
import FlexBetween from "../../components/FlexBetween";
import Friend from "../../components/Friend";
import WidgetWrapper from "../../components/WidgetWrapper";
import { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { selectCurrentUser, setPost } from "../../features/auth/authSlice";
import { useAddLikeToPostMutation } from '../../features/posts/postsApiSlice'

const API_IMG =  import.meta.env.VITE_REACT_APP_BACKEND_IMG;

const PostWidget = ({
  postId,
  postUserId,
  name,
  description,
  location,
  picturePath,
  userPicturePath,
  likes,
  comments,
}) => {

  const [isComments, setIsComments] = useState(false);
  const dispatch = useDispatch();
  const loggedInUser = useSelector(selectCurrentUser)
  const {id} = useSelector(selectCurrentUser)
  const [isLiked, setIsLiked] = useState(false)
  const [likeCount, setLikeCount] = useState(0)

  useEffect(() =>{
    if(likes?.length > 0){
      const found = likes.some(lk => lk.user_id === id);
      setIsLiked(found)
      setLikeCount(Object.keys(likes).length)
    } else {
      setIsLiked(false)
      setLikeCount(0)
    }
  },[isLiked,likeCount,likes,id])

  const { palette } = useTheme();
  const main = palette.neutral.main;
  const primary = palette.primary.main;

  const [addLikeToPost,{isSuccess}] = useAddLikeToPostMutation()

  const postLikeData = {
    isLikablePost:isLiked,
    postID: postId,
    userID: loggedInUser.id
  }

  const patchLike = async () =>{
      const postLikedResponse = await addLikeToPost(postLikeData).unwrap()
      if(isSuccess){
        dispatch(setPost({...postLikedResponse}))
        setIsLiked(!isLiked)
      }
  }


  return (
    <WidgetWrapper m="2rem 0">
      <Friend
        friendId={postUserId}
        name={name}
        subtitle={location}
        userPicturePath={userPicturePath}
      />
      <Typography color={main} sx={{ mt: "1rem" }}>
        {description}
      </Typography>
      {picturePath && (
        <img
          width="100%"
          height="auto"
          alt="post"
          style={{ borderRadius: "0.75rem", marginTop: "0.75rem" }}
          src={`${API_IMG}/${picturePath}`}
        />
      )}
      <FlexBetween mt="0.25rem">
        <FlexBetween gap="1rem">
          <FlexBetween gap="0.3rem">
            <IconButton onClick={patchLike}>
              {isLiked ? (
                <FavoriteOutlined sx={{ color: primary }} />
              ) : (
                <FavoriteBorderOutlined />
              )}
            </IconButton>
            <Typography>{likeCount}</Typography>
          </FlexBetween>

          <FlexBetween gap="0.3rem">
            <IconButton onClick={() => setIsComments(!isComments)}>
              <ChatBubbleOutlineOutlined />
            </IconButton>
            <Typography>{comments ? comments.length : 0}</Typography>
          </FlexBetween>
        </FlexBetween>

        <IconButton>
          <ShareOutlined />
        </IconButton>
      </FlexBetween>
      {isComments && (
        <Box mt="0.5rem">
          {comments?.length > 0 && (comments.map((comment, i) => (
            <Box key={`${name}-${i}`}>
              <p>{name}</p>            
              <Divider />
              <Typography sx={{ color: main, m: "0.5rem 0", pl: "1rem" }}>
                {comment.comment_description}
              </Typography>
            </Box>
          )))}
          <Divider />
        </Box>
      )}
    </WidgetWrapper>
  );
};

export default PostWidget;