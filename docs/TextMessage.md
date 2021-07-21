# TextMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Id** | **string** |  | 
**Text** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewTextMessage

`func NewTextMessage(id string, text string, type_ string, ) *TextMessage`

NewTextMessage instantiates a new TextMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextMessageWithDefaults

`func NewTextMessageWithDefaults() *TextMessage`

NewTextMessageWithDefaults instantiates a new TextMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *TextMessage) GetContext() map[string]map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TextMessage) GetContextOk() (*map[string]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TextMessage) SetContext(v map[string]map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *TextMessage) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *TextMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TextMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TextMessage) SetId(v string)`

SetId sets Id field to given value.


### GetText

`func (o *TextMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TextMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TextMessage) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *TextMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TextMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TextMessage) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


