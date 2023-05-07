import { IGetInstrumentType } from "./IGetInstrumentType";

export class GetInstrumentType implements IGetInstrumentType {
    sectorId: number;

    constructor(data: IGetInstrumentType) {
        this.sectorId = data.sectorId;
    }
}
