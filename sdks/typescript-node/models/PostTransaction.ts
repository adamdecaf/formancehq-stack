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

import { PostTransactionScript } from '../models/PostTransactionScript';
import { Posting } from '../models/Posting';
import { HttpFile } from '../http/http';

export class PostTransaction {
    'timestamp'?: Date;
    'postings'?: Array<Posting>;
    'script'?: PostTransactionScript;
    'reference'?: string;
    'metadata'?: { [key: string]: any; };

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "timestamp",
            "baseName": "timestamp",
            "type": "Date",
            "format": "date-time"
        },
        {
            "name": "postings",
            "baseName": "postings",
            "type": "Array<Posting>",
            "format": ""
        },
        {
            "name": "script",
            "baseName": "script",
            "type": "PostTransactionScript",
            "format": ""
        },
        {
            "name": "reference",
            "baseName": "reference",
            "type": "string",
            "format": ""
        },
        {
            "name": "metadata",
            "baseName": "metadata",
            "type": "{ [key: string]: any; }",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return PostTransaction.attributeTypeMap;
    }

    public constructor() {
    }
}

