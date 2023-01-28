export interface IResponse<T> {
    Status: {
        Status: string;
        Message: string;
    };
    Body: T | null;
    Time: string,
}
