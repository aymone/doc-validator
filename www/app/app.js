'use strict';

// Declare app level module which depends on views, and components
angular
  .module('App', ['angularMask'])
  .controller('AppController', AppController)
  .factory('documentsService', documentsService);

AppController.$inject = ['documentsService'];
function AppController(documentsService) {
  // bind scope
  let vm = this;

  // bind variables
  vm.valid = false;
  vm.input = "";
  vm.sorter = "";
  vm.documents = [];

  // bind methods
  vm.get = get;
  vm.del = del;
  vm.create = create;
  vm.blacklist = blacklist;
  vm.cleanFilter = cleanFilter;
  vm.setSorter = setSorter;

  function cleanFilter() {
    vm.input = ""
    vm.sorter = ""
    get();
  }

  function get() {
    const query = {};
    if (vm.input) {
      query.filter = vm.input
    }

    if (vm.sorter) {
      query.sorter = vm.sorter
    }

    return documentsService.validate(vm.input)
      .then(response => {
          vm.valid = response.data.isValid;
      })
      .catch(() => {
        vm.valid = false
      })
      .then(() => documentsService.getAll(query))
      .then(response => {
        vm.documents = response.data;
      });
  }

  function setSorter(sorter) {
    if (vm.sorter === sorter) {
      sorter = `-${sorter}`;
    }

    vm.sorter = sorter;
    get();
  }

  function blacklist(documentNumber, isBlackListed) {
    return documentsService.blacklist(documentNumber, !isBlackListed);
  }

  function del(documentNumber) {
    return documentsService.del(documentNumber)
      .then(()=> get());
  }

  function create() {
    const document = {documentNumber: vm.input};
    return documentsService.validate(document.documentNumber)
      .then(() => documentsService.create(document))
      .then(() => get());
  }

  get();
}

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
    validate(documentNumber) {
      return $http.get(`${host}/documents/${documentNumber}/validate`, document);
    },
    create(document) {
      return $http.post(`${host}/documents`, document);
    },
    del(documentNumber) {
      return $http.delete(`${host}/documents/${documentNumber}`);
    },
    blacklist(documentNumber, isBlackListed) {
      const status = isBlackListed ? 'remove' : 'add';
      return $http.put(`${host}/documents/${documentNumber}/blacklist/${status}`, document);
    }
  }
}
