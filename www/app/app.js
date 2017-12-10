(function (angular) {
    'use strict';

// Declare app level module which depends on views, and components
angular
    .module('App', ['angularMask'])
    .controller('AppController', AppController);

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

    function blacklist(id, isBlackListed) {
        return documentsService.blacklist(id, !isBlackListed);
    }

    function del(id) {
        return documentsService.del(id)
            .then(()=> get());
    }

    function create() {
        const document = {id: vm.input};
        return documentsService.validate(document.id)
            .then(() => documentsService.create(document))
            .then(() => get());
    }

    get();
}

})(angular);
