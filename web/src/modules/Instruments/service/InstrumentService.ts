import instrumentRepository from '../repository/InstrumentRepository';
import { handleErrorResponse } from '@/modules/Common/handler/ErrorResponseHandler';
import { IInstrumentSector } from '../model/IInstrumentSector';


export default {
    getInstrumentSectorList: async (): Promise<IInstrumentSector[] | undefined> | never => {
        try {
            return await instrumentRepository.getInstrumentSectorList().then(data => data.Body);
        } catch (e) {
            handleErrorResponse(e);
        }
    },
}
