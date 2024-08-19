import Functions from '@/utils/functions';

const funcs = new Functions();
export const API_URL = import.meta.env.VITE_APP_API_URL;
export const HOME_URL = '/dashboard?lang=' + funcs.getLang();
export const LOGIN_URL = '/login?lang=' + funcs.getLang();