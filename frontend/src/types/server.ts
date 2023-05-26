export class Server {
    constructor(iconURL: string, iconAlt: string, serverName: string, upStatus: boolean) {
        this.iconURL = iconURL;
        this.iconAlt = iconAlt;
        this.serverName = serverName;
        this.upStatus = upStatus;
    }

    iconURL: string;
    iconAlt: string;
    serverName: string = "undefined";
    upStatus: boolean;
}