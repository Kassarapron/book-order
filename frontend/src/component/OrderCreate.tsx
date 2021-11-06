import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {makeStyles,Theme,createStyles,alpha,} from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";

import { BookOrderInterface } from "../models/BookOrder";
import { BookInterface } from "../models/Book";
import { BookTypeInterface } from "../models/BookType";
import { CompanyInterface } from "../models/Company";
import { AdminInterface } from "../models/Admin";

import { TextField } from "@material-ui/core";

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    container: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
    },
  })
);

function OrderCreate() {
 const classes = useStyles();
 const [admins,    setAdmin] = useState<AdminInterface[]>([]);
 const [books,     setBook] = useState<BookInterface[]>([]);
 const [book_types, setBookType] = useState<BookTypeInterface[]>([]);
 const [companies, setCompanies] = useState<CompanyInterface[]>([]);
 const [bookOrder, setBookOrder] = useState<any>({});
 
 const [success, setSuccess] = useState(false);
 const [error, setError] = useState(false);

 const apiUrl = "http://localhost:8080";
 const requestOptions = {
   method: "GET",
   headers: {
     Authorization: `Bearer ${localStorage.getItem("token")}`,
     "Content-Type": "application/json",
   },
 };

 const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
   if (reason === "clickaway") {
     return;
   }
   setSuccess(false);
   setError(false);
 };

 const handleChange = (
   event: React.ChangeEvent<{ name?: string; value: unknown }>
 ) => {
   const name = event.target.name as keyof typeof bookOrder;
   const val = event.target.value
   setBookOrder({
     ...bookOrder,
     [name]: typeof val === 'string' ? parseInt(val) : undefined,
   });
 };

 const getAdmin = async () => {
  fetch(`${apiUrl}/admin`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      setAdmin(res);
    });
  };  
  const getCompany = async () => {
    fetch(`${apiUrl}/company`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        setCompanies(res);
      });
    }; 
 const getBook = async () => {
  fetch(`${apiUrl}/book`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      setBook(res);
    });
  }; 
  const getBookType = async () => {
  fetch(`${apiUrl}/book-type`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      setBookType(res);
    });
  }; 

 useEffect(() => {
  getAdmin();
  getCompany();
  getBook();
  getBookType();
  }, []);

function submit() {
  if(bookOrder.AdminID && bookOrder.Quantity && bookOrder.BookID && bookOrder.CompanyID){
     const requestOptionsPost = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(bookOrder),
    };
  
    fetch(`${apiUrl}/book-order`, requestOptionsPost)
        .then((response) => response.json())
        .then((res) => {
          if (res) {
            setSuccess(true);
          } else {
            setError(true);
          }
        });
  }else{
    setError(true);
  }
}

return (
  <Container className={classes.container} maxWidth="md">
    <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
      <Alert onClose={handleClose} severity="success">
        บันทึกข้อมูลสำเร็จ
      </Alert>
    </Snackbar>
    <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
      <Alert onClose={handleClose} severity="error">
        บันทึกข้อมูลผิดพลาด
      </Alert>
    </Snackbar>
    <Paper className={classes.paper}>
      <Box display="flex">
        <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
             Create Order
           </Typography>
         </Box>
       </Box>
       <Divider />
       <Grid container spacing={3} className={classes.root}>
         
       <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>แอดมิน</p>
              <Select
                native
                // value={book_orders.AdminID}
                onChange={handleChange}
                inputProps={{
                  name: "AdminID",
                }}
                defaultValue="none"
              >
                <option aria-label="None" value="none" disabled>
                  กรุณาเลือกแอดมิน
                </option>
                {admins.map((admin: AdminInterface) => (
                  <option value={admin.ID} key={admin.ID}>
                    {admin.AdminName}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>บริษัท</p>
              <Select
                native
                value={bookOrder.ID}
                onChange={handleChange}
                inputProps={{
                  name: "CompanyID",
                }}
                defaultValue="none"
              >
                <option aria-label="None" value="none" disabled>
                  กรุณาเลือกบริษัท
                </option>
                {companies.map((company: CompanyInterface) => (
                  <option value={company.ID} key={company.ID}>
                    {company.CompanyName}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ชื่อหนังสือ</p>
              <Select
                native
                // value = {book_orders.ID}
                onChange = {handleChange}
                inputProps={{
                name: "BookID",}}
                defaultValue="none"
              >
                <option aria-label="None" value="none" disabled>
                  กรุณาเลือกหนังสือ
                </option>
                {books.map((book: BookInterface) => (
                <option value={book.ID} key={book.ID}>
                  {book.ID}.{book.BookName}
                  </option>
                  ))}
                  </Select>
                  </FormControl>
          </Grid>

         <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
              <p>จำนวน</p>
              <TextField            
                id="quantity"
                variant="outlined"
                type="number"
                size="medium"
                name="Quantity"
                placeholder="กรุณาเลือกจำนวน"
                // value={book_orders.Quantity || ""}
                onChange={handleChange}
              />
            </FormControl>
          </Grid>

         <Grid item xs={12}>
           <Button component={RouterLink} to="/Orders" variant="contained">
             Order Table
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
       </Grid>
     </Paper>
   </Container>
 );
}
export default OrderCreate;