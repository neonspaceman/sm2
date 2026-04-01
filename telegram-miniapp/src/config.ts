export enum APP_ENV {
  DEV = "dev",
  PROD = "prod",
  TEST = "test"
}

export const CONFIG = {
  API_BASE_URL: import.meta.env.VITE_API_BASE_URL,
  APP_ENV: import.meta.env.VITE_APP_ENV,
};
