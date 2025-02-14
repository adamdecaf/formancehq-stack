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


class Response(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        
        class properties:
            data = schemas.DictSchema
            
            
            class cursor(
                schemas.DictSchema
            ):
            
            
                class MetaOapg:
                    
                    class properties:
                        
                        
                        class pageSize(
                            schemas.Int64Schema
                        ):
                            pass
                        hasMore = schemas.BoolSchema
                        
                        
                        class total(
                            schemas.DictSchema
                        ):
                        
                        
                            class MetaOapg:
                                
                                class properties:
                                    
                                    
                                    class value(
                                        schemas.Int64Schema
                                    ):
                                        pass
                                    relation = schemas.StrSchema
                                    __annotations__ = {
                                        "value": value,
                                        "relation": relation,
                                    }
                            
                            @typing.overload
                            def __getitem__(self, name: typing_extensions.Literal["value"]) -> MetaOapg.properties.value: ...
                            
                            @typing.overload
                            def __getitem__(self, name: typing_extensions.Literal["relation"]) -> MetaOapg.properties.relation: ...
                            
                            @typing.overload
                            def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
                            
                            def __getitem__(self, name: typing.Union[typing_extensions.Literal["value", "relation", ], str]):
                                # dict_instance[name] accessor
                                return super().__getitem__(name)
                            
                            
                            @typing.overload
                            def get_item_oapg(self, name: typing_extensions.Literal["value"]) -> typing.Union[MetaOapg.properties.value, schemas.Unset]: ...
                            
                            @typing.overload
                            def get_item_oapg(self, name: typing_extensions.Literal["relation"]) -> typing.Union[MetaOapg.properties.relation, schemas.Unset]: ...
                            
                            @typing.overload
                            def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
                            
                            def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["value", "relation", ], str]):
                                return super().get_item_oapg(name)
                            
                        
                            def __new__(
                                cls,
                                *_args: typing.Union[dict, frozendict.frozendict, ],
                                value: typing.Union[MetaOapg.properties.value, decimal.Decimal, int, schemas.Unset] = schemas.unset,
                                relation: typing.Union[MetaOapg.properties.relation, str, schemas.Unset] = schemas.unset,
                                _configuration: typing.Optional[schemas.Configuration] = None,
                                **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
                            ) -> 'total':
                                return super().__new__(
                                    cls,
                                    *_args,
                                    value=value,
                                    relation=relation,
                                    _configuration=_configuration,
                                    **kwargs,
                                )
                        next = schemas.StrSchema
                        previous = schemas.StrSchema
                        
                        
                        class data(
                            schemas.ListSchema
                        ):
                        
                        
                            class MetaOapg:
                                
                                
                                class items(
                                    schemas.ComposedSchema,
                                ):
                                
                                
                                    class MetaOapg:
                                        all_of_0 = schemas.DictSchema
                                        
                                        @classmethod
                                        @functools.lru_cache()
                                        def all_of(cls):
                                            # we need this here to make our import statements work
                                            # we must store _composed_schemas in here so the code is only run
                                            # when we invoke this method. If we kept this at the class
                                            # level we would get an error because the class level
                                            # code would be run when this module is imported, and these composed
                                            # classes don't exist yet because their module has not finished
                                            # loading
                                            return [
                                                cls.all_of_0,
                                            ]
                                
                                
                                    def __new__(
                                        cls,
                                        *_args: typing.Union[dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, bool, None, list, tuple, bytes, io.FileIO, io.BufferedReader, ],
                                        _configuration: typing.Optional[schemas.Configuration] = None,
                                        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
                                    ) -> 'items':
                                        return super().__new__(
                                            cls,
                                            *_args,
                                            _configuration=_configuration,
                                            **kwargs,
                                        )
                        
                            def __new__(
                                cls,
                                _arg: typing.Union[typing.Tuple[typing.Union[MetaOapg.items, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, bool, None, list, tuple, bytes, io.FileIO, io.BufferedReader, ]], typing.List[typing.Union[MetaOapg.items, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, bool, None, list, tuple, bytes, io.FileIO, io.BufferedReader, ]]],
                                _configuration: typing.Optional[schemas.Configuration] = None,
                            ) -> 'data':
                                return super().__new__(
                                    cls,
                                    _arg,
                                    _configuration=_configuration,
                                )
                        
                            def __getitem__(self, i: int) -> MetaOapg.items:
                                return super().__getitem__(i)
                        __annotations__ = {
                            "pageSize": pageSize,
                            "hasMore": hasMore,
                            "total": total,
                            "next": next,
                            "previous": previous,
                            "data": data,
                        }
                
                @typing.overload
                def __getitem__(self, name: typing_extensions.Literal["pageSize"]) -> MetaOapg.properties.pageSize: ...
                
                @typing.overload
                def __getitem__(self, name: typing_extensions.Literal["hasMore"]) -> MetaOapg.properties.hasMore: ...
                
                @typing.overload
                def __getitem__(self, name: typing_extensions.Literal["total"]) -> MetaOapg.properties.total: ...
                
                @typing.overload
                def __getitem__(self, name: typing_extensions.Literal["next"]) -> MetaOapg.properties.next: ...
                
                @typing.overload
                def __getitem__(self, name: typing_extensions.Literal["previous"]) -> MetaOapg.properties.previous: ...
                
                @typing.overload
                def __getitem__(self, name: typing_extensions.Literal["data"]) -> MetaOapg.properties.data: ...
                
                @typing.overload
                def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
                
                def __getitem__(self, name: typing.Union[typing_extensions.Literal["pageSize", "hasMore", "total", "next", "previous", "data", ], str]):
                    # dict_instance[name] accessor
                    return super().__getitem__(name)
                
                
                @typing.overload
                def get_item_oapg(self, name: typing_extensions.Literal["pageSize"]) -> typing.Union[MetaOapg.properties.pageSize, schemas.Unset]: ...
                
                @typing.overload
                def get_item_oapg(self, name: typing_extensions.Literal["hasMore"]) -> typing.Union[MetaOapg.properties.hasMore, schemas.Unset]: ...
                
                @typing.overload
                def get_item_oapg(self, name: typing_extensions.Literal["total"]) -> typing.Union[MetaOapg.properties.total, schemas.Unset]: ...
                
                @typing.overload
                def get_item_oapg(self, name: typing_extensions.Literal["next"]) -> typing.Union[MetaOapg.properties.next, schemas.Unset]: ...
                
                @typing.overload
                def get_item_oapg(self, name: typing_extensions.Literal["previous"]) -> typing.Union[MetaOapg.properties.previous, schemas.Unset]: ...
                
                @typing.overload
                def get_item_oapg(self, name: typing_extensions.Literal["data"]) -> typing.Union[MetaOapg.properties.data, schemas.Unset]: ...
                
                @typing.overload
                def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
                
                def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["pageSize", "hasMore", "total", "next", "previous", "data", ], str]):
                    return super().get_item_oapg(name)
                
            
                def __new__(
                    cls,
                    *_args: typing.Union[dict, frozendict.frozendict, ],
                    pageSize: typing.Union[MetaOapg.properties.pageSize, decimal.Decimal, int, schemas.Unset] = schemas.unset,
                    hasMore: typing.Union[MetaOapg.properties.hasMore, bool, schemas.Unset] = schemas.unset,
                    total: typing.Union[MetaOapg.properties.total, dict, frozendict.frozendict, schemas.Unset] = schemas.unset,
                    next: typing.Union[MetaOapg.properties.next, str, schemas.Unset] = schemas.unset,
                    previous: typing.Union[MetaOapg.properties.previous, str, schemas.Unset] = schemas.unset,
                    data: typing.Union[MetaOapg.properties.data, list, tuple, schemas.Unset] = schemas.unset,
                    _configuration: typing.Optional[schemas.Configuration] = None,
                    **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
                ) -> 'cursor':
                    return super().__new__(
                        cls,
                        *_args,
                        pageSize=pageSize,
                        hasMore=hasMore,
                        total=total,
                        next=next,
                        previous=previous,
                        data=data,
                        _configuration=_configuration,
                        **kwargs,
                    )
            __annotations__ = {
                "data": data,
                "cursor": cursor,
            }
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["data"]) -> MetaOapg.properties.data: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["cursor"]) -> MetaOapg.properties.cursor: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["data", "cursor", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["data"]) -> typing.Union[MetaOapg.properties.data, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["cursor"]) -> typing.Union[MetaOapg.properties.cursor, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["data", "cursor", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *_args: typing.Union[dict, frozendict.frozendict, ],
        data: typing.Union[MetaOapg.properties.data, dict, frozendict.frozendict, schemas.Unset] = schemas.unset,
        cursor: typing.Union[MetaOapg.properties.cursor, dict, frozendict.frozendict, schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'Response':
        return super().__new__(
            cls,
            *_args,
            data=data,
            cursor=cursor,
            _configuration=_configuration,
            **kwargs,
        )
