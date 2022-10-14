import React from "react";
import { Link as RouterLink } from "react-router-dom";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { UsersInterface } from "../models/IUser";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import Autocomplete from '@mui/lab/Autocomplete';
//import lightgreen from '@mui/material/colors'

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
 props,
 ref
) {
 return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function UserCreate() {
 const [date, setDate] = React.useState<Date | null>(null);
 const [user, setUser] = React.useState<Partial<UsersInterface>>({});
 const [success, setSuccess] = React.useState(false);
 const [error, setError] = React.useState(false);
 const handleClose = (
   event?: React.SyntheticEvent | Event,
   reason?: string
 ) => {
   if (reason === "clickaway") {
     return;
   }
   setSuccess(false);
   setError(false);
 };


 const handleInputChange = (
   event: React.ChangeEvent<{ id?: string; value: any }>
 ) => {
   const id = event.target.id as keyof typeof UserCreate;
   const { value } = event.target;
   setUser({ ...user, [id]: value });
 };


 function submit() {
   let data = {
     Name: user.Name ?? "",
     //LastName: user.LastName ?? "",
     //Age: typeof user.Age === "string" ? parseInt(user.Age) : 0,
     BirthDay: date,
   };


   const apiUrl = "http://localhost:8080/users";
   const requestOptions = {
     method: "POST",
     headers: { "Content-Type": "application/json" },
     body: JSON.stringify(data),
   };


   fetch(apiUrl, requestOptions)
     .then((response) => response.json())
     .then((res) => {
       if (res.data) {
         setSuccess(true);
       } else {
         setError(true);
       }
     });
 }


 return (
   <Container maxWidth="md">
     <Snackbar
       open={success}
       autoHideDuration={6000}
       onClose={handleClose}
       anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
     >
       <Alert onClose={handleClose} severity="success">
         บันทึกข้อมูลสำเร็จ
       </Alert>
     </Snackbar>
     <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
       <Alert onClose={handleClose} severity="error">
         บันทึกข้อมูลไม่สำเร็จ
       </Alert>
     </Snackbar>

     
     <Paper>
       <Box
         display="flex"
         sx={{
           marginTop: 2,
         }}
       >
         <Box sx={{ paddingX: 26, paddingY: 1 }}>
           <Typography
             component="h2"
             variant="h4"
             color= "#558b2f"
             gutterBottom
             //align="center"
             fontFamily="Arial"
           >
             ระบบบันทึกข้อมูลผู้ป่วยใน
           </Typography>

         </Box>
       </Box>
       <Divider />


       <Grid container spacing={3} sx={{ padding: 2 }}>


         <Grid item xs={12}>
           <p>รหัสประจำตัวประชาชน</p>
           <FormControl fullWidth variant="outlined">
             <TextField
               id="ID_Card"
               variant="outlined"
               type="string"
               size="medium"
               value={user.ID_Card || ""}
               onChange={handleInputChange}     
             />
           </FormControl>
         </Grid>




         <Grid item xs={12}>
           <FormControl fullWidth variant="outlined">
             <p>ชื่อ-สกุล</p>
             <TextField
               id="Name"
               variant="outlined"
               type="string"
               size="medium"
               value={user.Name || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>





         {/* <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p>วันเกิด</p>
             <TextField
               id="BirthDay"
               variant="outlined"
               type="number"
               size="medium"
               InputProps={{ inputProps: { min: 1 } }}
               InputLabelProps={{
                 shrink: true,
               }}
               value={user.BirthDay || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid> */}



         <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p>วันเกิด</p>
             <LocalizationProvider dateAdapter={AdapterDateFns}>
               <DatePicker
                 value={date}
                 onChange={(newValue) => {
                   setDate(newValue);
                 }}
                 renderInput={(params) => <TextField {...params} />}
               />
             </LocalizationProvider>
           </FormControl>
         </Grid>



        {/* test combobox1 */}
         <Grid item xs={12}>
         <FormControl fullWidth variant="outlined">
          <p>Test combobox</p>
            <Autocomplete
              id="combo-box-demo"
              options={top100Films}
              getOptionLabel={(option) => option.title}
              style={{ width: 300 }}
              renderInput={(params) => <TextField {...params} variant="outlined" />}
              //renderInput={(params) => <TextField {...params} label="Combo box" variant="outlined" />}
            />
         </FormControl>
         </Grid>



         <Grid item xs={12}>
           <Button component={RouterLink} to="/" variant="contained">
             Back
           </Button>
           <Button
             style={{ float: "right" }}
             onClick={submit}
             variant="contained"
             color="primary"
           >
             Submit
           </Button>
         </Grid>



        <Grid item xs={12}>
          <p>Test text laout</p>
        </Grid>



       </Grid>
     </Paper>
   </Container>
 );
}
export default UserCreate;



    

// data in combobox1
const top100Films = [
  { title: 'The Shawshank Redemption', year: 1994 },
  { title: 'The Godfather', year: 1972 },
  { title: 'The Godfather: Part II', year: 1974 },
  { title: 'The Dark Knight', year: 2008 },
  { title: '12 Angry Men', year: 1957 },
];