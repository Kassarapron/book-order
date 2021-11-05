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
import Orders from "./Orders";
import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";
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
 const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
 const [admins,    setAdmin] = useState<AdminInterface[]>([]);
 const [books,     setBook] = useState<BookInterface[]>([]);
 const [book_types, setBookType] = useState<BookTypeInterface[]>([]);
 const [companys, setCompany] = useState<CompanyInterface[]>([]);
 const [book_orders, setBookOrder] = useState<Partial<BookOrderInterface>>({});
 
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
 const handleInputChange = (
  event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof book_orders;
    const { value } = event.target;
    setBookOrder({ ...book_orders, [id]: value });
  };

 const handleChange = (
   event: React.ChangeEvent<{ name?: string; value: unknown }>
 ) => {
   const name = event.target.name as keyof typeof book_orders;
   setBookOrder({
     ...book_orders,
     [name]: event.target.value,
   });
 };

 const handleDateChange = (date: Date | null) => {
   console.log(date);
   setSelectedDate(date);
 };

 const getAdmin = async () => {
  fetch(`${apiUrl}/category`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {setAdmin(res.data);} 
      else {console.log("else");}
    });
  };  
  const getCompany = async () => {
    fetch(`${apiUrl}/category`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {setCompany(res.data);} 
        else {console.log("else");}
      });
    }; 
 const getBook = async () => {
  fetch(`${apiUrl}/category`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {setBook(res.data);} 
      else {console.log("else");}
    });
  }; 
  const getBookType = async () => {
  fetch(`${apiUrl}/category`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {setBookType(res.data);} 
      else {console.log("else");}
    });
  }; 

 useEffect(() => {
  getAdmin();
  getCompany();
  getBook();
  getBookType();
  }, []);
  
 const convertType = (data: string | number | undefined) => {
  let val = typeof data === "string" ? parseInt(data) : data;
  return val;
  };  

function submit() {
   let data = {
    //  AdminID : convertType(book_orders.AdminID),
    //  CompanyID : convertType(book_orders.CompanyID),
    //  BookID : convertType(book_orders.BookID),
    //  BookTypeID : convertType(book_orders.BookTypeID),
    //  Quantity : book_orders.Quantity ?? "",
   };

   console.log(data)

   const requestOptionsPost = {
    method: "POST",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  };

  fetch(`${apiUrl}/bookorder`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true);
        } else {
          console.log("บันทึกผิดพลาด")
          setError(true);
        }
      });
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
              <p>Email</p>
              <Select
                native
                // value={book_orders.AdminID}
                onChange={handleChange}
                inputProps={{
                  name: "VideoID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกวีดีโอ
                </option>
                {admins.map((item: AdminInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {/* {item.AName} */}
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
                value={book_orders.ID}
                onChange={handleChange}
                inputProps={{
                  name: "ID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกบริษัท
                </option>
                {companys.map((item: CompanyInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {/* {item.CName} */}
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
                value = {book_orders.ID}
                onChange = {handleChange}
                inputProps={{
                name: "BookID",}}
              >
                <option aria-label="None" value="">กรุณาเลือกหนังสือ</option>
                {books.map((item: BookInterface) => (
                <option value={item.ID} key={item.ID}>
                  {/* {item.ID}.{item.BName} */}
                  </option>
                  ))}
                  </Select>
                  </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ประเภทหนังสือ</p>
              <Select
              native
              // value={book_orders.BookTypeID}
              onChange={handleChange}
              inputProps={{
                name: "BookTypeID",
                }}
                >
                  <option aria-label="None" value="">
                    กรุณาเลือกประเภท
                    </option>
                    {book_types.map((item: BookTypeInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {/* {item.ID}.{item.BTName} */}
                      </option>
                      ))}
                      </Select>
                      </FormControl>
                      </Grid>

          {/* <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ประเภทหนังสือ</p>
              <Select
              native
              value={book_orders.BookTypeID}
              onChange={handleChange}
              inputProps={{
              name: "BookTypeID",
              }}
              >
              <option aria-label="None" value="">
                กรุณาระบุหนังสือ
              </option>
              {book_types.map((item: BookTypeInterface) => (
                <option value={item.BTName} key={item.ID}>
                รหัสหนังสือ {item.ID} ประเภทหนังสือ {item.BTName}
              </option>
              ))}
              </Select>
            </FormControl>
          </Grid> */}

          {/* <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ประเภทหนังสือ</p>
              <Select
                native
                value = {book_orders.ID}
                onChange = {handleChange}
                inputProps={{
                name: "BookTypeID",}}
              >
                <option aria-label="None" value="">กรุณาเลือกประเภทหนังสือ</option>
                {book_types.map((item: BookTypeInterface) => (
                <option value={item.ID} key={item.ID}>
                  {item.ID}.{item.BTName}
                  </option>
                ))}
              </Select>
            </FormControl>
        </Grid> */}

         <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
              <p>จำนวน</p>
              <TextField            
                id="quantity"
                variant="outlined"
                type="uint"
                size="medium"
                value={book_orders.Quantity || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่</p>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  name="Date"
                  value={selectedDate}
                  onChange={handleDateChange}
                  label="กรุณาเลือกวันที่"
                  minDate={new Date("2018-01-01")}
                  format="yyyy/MM/dd"
                />
              </MuiPickersUtilsProvider>
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