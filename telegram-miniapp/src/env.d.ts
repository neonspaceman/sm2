/// <reference types="vite/client" />

import {APP_ENV} from "@/config.ts";

interface ImportMetaEnv {
  readonly VITE_API_BASE_URL: string;
  readonly VITE_APP_ENV: APP_ENV
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
