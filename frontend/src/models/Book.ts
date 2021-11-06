import { BookTypeInterface } from "./BookType";

export interface BookInterface {
    ID              : number,
    CreatedAt       : string,
    UpdatedAt       : string,
    DeletedAt?      : string,
    BookName        : string,
    BookNumber      : string,
    BookPublicher   : string,
    BookOrders?     : any,
    BookTypeId      : number,
    BookType        : BookTypeInterface,
}