/*
GCOM API

Grafana.com API (public).  Looking for GCOM API client packages? You can find them at [grafana-com-public-clients](https://github.com/grafana/grafana-com-public-clients) repository.  If you have any questions, please contact support in the Grafana Cloud UI.  This spec is in *Beta* stage, so use it with caution: - Not all endpoint responses are properly typed for the time being. - Some request parameter types may not be precise

API version: public
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gcom

import (
	"encoding/json"
)

// checks if the GetInstancePlugins200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInstancePlugins200Response{}

// GetInstancePlugins200Response struct for GetInstancePlugins200Response
type GetInstancePlugins200Response struct {
	Items                []FormattedApiInstancePlugin `json:"items"`
	OrderBy              string                       `json:"orderBy"`
	Direction            string                       `json:"direction"`
	Total                float32                      `json:"total"`
	Pages                float32                      `json:"pages"`
	PageSize             float32                      `json:"pageSize"`
	Page                 float32                      `json:"page"`
	AdditionalProperties map[string]interface{}
}

type _GetInstancePlugins200Response GetInstancePlugins200Response

// NewGetInstancePlugins200Response instantiates a new GetInstancePlugins200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInstancePlugins200Response(items []FormattedApiInstancePlugin, orderBy string, direction string, total float32, pages float32, pageSize float32, page float32) *GetInstancePlugins200Response {
	this := GetInstancePlugins200Response{}
	this.Items = items
	this.OrderBy = orderBy
	this.Direction = direction
	this.Total = total
	this.Pages = pages
	this.PageSize = pageSize
	this.Page = page
	return &this
}

// NewGetInstancePlugins200ResponseWithDefaults instantiates a new GetInstancePlugins200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInstancePlugins200ResponseWithDefaults() *GetInstancePlugins200Response {
	this := GetInstancePlugins200Response{}
	return &this
}

// GetItems returns the Items field value
func (o *GetInstancePlugins200Response) GetItems() []FormattedApiInstancePlugin {
	if o == nil {
		var ret []FormattedApiInstancePlugin
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *GetInstancePlugins200Response) GetItemsOk() ([]FormattedApiInstancePlugin, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *GetInstancePlugins200Response) SetItems(v []FormattedApiInstancePlugin) {
	o.Items = v
}

// GetOrderBy returns the OrderBy field value
func (o *GetInstancePlugins200Response) GetOrderBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value
// and a boolean to check if the value has been set.
func (o *GetInstancePlugins200Response) GetOrderByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderBy, true
}

// SetOrderBy sets field value
func (o *GetInstancePlugins200Response) SetOrderBy(v string) {
	o.OrderBy = v
}

// GetDirection returns the Direction field value
func (o *GetInstancePlugins200Response) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *GetInstancePlugins200Response) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *GetInstancePlugins200Response) SetDirection(v string) {
	o.Direction = v
}

// GetTotal returns the Total field value
func (o *GetInstancePlugins200Response) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetInstancePlugins200Response) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetInstancePlugins200Response) SetTotal(v float32) {
	o.Total = v
}

// GetPages returns the Pages field value
func (o *GetInstancePlugins200Response) GetPages() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value
// and a boolean to check if the value has been set.
func (o *GetInstancePlugins200Response) GetPagesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pages, true
}

// SetPages sets field value
func (o *GetInstancePlugins200Response) SetPages(v float32) {
	o.Pages = v
}

// GetPageSize returns the PageSize field value
func (o *GetInstancePlugins200Response) GetPageSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *GetInstancePlugins200Response) GetPageSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *GetInstancePlugins200Response) SetPageSize(v float32) {
	o.PageSize = v
}

// GetPage returns the Page field value
func (o *GetInstancePlugins200Response) GetPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *GetInstancePlugins200Response) GetPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *GetInstancePlugins200Response) SetPage(v float32) {
	o.Page = v
}

func (o GetInstancePlugins200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInstancePlugins200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["orderBy"] = o.OrderBy
	toSerialize["direction"] = o.Direction
	toSerialize["total"] = o.Total
	toSerialize["pages"] = o.Pages
	toSerialize["pageSize"] = o.PageSize
	toSerialize["page"] = o.Page

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetInstancePlugins200Response) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varGetInstancePlugins200Response := _GetInstancePlugins200Response{}

	err = json.Unmarshal(data, &varGetInstancePlugins200Response)

	if err != nil {
		return err
	}

	*o = GetInstancePlugins200Response(varGetInstancePlugins200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "orderBy")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "total")
		delete(additionalProperties, "pages")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "page")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetInstancePlugins200Response struct {
	value *GetInstancePlugins200Response
	isSet bool
}

func (v NullableGetInstancePlugins200Response) Get() *GetInstancePlugins200Response {
	return v.value
}

func (v *NullableGetInstancePlugins200Response) Set(val *GetInstancePlugins200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInstancePlugins200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInstancePlugins200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInstancePlugins200Response(val *GetInstancePlugins200Response) *NullableGetInstancePlugins200Response {
	return &NullableGetInstancePlugins200Response{value: val, isSet: true}
}

func (v NullableGetInstancePlugins200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInstancePlugins200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
