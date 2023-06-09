// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {remote} from '../models';
import {server_repo} from '../models';

export function AddDevice(arg1:string,arg2:string,arg3:number,arg4:remote.DeviceType):Promise<void>;

export function GetServer(arg1:string):Promise<void>;

export function GetServers():Promise<{[key: string]: server_repo.Server}>;

export function PatchDevice(arg1:string,arg2:string):Promise<void>;

export function PatchFile(arg1:string,arg2:string,arg3:string):Promise<void>;
