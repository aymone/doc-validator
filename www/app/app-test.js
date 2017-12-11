'use strict';

describe('Controller: AppController', function() {
    beforeEach(angular.mock.module('App'));

    let $controller, documentsService, $mdToast;

    beforeEach(function(){
        inject(function(_$controller_, _documentsService_, _$mdToast_) {
            $controller = _$controller_;
            documentsService = _documentsService_;
            $mdToast = _$mdToast_;
        });
    });

    it('should have some props defined', function() {
        const AppController = $controller('AppController', {documentsService, $mdToast});
        expect(AppController).toBeDefined();
        expect(AppController.valid).toBeDefined();
        expect(AppController.input).toBeDefined();
        expect(AppController.sorter).toBeDefined();
        expect(AppController.documents).toBeDefined();
    });
});
