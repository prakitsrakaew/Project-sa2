/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    ControllersBill,
    ControllersBillFromJSON,
    ControllersBillToJSON,
    EntBill,
    EntBillFromJSON,
    EntBillToJSON,
    EntBillingstatus,
    EntBillingstatusFromJSON,
    EntBillingstatusToJSON,
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeToJSON,
    EntRepairinvoice,
    EntRepairinvoiceFromJSON,
    EntRepairinvoiceToJSON,
} from '../models';

export interface CreateBillRequest {
    bill: ControllersBill;
}

export interface CreateBillingstatusRequest {
    billingstatus: EntBillingstatus;
}

export interface CreateEmployeeRequest {
    employee: EntEmployee;
}

export interface CreateRepairinvoiceRequest {
    repairinvoice: EntRepairinvoice;
}

export interface GetBillingstatusRequest {
    id: number;
}

export interface GetEmployeeRequest {
    id: number;
}

export interface GetRepairinvoiceRequest {
    id: number;
}

export interface ListBillRequest {
    limit?: number;
    offset?: number;
}

export interface ListBillingstatusRequest {
    limit?: number;
    offset?: number;
}

export interface ListEmployeeRequest {
    limit?: number;
    offset?: number;
}

export interface ListRepairinvoiceRequest {
    limit?: number;
    offset?: number;
}

/**
 * 
 */
export class DefaultApi extends runtime.BaseAPI {

    /**
     * Create bill
     * Create bill
     */
    async createBillRaw(requestParameters: CreateBillRequest): Promise<runtime.ApiResponse<ControllersBill>> {
        if (requestParameters.bill === null || requestParameters.bill === undefined) {
            throw new runtime.RequiredError('bill','Required parameter requestParameters.bill was null or undefined when calling createBill.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/bills`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ControllersBillToJSON(requestParameters.bill),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ControllersBillFromJSON(jsonValue));
    }

    /**
     * Create bill
     * Create bill
     */
    async createBill(requestParameters: CreateBillRequest): Promise<ControllersBill> {
        const response = await this.createBillRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create billingstatus
     * Create billingstatus
     */
    async createBillingstatusRaw(requestParameters: CreateBillingstatusRequest): Promise<runtime.ApiResponse<EntBillingstatus>> {
        if (requestParameters.billingstatus === null || requestParameters.billingstatus === undefined) {
            throw new runtime.RequiredError('billingstatus','Required parameter requestParameters.billingstatus was null or undefined when calling createBillingstatus.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/billingstatuss`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntBillingstatusToJSON(requestParameters.billingstatus),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntBillingstatusFromJSON(jsonValue));
    }

    /**
     * Create billingstatus
     * Create billingstatus
     */
    async createBillingstatus(requestParameters: CreateBillingstatusRequest): Promise<EntBillingstatus> {
        const response = await this.createBillingstatusRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create employee
     * Create employee
     */
    async createEmployeeRaw(requestParameters: CreateEmployeeRequest): Promise<runtime.ApiResponse<EntEmployee>> {
        if (requestParameters.employee === null || requestParameters.employee === undefined) {
            throw new runtime.RequiredError('employee','Required parameter requestParameters.employee was null or undefined when calling createEmployee.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/employees`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntEmployeeToJSON(requestParameters.employee),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntEmployeeFromJSON(jsonValue));
    }

    /**
     * Create employee
     * Create employee
     */
    async createEmployee(requestParameters: CreateEmployeeRequest): Promise<EntEmployee> {
        const response = await this.createEmployeeRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create repairinvoice
     * Create repairinvoice
     */
    async createRepairinvoiceRaw(requestParameters: CreateRepairinvoiceRequest): Promise<runtime.ApiResponse<EntRepairinvoice>> {
        if (requestParameters.repairinvoice === null || requestParameters.repairinvoice === undefined) {
            throw new runtime.RequiredError('repairinvoice','Required parameter requestParameters.repairinvoice was null or undefined when calling createRepairinvoice.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/repairinvoices`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntRepairinvoiceToJSON(requestParameters.repairinvoice),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntRepairinvoiceFromJSON(jsonValue));
    }

    /**
     * Create repairinvoice
     * Create repairinvoice
     */
    async createRepairinvoice(requestParameters: CreateRepairinvoiceRequest): Promise<EntRepairinvoice> {
        const response = await this.createRepairinvoiceRaw(requestParameters);
        return await response.value();
    }

    /**
     * get billingstatus by ID
     * Get a billingstatus entity by ID
     */
    async getBillingstatusRaw(requestParameters: GetBillingstatusRequest): Promise<runtime.ApiResponse<EntBillingstatus>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getBillingstatus.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/billingstatuss/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntBillingstatusFromJSON(jsonValue));
    }

    /**
     * get billingstatus by ID
     * Get a billingstatus entity by ID
     */
    async getBillingstatus(requestParameters: GetBillingstatusRequest): Promise<EntBillingstatus> {
        const response = await this.getBillingstatusRaw(requestParameters);
        return await response.value();
    }

    /**
     * get employee by ID
     * Get a employee entity by ID
     */
    async getEmployeeRaw(requestParameters: GetEmployeeRequest): Promise<runtime.ApiResponse<EntEmployee>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getEmployee.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/employees/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntEmployeeFromJSON(jsonValue));
    }

    /**
     * get employee by ID
     * Get a employee entity by ID
     */
    async getEmployee(requestParameters: GetEmployeeRequest): Promise<EntEmployee> {
        const response = await this.getEmployeeRaw(requestParameters);
        return await response.value();
    }

    /**
     * get repairinvoice by ID
     * Get a repairinvoice entity by ID
     */
    async getRepairinvoiceRaw(requestParameters: GetRepairinvoiceRequest): Promise<runtime.ApiResponse<EntRepairinvoice>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getRepairinvoice.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/repairinvoices/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntRepairinvoiceFromJSON(jsonValue));
    }

    /**
     * get repairinvoice by ID
     * Get a repairinvoice entity by ID
     */
    async getRepairinvoice(requestParameters: GetRepairinvoiceRequest): Promise<EntRepairinvoice> {
        const response = await this.getRepairinvoiceRaw(requestParameters);
        return await response.value();
    }

    /**
     * list bill entities
     * List bill entities
     */
    async listBillRaw(requestParameters: ListBillRequest): Promise<runtime.ApiResponse<Array<EntBill>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/bills`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntBillFromJSON));
    }

    /**
     * list bill entities
     * List bill entities
     */
    async listBill(requestParameters: ListBillRequest): Promise<Array<EntBill>> {
        const response = await this.listBillRaw(requestParameters);
        return await response.value();
    }

    /**
     * list billingstatus entities
     * List billingstatus entities
     */
    async listBillingstatusRaw(requestParameters: ListBillingstatusRequest): Promise<runtime.ApiResponse<Array<EntBillingstatus>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/billingstatuss`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntBillingstatusFromJSON));
    }

    /**
     * list billingstatus entities
     * List billingstatus entities
     */
    async listBillingstatus(requestParameters: ListBillingstatusRequest): Promise<Array<EntBillingstatus>> {
        const response = await this.listBillingstatusRaw(requestParameters);
        return await response.value();
    }

    /**
     * list employee entities
     * List employee entities
     */
    async listEmployeeRaw(requestParameters: ListEmployeeRequest): Promise<runtime.ApiResponse<Array<EntEmployee>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/employees`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntEmployeeFromJSON));
    }

    /**
     * list employee entities
     * List employee entities
     */
    async listEmployee(requestParameters: ListEmployeeRequest): Promise<Array<EntEmployee>> {
        const response = await this.listEmployeeRaw(requestParameters);
        return await response.value();
    }

    /**
     * list repairinvoice entities
     * List repairinvoice entities
     */
    async listRepairinvoiceRaw(requestParameters: ListRepairinvoiceRequest): Promise<runtime.ApiResponse<Array<EntRepairinvoice>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/repairinvoices`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntRepairinvoiceFromJSON));
    }

    /**
     * list repairinvoice entities
     * List repairinvoice entities
     */
    async listRepairinvoice(requestParameters: ListRepairinvoiceRequest): Promise<Array<EntRepairinvoice>> {
        const response = await this.listRepairinvoiceRaw(requestParameters);
        return await response.value();
    }

}
