/**
 * Formance Stack API
 * Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 
 *
 * OpenAPI spec version: develop
 * Contact: support@formance.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { HttpFile } from '../http/http';

export class StageStatus {
    'stage': number;
    'instanceID': string;
    'startedAt': Date;
    'terminatedAt'?: Date;
    'error'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "stage",
            "baseName": "stage",
            "type": "number",
            "format": ""
        },
        {
            "name": "instanceID",
            "baseName": "instanceID",
            "type": "string",
            "format": ""
        },
        {
            "name": "startedAt",
            "baseName": "startedAt",
            "type": "Date",
            "format": "date-time"
        },
        {
            "name": "terminatedAt",
            "baseName": "terminatedAt",
            "type": "Date",
            "format": "date-time"
        },
        {
            "name": "error",
            "baseName": "error",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return StageStatus.attributeTypeMap;
    }

    public constructor() {
    }
}

