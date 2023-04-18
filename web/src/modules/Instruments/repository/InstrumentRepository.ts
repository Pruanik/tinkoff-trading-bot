import HttpClient from "@/http/HttpClient";
import { IResponse } from "@/modules/Common/model/IResponse";
import { IInstrumentSector } from "../model/IInstrumentSector";

export default {
    getInstrumentSectorList: async (): Promise<IResponse<IInstrumentSector[]>> => {
        const response = await HttpClient.get<IResponse<IInstrumentSector[]>>('/instruments/sectors');
        return response.data;
    },
};
