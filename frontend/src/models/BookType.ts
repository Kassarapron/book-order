import { BookInterface } from "./Book";
export interface BookTypeInterface {
    ID      : number,
    BTName  : string;

    BookID  : number;
    Bookname: BookInterface;
}