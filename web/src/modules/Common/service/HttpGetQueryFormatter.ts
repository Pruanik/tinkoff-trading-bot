export default {
    getGetQueryString(data: Object) {
        let resultQuery = '';

        Object.keys(data).forEach(key => {
            if (data[key] !== undefined && data[key] !== null) {
                resultQuery += (resultQuery === '' ? '?' : '&') + key + '=' + data[key];
            }
        });

        return resultQuery;
    }
}
