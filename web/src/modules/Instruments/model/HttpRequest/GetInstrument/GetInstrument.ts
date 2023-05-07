import { IGetInstrument } from "./IGetInstrument";

export class GetInstrument implements IGetInstrument {
    sectorId: number;
    type: string;

    constructor(data: IGetInstrument) {
        this.sectorId = data.sectorId;
        this.type = data.type;
    }
}
