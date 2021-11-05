// import { DatePickerView } from "@material-ui/pickers";
// import { NamedTupleMember } from "typescript";
import { AdminInterface } from "./Admin";
import { BookInterface } from "./Book";
import { BookTypeInterface } from "./BookType";
import { CompanyInterface } from "./Company";

export interface BookOrderInterface {
    ID          : number,
    Quantity    : string;
    Date        : Date;

    AdminID     : number;
    AdminName   : AdminInterface;

    BookID      : number;
    BookName    : BookInterface;

    BookTypeID  : number;
    BookTypeName: BookTypeInterface;

    CompanyID   : number;
    CompanyName : CompanyInterface;
    
}