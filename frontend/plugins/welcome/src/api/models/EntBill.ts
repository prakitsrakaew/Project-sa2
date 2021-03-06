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

import { exists, mapValues } from '../runtime';
import {
    EntBillEdges,
    EntBillEdgesFromJSON,
    EntBillEdgesFromJSONTyped,
    EntBillEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBill
 */
export interface EntBill {
    /**
     * 
     * @type {EntBillEdges}
     * @memberof EntBill
     */
    edges?: EntBillEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntBill
     */
    id?: number;
    /**
     * Price holds the value of the "price" field.
     * @type {number}
     * @memberof EntBill
     */
    price?: number;
    /**
     * Time holds the value of the "time" field.
     * @type {number}
     * @memberof EntBill
     */
    time?: number;
}

export function EntBillFromJSON(json: any): EntBill {
    return EntBillFromJSONTyped(json, false);
}

export function EntBillFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBill {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntBillEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'price': !exists(json, 'price') ? undefined : json['price'],
        'time': !exists(json, 'time') ? undefined : json['time'],
    };
}

export function EntBillToJSON(value?: EntBill | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntBillEdgesToJSON(value.edges),
        'id': value.id,
        'price': value.price,
        'time': value.time,
    };
}


