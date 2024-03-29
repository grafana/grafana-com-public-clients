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

// checks if the GetStackRegions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStackRegions200Response{}

// GetStackRegions200Response struct for GetStackRegions200Response
type GetStackRegions200Response struct {
	Items                []FormattedApiStackRegion `json:"items"`
	OrderBy              string                    `json:"orderBy"`
	Direction            string                    `json:"direction"`
	Total                float32                   `json:"total"`
	Pages                float32                   `json:"pages"`
	PageSize             float32                   `json:"pageSize"`
	Page                 float32                   `json:"page"`
	AdditionalProperties map[string]interface{}
}

type _GetStackRegions200Response GetStackRegions200Response

// NewGetStackRegions200Response instantiates a new GetStackRegions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStackRegions200Response(items []FormattedApiStackRegion, orderBy string, direction string, total float32, pages float32, pageSize float32, page float32) *GetStackRegions200Response {
	this := GetStackRegions200Response{}
	this.Items = items
	this.OrderBy = orderBy
	this.Direction = direction
	this.Total = total
	this.Pages = pages
	this.PageSize = pageSize
	this.Page = page
	return &this
}

// NewGetStackRegions200ResponseWithDefaults instantiates a new GetStackRegions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStackRegions200ResponseWithDefaults() *GetStackRegions200Response {
	this := GetStackRegions200Response{}
	return &this
}

// GetItems returns the Items field value
func (o *GetStackRegions200Response) GetItems() []FormattedApiStackRegion {
	if o == nil {
		var ret []FormattedApiStackRegion
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *GetStackRegions200Response) GetItemsOk() ([]FormattedApiStackRegion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *GetStackRegions200Response) SetItems(v []FormattedApiStackRegion) {
	o.Items = v
}

// GetOrderBy returns the OrderBy field value
func (o *GetStackRegions200Response) GetOrderBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value
// and a boolean to check if the value has been set.
func (o *GetStackRegions200Response) GetOrderByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderBy, true
}

// SetOrderBy sets field value
func (o *GetStackRegions200Response) SetOrderBy(v string) {
	o.OrderBy = v
}

// GetDirection returns the Direction field value
func (o *GetStackRegions200Response) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *GetStackRegions200Response) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *GetStackRegions200Response) SetDirection(v string) {
	o.Direction = v
}

// GetTotal returns the Total field value
func (o *GetStackRegions200Response) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetStackRegions200Response) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetStackRegions200Response) SetTotal(v float32) {
	o.Total = v
}

// GetPages returns the Pages field value
func (o *GetStackRegions200Response) GetPages() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value
// and a boolean to check if the value has been set.
func (o *GetStackRegions200Response) GetPagesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pages, true
}

// SetPages sets field value
func (o *GetStackRegions200Response) SetPages(v float32) {
	o.Pages = v
}

// GetPageSize returns the PageSize field value
func (o *GetStackRegions200Response) GetPageSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *GetStackRegions200Response) GetPageSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *GetStackRegions200Response) SetPageSize(v float32) {
	o.PageSize = v
}

// GetPage returns the Page field value
func (o *GetStackRegions200Response) GetPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *GetStackRegions200Response) GetPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *GetStackRegions200Response) SetPage(v float32) {
	o.Page = v
}

func (o GetStackRegions200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStackRegions200Response) ToMap() (map[string]interface{}, error) {
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

func (o *GetStackRegions200Response) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varGetStackRegions200Response := _GetStackRegions200Response{}

	err = json.Unmarshal(data, &varGetStackRegions200Response)

	if err != nil {
		return err
	}

	*o = GetStackRegions200Response(varGetStackRegions200Response)

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

type NullableGetStackRegions200Response struct {
	value *GetStackRegions200Response
	isSet bool
}

func (v NullableGetStackRegions200Response) Get() *GetStackRegions200Response {
	return v.value
}

func (v *NullableGetStackRegions200Response) Set(val *GetStackRegions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStackRegions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStackRegions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStackRegions200Response(val *GetStackRegions200Response) *NullableGetStackRegions200Response {
	return &NullableGetStackRegions200Response{value: val, isSet: true}
}

func (v NullableGetStackRegions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStackRegions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
