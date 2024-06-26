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

// checks if the FormattedApiApiKeyListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormattedApiApiKeyListResponse{}

// FormattedApiApiKeyListResponse struct for FormattedApiApiKeyListResponse
type FormattedApiApiKeyListResponse struct {
	Items                []ItemsInner `json:"items"`
	OrderBy              *string      `json:"orderBy,omitempty"`
	Direction            *string      `json:"direction,omitempty"`
	Total                *float32     `json:"total,omitempty"`
	Pages                *float32     `json:"pages,omitempty"`
	PageSize             *float32     `json:"pageSize,omitempty"`
	Page                 *float32     `json:"page,omitempty"`
	Cursor               *float32     `json:"cursor,omitempty"`
	NextCursor           *float32     `json:"nextCursor,omitempty"`
	Links                []LinksInner `json:"links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FormattedApiApiKeyListResponse FormattedApiApiKeyListResponse

// NewFormattedApiApiKeyListResponse instantiates a new FormattedApiApiKeyListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormattedApiApiKeyListResponse(items []ItemsInner) *FormattedApiApiKeyListResponse {
	this := FormattedApiApiKeyListResponse{}
	this.Items = items
	return &this
}

// NewFormattedApiApiKeyListResponseWithDefaults instantiates a new FormattedApiApiKeyListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormattedApiApiKeyListResponseWithDefaults() *FormattedApiApiKeyListResponse {
	this := FormattedApiApiKeyListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *FormattedApiApiKeyListResponse) GetItems() []ItemsInner {
	if o == nil {
		var ret []ItemsInner
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKeyListResponse) GetItemsOk() ([]ItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *FormattedApiApiKeyListResponse) SetItems(v []ItemsInner) {
	o.Items = v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *FormattedApiApiKeyListResponse) GetOrderBy() string {
	if o == nil || IsNil(o.OrderBy) {
		var ret string
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKeyListResponse) GetOrderByOk() (*string, bool) {
	if o == nil || IsNil(o.OrderBy) {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *FormattedApiApiKeyListResponse) HasOrderBy() bool {
	if o != nil && !IsNil(o.OrderBy) {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given string and assigns it to the OrderBy field.
func (o *FormattedApiApiKeyListResponse) SetOrderBy(v string) {
	o.OrderBy = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *FormattedApiApiKeyListResponse) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKeyListResponse) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *FormattedApiApiKeyListResponse) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *FormattedApiApiKeyListResponse) SetDirection(v string) {
	o.Direction = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *FormattedApiApiKeyListResponse) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKeyListResponse) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *FormattedApiApiKeyListResponse) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *FormattedApiApiKeyListResponse) SetTotal(v float32) {
	o.Total = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *FormattedApiApiKeyListResponse) GetPages() float32 {
	if o == nil || IsNil(o.Pages) {
		var ret float32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKeyListResponse) GetPagesOk() (*float32, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *FormattedApiApiKeyListResponse) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given float32 and assigns it to the Pages field.
func (o *FormattedApiApiKeyListResponse) SetPages(v float32) {
	o.Pages = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *FormattedApiApiKeyListResponse) GetPageSize() float32 {
	if o == nil || IsNil(o.PageSize) {
		var ret float32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKeyListResponse) GetPageSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *FormattedApiApiKeyListResponse) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given float32 and assigns it to the PageSize field.
func (o *FormattedApiApiKeyListResponse) SetPageSize(v float32) {
	o.PageSize = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *FormattedApiApiKeyListResponse) GetPage() float32 {
	if o == nil || IsNil(o.Page) {
		var ret float32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKeyListResponse) GetPageOk() (*float32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *FormattedApiApiKeyListResponse) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given float32 and assigns it to the Page field.
func (o *FormattedApiApiKeyListResponse) SetPage(v float32) {
	o.Page = &v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *FormattedApiApiKeyListResponse) GetCursor() float32 {
	if o == nil || IsNil(o.Cursor) {
		var ret float32
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKeyListResponse) GetCursorOk() (*float32, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *FormattedApiApiKeyListResponse) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given float32 and assigns it to the Cursor field.
func (o *FormattedApiApiKeyListResponse) SetCursor(v float32) {
	o.Cursor = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *FormattedApiApiKeyListResponse) GetNextCursor() float32 {
	if o == nil || IsNil(o.NextCursor) {
		var ret float32
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKeyListResponse) GetNextCursorOk() (*float32, bool) {
	if o == nil || IsNil(o.NextCursor) {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *FormattedApiApiKeyListResponse) HasNextCursor() bool {
	if o != nil && !IsNil(o.NextCursor) {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given float32 and assigns it to the NextCursor field.
func (o *FormattedApiApiKeyListResponse) SetNextCursor(v float32) {
	o.NextCursor = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FormattedApiApiKeyListResponse) GetLinks() []LinksInner {
	if o == nil || IsNil(o.Links) {
		var ret []LinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKeyListResponse) GetLinksOk() ([]LinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FormattedApiApiKeyListResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinksInner and assigns it to the Links field.
func (o *FormattedApiApiKeyListResponse) SetLinks(v []LinksInner) {
	o.Links = v
}

func (o FormattedApiApiKeyListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormattedApiApiKeyListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	if !IsNil(o.OrderBy) {
		toSerialize["orderBy"] = o.OrderBy
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Pages) {
		toSerialize["pages"] = o.Pages
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	if !IsNil(o.NextCursor) {
		toSerialize["nextCursor"] = o.NextCursor
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormattedApiApiKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varFormattedApiApiKeyListResponse := _FormattedApiApiKeyListResponse{}

	err = json.Unmarshal(data, &varFormattedApiApiKeyListResponse)

	if err != nil {
		return err
	}

	*o = FormattedApiApiKeyListResponse(varFormattedApiApiKeyListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "orderBy")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "total")
		delete(additionalProperties, "pages")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "page")
		delete(additionalProperties, "cursor")
		delete(additionalProperties, "nextCursor")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormattedApiApiKeyListResponse struct {
	value *FormattedApiApiKeyListResponse
	isSet bool
}

func (v NullableFormattedApiApiKeyListResponse) Get() *FormattedApiApiKeyListResponse {
	return v.value
}

func (v *NullableFormattedApiApiKeyListResponse) Set(val *FormattedApiApiKeyListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFormattedApiApiKeyListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFormattedApiApiKeyListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormattedApiApiKeyListResponse(val *FormattedApiApiKeyListResponse) *NullableFormattedApiApiKeyListResponse {
	return &NullableFormattedApiApiKeyListResponse{value: val, isSet: true}
}

func (v NullableFormattedApiApiKeyListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormattedApiApiKeyListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
