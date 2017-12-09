'use strict';

angular
.module('App')
.directive('documentList', documentList);

function documentList() {
    return {
        templateUrl: 'app/components/ui/document_list.html'
    };
};
