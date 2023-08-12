export namespace gofeed {
	
	export class Enclosure {
	    url?: string;
	    length?: string;
	    type?: string;
	
	    static createFrom(source: any = {}) {
	        return new Enclosure(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.length = source["length"];
	        this.type = source["type"];
	    }
	}
	export class Image {
	    url?: string;
	    title?: string;
	
	    static createFrom(source: any = {}) {
	        return new Image(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.title = source["title"];
	    }
	}

}

export namespace news {
	
	export class Item {
	    title?: string;
	    description?: string;
	    content?: string;
	    link?: string;
	    links?: string[];
	    updated?: string;
	    // Go type: time
	    updated_parsed?: any;
	    published?: string;
	    // Go type: time
	    published_parsed?: any;
	    guid?: string;
	    image?: gofeed.Image;
	    categories?: string[];
	    enclosures?: gofeed.Enclosure[];
	    stupid_enclosure?: gofeed.Enclosure;
	    custom?: {[key: string]: string};
	
	    static createFrom(source: any = {}) {
	        return new Item(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.description = source["description"];
	        this.content = source["content"];
	        this.link = source["link"];
	        this.links = source["links"];
	        this.updated = source["updated"];
	        this.updated_parsed = this.convertValues(source["updated_parsed"], null);
	        this.published = source["published"];
	        this.published_parsed = this.convertValues(source["published_parsed"], null);
	        this.guid = source["guid"];
	        this.image = this.convertValues(source["image"], gofeed.Image);
	        this.categories = source["categories"];
	        this.enclosures = this.convertValues(source["enclosures"], gofeed.Enclosure);
	        this.stupid_enclosure = this.convertValues(source["stupid_enclosure"], gofeed.Enclosure);
	        this.custom = source["custom"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace remote {
	
	export class Device {
	    name?: string;
	    hostname?: string;
	    port?: number;
	    type?: number;
	
	    static createFrom(source: any = {}) {
	        return new Device(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.hostname = source["hostname"];
	        this.port = source["port"];
	        this.type = source["type"];
	    }
	}

}

export namespace server_repo {
	
	export class SupportedGames {
	    lbp1: boolean;
	    lbp2: boolean;
	    lbp3: boolean;
	    lbp_vita: boolean;
	    lbp_psp: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SupportedGames(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.lbp1 = source["lbp1"];
	        this.lbp2 = source["lbp2"];
	        this.lbp3 = source["lbp3"];
	        this.lbp_vita = source["lbp_vita"];
	        this.lbp_psp = source["lbp_psp"];
	    }
	}
	export class Server {
	    name: string;
	    website: string;
	    game_url: string;
	    api_url?: string;
	    feed_url?: string;
	    open_auth: boolean;
	    icon_url?: string;
	    mods_allowed: boolean;
	    display_name: string;
	    description: string;
	    owner: string;
	    accent_colour: string;
	    background_image?: string;
	    supported_games?: SupportedGames;
	
	    static createFrom(source: any = {}) {
	        return new Server(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.website = source["website"];
	        this.game_url = source["game_url"];
	        this.api_url = source["api_url"];
	        this.feed_url = source["feed_url"];
	        this.open_auth = source["open_auth"];
	        this.icon_url = source["icon_url"];
	        this.mods_allowed = source["mods_allowed"];
	        this.display_name = source["display_name"];
	        this.description = source["description"];
	        this.owner = source["owner"];
	        this.accent_colour = source["accent_colour"];
	        this.background_image = source["background_image"];
	        this.supported_games = this.convertValues(source["supported_games"], SupportedGames);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

