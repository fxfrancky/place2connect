import { Box, useMediaQuery } from "@mui/material";
import { useParams } from "react-router-dom";
import Navbar from "../../scenes/navbar";
import FriendListWidget from "../../scenes/widgets/FriendListWidget";
import MyPostWidget from "../../scenes/widgets/MyPostWidget";
import PostsWidget from "../../scenes/widgets/PostsWidget";
import UserWidget from "../../scenes/widgets/UserWidget";
import {useGetUserQuery} from '../../features/users/usersApiSlice'
import { useDispatch } from 'react-redux';
import { apiSlice } from '../../app/api/apiSlice'

const ProfilePage = () => {
  // const [user, setUser] = useState(null);
  const { userID } = useParams();
  const dispatch = useDispatch();
  const isNonMobileScreens = useMediaQuery("(min-width:1000px)");
   // Get the user By ID
  const {
  data:user,
  isSuccess,
  isError,
  error,
  isLoading
} = useGetUserQuery(userID) 

  // setUser(response)
  dispatch(apiSlice.util.resetApiState());


  if (!user){
    console.log(`user not found with id `, userID)
    return null;
  }

  // const getUser = async () => {
  //   const response = await fetch(`http://localhost:3001/users/${userId}`, {
  //     method: "GET",
  //     headers: { Authorization: `Bearer ${token}` },
  //   });
  //   const data = await response.json();
  //   setUser(data);
  // };

  // useEffect(() => {
  //   getUser();
  // }, []); // eslint-disable-line react-hooks/exhaustive-deps
  console.log('IS Success', isSuccess)
  console.log('IS Error', isError, error)
  console.log('IS loading', isLoading)
 
  {isError && (<p>Saddly an error occured {error.message}</p>)}
  {isLoading && (<p>Wait a minute. The page is still loading</p>)}

  if (!user) return null;


  return ( <Box>
      <Navbar />
      <Box
        width="100%"
        padding="2rem 6%"
        display={isNonMobileScreens ? "flex" : "block"}
        gap="2rem"
        justifyContent="center"
      >
        { isSuccess && (
          <>
          <Box flexBasis={isNonMobileScreens ? "26%" : undefined}>
          <UserWidget userID={userID} picturePath={user.picturePath} />
          <Box m="2rem 0" />
          <FriendListWidget userID={userID} />
        </Box>
        <Box
          flexBasis={isNonMobileScreens ? "42%" : undefined}
          mt={isNonMobileScreens ? undefined : "2rem"}
        >
          <MyPostWidget picturePath={user.picturePath} />
          <Box m="2rem 0" />
          <PostsWidget userID={userID} isProfile />
        </Box>
        </>
        )}  
      </Box>
    </Box>
  );}

  // return (
  //   <Box>
  //     <Navbar />
  //     <Box
  //       width="100%"
  //       padding="2rem 6%"
  //       display={isNonMobileScreens ? "flex" : "block"}
  //       gap="2rem"
  //       justifyContent="center"
  //     >
  //       { isSuccess && (
  //         <>
  //         <Box flexBasis={isNonMobileScreens ? "26%" : undefined}>
  //         <UserWidget userID={userID} picturePath={user.picturePath} />
  //         <Box m="2rem 0" />
  //         <FriendListWidget userID={userID} />
  //       </Box>
  //       <Box
  //         flexBasis={isNonMobileScreens ? "42%" : undefined}
  //         mt={isNonMobileScreens ? undefined : "2rem"}
  //       >
  //         <MyPostWidget picturePath={user.picturePath} />
  //         <Box m="2rem 0" />
  //         <PostsWidget userID={userID} isProfile />
  //       </Box>
  //       </>
  //       )}   



  //     </Box>
  //   </Box>
  // );
// };

export default ProfilePage;