/*
Load Balancer Management API

Load Balancer Management API is an API for managing load balancers.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// checks if the LoadBalancerPoolOrigin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerPoolOrigin{}

// LoadBalancerPoolOrigin struct for LoadBalancerPoolOrigin
type LoadBalancerPoolOrigin struct {
	// ID of a port
	Id string `json:"id"`
	// Date and time of creation
	CreatedAt time.Time `json:"created_at"`
	// Date and time of last update
	UpdatedAt time.Time `json:"updated_at"`
	// A name for the origin
	Name string `json:"name"`
	// IP address of the origin
	Target     string                           `json:"target"`
	PortNumber LoadBalancerPoolOriginPortNumber `json:"port_number"`
	// If the origin is enabled
	Active bool `json:"active"`
	// ID of the pool the origin belongs to
	PoolId               string `json:"pool_id"`
	AdditionalProperties map[string]interface{}
}

type _LoadBalancerPoolOrigin LoadBalancerPoolOrigin

// NewLoadBalancerPoolOrigin instantiates a new LoadBalancerPoolOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerPoolOrigin(id string, createdAt time.Time, updatedAt time.Time, name string, target string, portNumber LoadBalancerPoolOriginPortNumber, active bool, poolId string) *LoadBalancerPoolOrigin {
	this := LoadBalancerPoolOrigin{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	this.Target = target
	this.PortNumber = portNumber
	this.Active = active
	this.PoolId = poolId
	return &this
}

// NewLoadBalancerPoolOriginWithDefaults instantiates a new LoadBalancerPoolOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerPoolOriginWithDefaults() *LoadBalancerPoolOrigin {
	this := LoadBalancerPoolOrigin{}
	return &this
}

// GetId returns the Id field value
func (o *LoadBalancerPoolOrigin) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolOrigin) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoadBalancerPoolOrigin) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *LoadBalancerPoolOrigin) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolOrigin) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LoadBalancerPoolOrigin) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *LoadBalancerPoolOrigin) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolOrigin) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *LoadBalancerPoolOrigin) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *LoadBalancerPoolOrigin) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolOrigin) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LoadBalancerPoolOrigin) SetName(v string) {
	o.Name = v
}

// GetTarget returns the Target field value
func (o *LoadBalancerPoolOrigin) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolOrigin) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *LoadBalancerPoolOrigin) SetTarget(v string) {
	o.Target = v
}

// GetPortNumber returns the PortNumber field value
func (o *LoadBalancerPoolOrigin) GetPortNumber() LoadBalancerPoolOriginPortNumber {
	if o == nil {
		var ret LoadBalancerPoolOriginPortNumber
		return ret
	}

	return o.PortNumber
}

// GetPortNumberOk returns a tuple with the PortNumber field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolOrigin) GetPortNumberOk() (*LoadBalancerPoolOriginPortNumber, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PortNumber, true
}

// SetPortNumber sets field value
func (o *LoadBalancerPoolOrigin) SetPortNumber(v LoadBalancerPoolOriginPortNumber) {
	o.PortNumber = v
}

// GetActive returns the Active field value
func (o *LoadBalancerPoolOrigin) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolOrigin) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *LoadBalancerPoolOrigin) SetActive(v bool) {
	o.Active = v
}

// GetPoolId returns the PoolId field value
func (o *LoadBalancerPoolOrigin) GetPoolId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolOrigin) GetPoolIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolId, true
}

// SetPoolId sets field value
func (o *LoadBalancerPoolOrigin) SetPoolId(v string) {
	o.PoolId = v
}

func (o LoadBalancerPoolOrigin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerPoolOrigin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["name"] = o.Name
	toSerialize["target"] = o.Target
	toSerialize["port_number"] = o.PortNumber
	toSerialize["active"] = o.Active
	toSerialize["pool_id"] = o.PoolId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadBalancerPoolOrigin) UnmarshalJSON(bytes []byte) (err error) {
	varLoadBalancerPoolOrigin := _LoadBalancerPoolOrigin{}

	err = json.Unmarshal(bytes, &varLoadBalancerPoolOrigin)

	if err != nil {
		return err
	}

	*o = LoadBalancerPoolOrigin(varLoadBalancerPoolOrigin)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "name")
		delete(additionalProperties, "target")
		delete(additionalProperties, "port_number")
		delete(additionalProperties, "active")
		delete(additionalProperties, "pool_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadBalancerPoolOrigin struct {
	value *LoadBalancerPoolOrigin
	isSet bool
}

func (v NullableLoadBalancerPoolOrigin) Get() *LoadBalancerPoolOrigin {
	return v.value
}

func (v *NullableLoadBalancerPoolOrigin) Set(val *LoadBalancerPoolOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerPoolOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerPoolOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerPoolOrigin(val *LoadBalancerPoolOrigin) *NullableLoadBalancerPoolOrigin {
	return &NullableLoadBalancerPoolOrigin{value: val, isSet: true}
}

func (v NullableLoadBalancerPoolOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerPoolOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
