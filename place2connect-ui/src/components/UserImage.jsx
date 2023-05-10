import { Box } from "@mui/material";

const API_IMG =  import.meta.env.VITE_REACT_APP_BACKEND_IMG;

const UserImage = ({image, size ="60px"}) => {
  return (
    <Box >
      <img 
          style={{ objectFit : "cover", borderRadius: "50%" }}
          width={size}
          height={size}
          alt="user"
          src={`${API_IMG}/${image}`}
          />
    </Box>
  )
}

export default UserImage