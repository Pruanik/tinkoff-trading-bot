import { IPostActionSetInstrumentObservable } from "./IPostActionSetInstrumentObservable";

export class PostActionSetInstrumentObservable implements IPostActionSetInstrumentObservable {
    figi: string;

    constructor(data: IPostActionSetInstrumentObservable) {
        this.figi = data.figi;
    }
}
