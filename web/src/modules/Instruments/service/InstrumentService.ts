import instrumentRepository from '../repository/InstrumentRepository';
import { handleErrorResponse } from '@/modules/Common/handler/ErrorResponseHandler';
import { IInstrumentSector } from '../model/IInstrumentSector';
import { IInstrumentType } from '../model/IInstrumentType';
import { GetInstrumentType } from '../model/HttpRequest/GetInstrumentType/GetInstrumentType';
import { GetInstrument } from '../model/HttpRequest/GetInstrument/GetInstrument';
import { PostActionSetInstrumentObservable } from '../model/HttpRequest/PostActionSetInstrumentObservable/PostActionSetInstrumentObservable';


export default {
    getInstrumentSectorList: async (): Promise<IInstrumentSector[] | undefined> | never => {
        try {
            return await instrumentRepository.getInstrumentSectorList().then(data => data.Body);
        } catch (e) {
            handleErrorResponse(e);
        }
    },

    getInstrumentTypeList: async (sectorId: number | undefined = null): Promise<IInstrumentType[] | undefined> | never => {
        try {
            const request = new GetInstrumentType({sectorId});
            return await instrumentRepository.getInstrumentTypeList(request).then(data => data.Body);
        } catch (e) {
            handleErrorResponse(e);
        }
    },

    getInstrumentList: async (
        sectorId: number | undefined = null,
        type: string | undefined = null,
    ): Promise<IInstrument[] | undefined> | never => {
        try {
            const request = new GetInstrument({sectorId, type});
            return await instrumentRepository.getInstrumentList(request).then(data => data.Body);
        } catch (e) {
            handleErrorResponse(e);
        }
    },

    setInstrumentObservable: async (figi: string) => {
        try {
            const request = new PostActionSetInstrumentObservable({figi});
            await instrumentRepository.setInstrumentObservable(request);
        } catch (e) {
            handleErrorResponse(e);
        }
    }
}
