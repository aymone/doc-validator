'use strict';

describe('Controller: AppController', function() {
    beforeEach(module('App'));

    let AppController;

    beforeEach(inject(function($controller) {
        let scope = {};
        AppController = $controller('AppController', {});
    }));

    it('should have some props defined', function() {
        expect(AppController).toBeDefined();
        expect(AppController.valid).toBeDefined();
        expect(AppController.input).toBeDefined();
        expect(AppController.sorter).toBeDefined();
        expect(AppController.documents).toBeDefined();
    });
});
