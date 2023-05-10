import { useDispatch, useSelector } from "react-redux"
import PostWidget from './PostWidget'
import { selectCurrentPosts,setPosts } from "../../features/auth/authSlice";
import { useGetPostsQuery, useGetUserPostsQuery } from '../../features/posts/postsApiSlice'
import { useEffect } from "react";


const PostsWidget = ({ userID, isProfile = false}) => {
  
  const dispatch = useDispatch();
  const posts = useSelector(selectCurrentPosts)

  // FeedPosts (All Posts)
  const postQuery = useGetPostsQuery()
  // User Posts (Only User Posts)
  const userPostsQuery = useGetUserPostsQuery(userID)

  useEffect(() =>{ 
    if(isProfile && userPostsQuery.isSuccess){ 
      // console.log(userPostsQuery.data)
      if(userPostsQuery.data){
        dispatch(setPosts(userPostsQuery.data));
      }
      
    } else if (!isProfile && postQuery.isSuccess){ 
      // dispatch(setPosts({...postQuery.data })); 
      //  console.log(userPostsQuery.data)
      if(userPostsQuery.data){
        dispatch(setPosts(postQuery.data));
      }
    }
  },[userID,isProfile,posts,dispatch, postQuery.data, postQuery.isSuccess, userPostsQuery.data, userPostsQuery.isSuccess]);

return (  

    <>
      {posts && (posts.map(
        ({
          id,
          userID,
          firstName,
          lastName,
          description,
          location,
          picturePath,
          userPicturePath,
          likes,
          comments,
        }) => (
          <PostWidget
            key={id}
            postId={id}
            postUserId={userID}
            name={`${firstName} ${lastName}`}
            description={description}
            location={location}
            picturePath={picturePath}
            userPicturePath={userPicturePath}
            likes={likes}
            comments={comments}
          />
        )
      )
      
      )}
    </>
  )

};
export default PostsWidget