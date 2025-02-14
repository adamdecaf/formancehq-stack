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

import { BankingCircleConfig } from '../models/BankingCircleConfig';
import { CurrencyCloudConfig } from '../models/CurrencyCloudConfig';
import { DummyPayConfig } from '../models/DummyPayConfig';
import { ModulrConfig } from '../models/ModulrConfig';
import { StripeConfig } from '../models/StripeConfig';
import { WiseConfig } from '../models/WiseConfig';
import { HttpFile } from '../http/http';

export class ConnectorConfig {
    /**
    * The frequency at which the connector will fetch transactions
    */
    'pollingPeriod'?: string;
    'apiKey': string;
    /**
    * Number of BalanceTransaction to fetch at each polling interval. 
    */
    'pageSize'?: number;
    /**
    * The frequency at which the connector will try to fetch new payment objects from the directory
    */
    'filePollingPeriod'?: string;
    /**
    * The frequency at which the connector will create new payment objects in the directory
    */
    'fileGenerationPeriod'?: string;
    'directory': string;
    'apiSecret': string;
    'endpoint': string;
    /**
    * Username of the API Key holder
    */
    'loginID': string;
    'username': string;
    'password': string;
    'authorizationEndpoint': string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "pollingPeriod",
            "baseName": "pollingPeriod",
            "type": "string",
            "format": ""
        },
        {
            "name": "apiKey",
            "baseName": "apiKey",
            "type": "string",
            "format": ""
        },
        {
            "name": "pageSize",
            "baseName": "pageSize",
            "type": "number",
            "format": "int64"
        },
        {
            "name": "filePollingPeriod",
            "baseName": "filePollingPeriod",
            "type": "string",
            "format": ""
        },
        {
            "name": "fileGenerationPeriod",
            "baseName": "fileGenerationPeriod",
            "type": "string",
            "format": ""
        },
        {
            "name": "directory",
            "baseName": "directory",
            "type": "string",
            "format": ""
        },
        {
            "name": "apiSecret",
            "baseName": "apiSecret",
            "type": "string",
            "format": ""
        },
        {
            "name": "endpoint",
            "baseName": "endpoint",
            "type": "string",
            "format": ""
        },
        {
            "name": "loginID",
            "baseName": "loginID",
            "type": "string",
            "format": ""
        },
        {
            "name": "username",
            "baseName": "username",
            "type": "string",
            "format": ""
        },
        {
            "name": "password",
            "baseName": "password",
            "type": "string",
            "format": ""
        },
        {
            "name": "authorizationEndpoint",
            "baseName": "authorizationEndpoint",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return ConnectorConfig.attributeTypeMap;
    }

    public constructor() {
    }
}

