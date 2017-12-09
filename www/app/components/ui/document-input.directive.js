'use strict';

angular
.module('App')
.directive('documentInput', documentInput);

function documentInput() {
    return {
        templateUrl: 'app/components/ui/document_input.html'
    };
};
