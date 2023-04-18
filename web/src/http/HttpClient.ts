import axios from 'axios';

axios.defaults.baseURL = '/api/v1/';
export default axios.create();
