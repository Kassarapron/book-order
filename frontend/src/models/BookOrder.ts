// import { DatePickerView } from "@material-ui/pickers";
// import { NamedTupleMember } from "typescript";
import { AdminInterface } from "./Admin";
import { BookInterface } from "./Book";
// import { BookTypeInterface } from "./BookType";
import { CompanyInterface } from "./Company";

export interface BookOrderInterface {
    ID          : string,
    CreatedAt   : string,
    UpdatedAt   : string,
    DeletedAt?  : string,
    Quantity    : number,
    AdminId     : number,
    Admin       : AdminInterface,
    CompanyId   : number,
    Company     : CompanyInterface,
    BookId      : number,
    Book        : BookInterface,
}