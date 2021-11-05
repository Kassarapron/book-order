export interface AdminInterface {
    ID          : number,
    CreatedAt   : string,
    UpdatedAt   : string,
    DeletedAt?  : string,
    AdminName   : string,
    Email       : String,
    BookOrders? : any,
}