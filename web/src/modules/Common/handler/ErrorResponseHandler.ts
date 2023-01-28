import axios from 'axios';

export function handleErrorResponse(reason: unknown): never {
    if (
        axios.isAxiosError(reason) &&
        reason.response?.data.Status.Message !== undefined &&
        reason.response?.data.Status.Message !== null
    ) {
        throw new Error(reason.response.data.message);
    } else {
        throw new Error(
            'Произошла непредвиденная ошибка, обратитесь, пожалуйста, в службу поддержки' +
                ' или попробуйте повторить действие позже'
        );
    }
}
