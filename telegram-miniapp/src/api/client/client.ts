import axios from "axios";
import {CONFIG} from "@/config.ts";

export const client = axios.create({
    baseURL: CONFIG.API_BASE_URL,
})