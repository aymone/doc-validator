'use strict';

describe('Service: documentsService', function() {
    let $httpBackend, documentsService, result;
    const host = 'http://localhost:3000';
    const id = '36673263865';
    const sorter = 'id';
    const response = [
        {
            id: "some-id"
        }
    ];

    beforeEach(module('App'));
    beforeEach(inject(function (_documentsService_, _$httpBackend_) {
        $httpBackend = _$httpBackend_;
        documentsService = _documentsService_;
        result = null;
    }));

    it('should get the documents', function () {
        $httpBackend
            .expect('GET', `${host}/documents`)
            .respond(200, response);

        documentsService.getAll().then(res => {
            result = res.data;
        });

        $httpBackend.flush();
        expect(result).toEqual(response);
    });

    it('should get the documents with filter', function () {
        $httpBackend
            .expect('GET', `${host}/documents?filter=${id}`)
            .respond(200, response);

        documentsService.getAll({filter: id}).then(res => {
            result = res.data;
        });

        $httpBackend.flush();
        expect(result).toEqual(response);
    });

    it('should get the documents with sorter', function () {
        $httpBackend
            .expect('GET', `${host}/documents?sort=id`)
            .respond(200, response);

        documentsService.getAll({sorter}).then(res => {
            result = res.data;
        });

        $httpBackend.flush();
        expect(result).toEqual(response);
    });

    it('should get the documents with filter and sorter', function () {
        $httpBackend
            .expect('GET', `${host}/documents?filter=${id}&sort=id`)
            .respond(200, response);

        documentsService.getAll({sorter, filter: id}).then(res => {
            result = res.data;
        });

        $httpBackend.flush();
        expect(result).toEqual(response);
    });

    it('should validate document', function () {
        $httpBackend
            .expect('GET', `${host}/documents/${id}/validate`)
            .respond(200);

        documentsService.validate(id).then(res => {
            result = true;
        });

        $httpBackend.flush();
        expect(result).toEqual(true);
    });

    it('should create document', function () {
        $httpBackend
            .expect('POST', `${host}/documents`)
            .respond(200, response);

        documentsService.create({id}).then(res => {
            result = res.data;
        });

        $httpBackend.flush();
        expect(result).toEqual(response);
    });

    it('should delete document', function () {
        $httpBackend
            .expect('DELETE', `${host}/documents/${id}`)
            .respond(200, response);

        documentsService.del(id).then(res => {
            result = true;
        });

        $httpBackend.flush();
        expect(result).toEqual(true);
    });

    it('should blacklist document', function () {
        $httpBackend
            .expect('PUT', `${host}/documents/${id}/blacklist/add`)
            .respond(200, response);

        documentsService.blacklist(id, false).then(res => {
            result = true;
        });

        $httpBackend.flush();
        expect(result).toEqual(true);
    });

    afterEach(function () {
        $httpBackend.verifyNoOutstandingExpectation();
        $httpBackend.verifyNoOutstandingRequest();
    });
});
