# coding: utf-8

"""
    Formance Stack API

    Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions />   # noqa: E501

    The version of the OpenAPI document: develop
    Contact: support@formance.com
    Generated by: https://openapi-generator.tech
"""

from datetime import date, datetime  # noqa: F401
import decimal  # noqa: F401
import functools  # noqa: F401
import io  # noqa: F401
import re  # noqa: F401
import typing  # noqa: F401
import typing_extensions  # noqa: F401
import uuid  # noqa: F401

import frozendict  # noqa: F401

from Formance import schemas  # noqa: F401


class PaymentAdjustment(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        required = {
            "date",
            "amount",
            "absolute",
            "raw",
            "status",
        }
        
        class properties:
        
            @staticmethod
            def status() -> typing.Type['PaymentStatus']:
                return PaymentStatus
            
            
            class amount(
                schemas.Int64Schema
            ):
                pass
            date = schemas.DateTimeSchema
            raw = schemas.DictSchema
            absolute = schemas.BoolSchema
            __annotations__ = {
                "status": status,
                "amount": amount,
                "date": date,
                "raw": raw,
                "absolute": absolute,
            }
    
    date: MetaOapg.properties.date
    amount: MetaOapg.properties.amount
    absolute: MetaOapg.properties.absolute
    raw: MetaOapg.properties.raw
    status: 'PaymentStatus'
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["status"]) -> 'PaymentStatus': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["amount"]) -> MetaOapg.properties.amount: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["date"]) -> MetaOapg.properties.date: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["raw"]) -> MetaOapg.properties.raw: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["absolute"]) -> MetaOapg.properties.absolute: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["status", "amount", "date", "raw", "absolute", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["status"]) -> 'PaymentStatus': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["amount"]) -> MetaOapg.properties.amount: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["date"]) -> MetaOapg.properties.date: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["raw"]) -> MetaOapg.properties.raw: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["absolute"]) -> MetaOapg.properties.absolute: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["status", "amount", "date", "raw", "absolute", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *_args: typing.Union[dict, frozendict.frozendict, ],
        date: typing.Union[MetaOapg.properties.date, str, datetime, ],
        amount: typing.Union[MetaOapg.properties.amount, decimal.Decimal, int, ],
        absolute: typing.Union[MetaOapg.properties.absolute, bool, ],
        raw: typing.Union[MetaOapg.properties.raw, dict, frozendict.frozendict, ],
        status: 'PaymentStatus',
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'PaymentAdjustment':
        return super().__new__(
            cls,
            *_args,
            date=date,
            amount=amount,
            absolute=absolute,
            raw=raw,
            status=status,
            _configuration=_configuration,
            **kwargs,
        )

from Formance.model.payment_status import PaymentStatus
