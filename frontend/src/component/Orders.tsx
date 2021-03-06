import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import { BookOrderInterface } from "../models/BookOrder";
//import moment from 'moment';
import { format } from "date-fns";
import moment from "moment";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: { marginTop: theme.spacing(2) },
    table: { minWidth: 800 },
    tableSpace: { marginTop: 20 },
  })
);

function Orders() {
  const classes = useStyles();
  const [Orders, setOrders] = React.useState<BookOrderInterface[]>([]);
  const apiUrl = "http://localhost8080/Orders";
  const getOrders = async () => {
      const apiUrl = "http://localhost:8080/Orders";
      const requestOptions = {
      method: "GET",
      headers: { "Content-Type": "application/json" },
    };

    fetch(apiUrl, requestOptions)
       .then((response) => response.json())
       .then((res) => {
         console.log(res.data);
         if (res.data) {
           setOrders(res.data);
         } else {
           console.log("else");
         }
       });
  };

  useEffect(() => {
    getOrders();
  }, []);
  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              Orders
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/OrderCreate"
              variant="contained"
              color="primary"
            >
              Create Order
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="5%">
                  ID
                </TableCell>
                <TableCell align="center" width="20%">
                  Email
                </TableCell>
                <TableCell align="center" width="20%">
                  ??????????????????????????????
                </TableCell>
                <TableCell align="center" width="15%">
                  ???????????????????????????????????????
                </TableCell>
                <TableCell align="center" width="20%">
                  ?????????????????????????????????
                </TableCell>
                <TableCell align="center" width="5%">
                  ???????????????
                </TableCell>
                <TableCell align="center" width="15%">
                  ??????????????????????????????
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {Orders.map((order: BookOrderInterface) => (
                <TableRow key={order.ID}>
                  <TableCell align="right"  size="medium"> {order.ID}           </TableCell>
                  <TableCell align="left"   size="medium"> {order.AdminName}    </TableCell>
                  <TableCell align="left"   size="medium"> {order.CompanyName}  </TableCell>
                  <TableCell align="left"   size="medium"> {order.BookTypeName} </TableCell>
                  <TableCell align="left"   size="medium"> {order.BookName}     </TableCell>
                  <TableCell align="right"  size="medium"> {order.Quantity}     </TableCell>
                  <TableCell align="center" size="medium"> {moment(order.Date).format("DD/MM/YYYY")}    </TableCell>             
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}
export default Orders;