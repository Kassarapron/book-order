export interface CompanyInterface {
    ID          : number,
    CreatedAt   : string,
    UpdatedAt   : string,
    DeletedAt?  : string,
    CompanyName : string;
    BookOrders? : any,
}