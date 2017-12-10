(function (angular) {
    'use strict';

angular
    .module('App')
    .factory('documentsService', documentsService);

documentsService.$inject = ['$http'];
function documentsService($http) {
    const host = 'http://localhost:3000';
    return {
        getAll(query) {
            let endpoint = `${host}/documents`;
            const params = [];
            if (query) {
                if (query.filter) {
                    params.push(`filter=${query.filter}`);
                }
                if (query.sorter) {
                    params.push(`sort=${query.sorter}`);
                }
                if (params.length) {
                    endpoint = `${endpoint}?${params.join('&')}`
                }
            }

            return $http.get(endpoint);
        },
        validate(id) {
            return $http.get(`${host}/documents/${id}/validate`, document);
        },
        create(document) {
            return $http.post(`${host}/documents`, document);
        },
        del(id) {
            return $http.delete(`${host}/documents/${id}`);
        },
        blacklist(id, isBlackListed) {
        const status = isBlackListed ? 'remove' : 'add';
            return $http.put(`${host}/documents/${id}/blacklist/${status}`, document);
        }
    }
}

})(angular);
