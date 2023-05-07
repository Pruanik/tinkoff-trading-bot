import HttpClient from "@/http/HttpClient";
import { IResponse } from "@/modules/Common/model/IResponse";
import { IInstrumentSector } from "../model/IInstrumentSector";
import { IGetInstrumentType } from "../model/HttpRequest/GetInstrumentType/IGetInstrumentType";
import HttpGetQueryFormatter from "@/modules/Common/service/HttpGetQueryFormatter";
import { IInstrumentType } from "../model/IInstrumentType";
import { IGetInstrument } from "../model/HttpRequest/GetInstrument/IGetInstrument";
import { IInstrument } from "../model/IInstrument";
import { IPostActionSetInstrumentObservable } from "../model/HttpRequest/PostActionSetInstrumentObservable/IPostActionSetInstrumentObservable";

export default {
    getInstrumentSectorList: async (): Promise<IResponse<IInstrumentSector[]>> => {
        const response = await HttpClient.get<IResponse<IInstrumentSector[]>>('/instruments/sectors');
        return response.data;
    },

    getInstrumentTypeList: async (getInstrumentTypeRequest: IGetInstrumentType): Promise<IResponse<IInstrumentType[]>> => {
        const response = await HttpClient.get<IResponse<IInstrumentType[]>>(
            '/instruments/types' + HttpGetQueryFormatter.getGetQueryString(getInstrumentTypeRequest)
        );
        return response.data;
    },

    getInstrumentList: async (getInstrumentRequest: IGetInstrument): Promise<IResponse<IInstrument[]>> => {
        const response = await HttpClient.get<IResponse<IInstrument[]>>(
            '/instruments' + HttpGetQueryFormatter.getGetQueryString(getInstrumentRequest)
        );
        return response.data;
    },

    setInstrumentObservable: async(postActionSetInstrumentObservable: IPostActionSetInstrumentObservable): Promise<IResponse<null>> => {
        const response = await HttpClient.post<IResponse<null>>(
            '/instruments/action/setInstrumentObservable',
            postActionSetInstrumentObservable
        );

        return response.data;
    }
};
