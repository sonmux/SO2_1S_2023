export namespace sys {
	
	export class CPUUsage {
	    avg: number;
	
	    static createFrom(source: any = {}) {
	        return new CPUUsage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.avg = source["avg"];
	    }
	}
	export class DISKUsage {
	    used: number;
	    free: number;
	
	    static createFrom(source: any = {}) {
	        return new DISKUsage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.used = source["used"];
	        this.free = source["free"];
	    }
	}

}
