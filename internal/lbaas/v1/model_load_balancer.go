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

// checks if the LoadBalancer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancer{}

// LoadBalancer struct for LoadBalancer
type LoadBalancer struct {
	// ID of the load balancer
	Id string `json:"id"`
	// Date and time of creation
	CreatedAt time.Time `json:"created_at"`
	// Date and time of last update
	UpdatedAt time.Time `json:"updated_at"`
	// A name for the load balancer
	Name     string   `json:"name"`
	Provider Provider `json:"provider"`
	// A list of ports assigned to the load balancer
	Ports []LoadBalancerPort `json:"ports"`
	// A list of pool names and ids assigned to the load balancer
	Pools []LoadBalancerPoolShort `json:"pools,omitempty"`
	// A list of associated ip addresses
	Ips                  []string              `json:"ips,omitempty"`
	Location             *LoadBalancerLocation `json:"location,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoadBalancer LoadBalancer

// NewLoadBalancer instantiates a new LoadBalancer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancer(id string, createdAt time.Time, updatedAt time.Time, name string, provider Provider, ports []LoadBalancerPort) *LoadBalancer {
	this := LoadBalancer{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	this.Provider = provider
	this.Ports = ports
	return &this
}

// NewLoadBalancerWithDefaults instantiates a new LoadBalancer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerWithDefaults() *LoadBalancer {
	this := LoadBalancer{}
	return &this
}

// GetId returns the Id field value
func (o *LoadBalancer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoadBalancer) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *LoadBalancer) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LoadBalancer) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *LoadBalancer) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *LoadBalancer) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *LoadBalancer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LoadBalancer) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value
func (o *LoadBalancer) GetProvider() Provider {
	if o == nil {
		var ret Provider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetProviderOk() (*Provider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *LoadBalancer) SetProvider(v Provider) {
	o.Provider = v
}

// GetPorts returns the Ports field value
func (o *LoadBalancer) GetPorts() []LoadBalancerPort {
	if o == nil {
		var ret []LoadBalancerPort
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetPortsOk() ([]LoadBalancerPort, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ports, true
}

// SetPorts sets field value
func (o *LoadBalancer) SetPorts(v []LoadBalancerPort) {
	o.Ports = v
}

// GetPools returns the Pools field value if set, zero value otherwise.
func (o *LoadBalancer) GetPools() []LoadBalancerPoolShort {
	if o == nil || IsNil(o.Pools) {
		var ret []LoadBalancerPoolShort
		return ret
	}
	return o.Pools
}

// GetPoolsOk returns a tuple with the Pools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetPoolsOk() ([]LoadBalancerPoolShort, bool) {
	if o == nil || IsNil(o.Pools) {
		return nil, false
	}
	return o.Pools, true
}

// HasPools returns a boolean if a field has been set.
func (o *LoadBalancer) HasPools() bool {
	if o != nil && !IsNil(o.Pools) {
		return true
	}

	return false
}

// SetPools gets a reference to the given []LoadBalancerPoolShort and assigns it to the Pools field.
func (o *LoadBalancer) SetPools(v []LoadBalancerPoolShort) {
	o.Pools = v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *LoadBalancer) GetIps() []string {
	if o == nil || IsNil(o.Ips) {
		var ret []string
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *LoadBalancer) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given []string and assigns it to the Ips field.
func (o *LoadBalancer) SetIps(v []string) {
	o.Ips = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *LoadBalancer) GetLocation() LoadBalancerLocation {
	if o == nil || IsNil(o.Location) {
		var ret LoadBalancerLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetLocationOk() (*LoadBalancerLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *LoadBalancer) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given LoadBalancerLocation and assigns it to the Location field.
func (o *LoadBalancer) SetLocation(v LoadBalancerLocation) {
	o.Location = &v
}

func (o LoadBalancer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["name"] = o.Name
	toSerialize["provider"] = o.Provider
	toSerialize["ports"] = o.Ports
	if !IsNil(o.Pools) {
		toSerialize["pools"] = o.Pools
	}
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadBalancer) UnmarshalJSON(bytes []byte) (err error) {
	varLoadBalancer := _LoadBalancer{}

	err = json.Unmarshal(bytes, &varLoadBalancer)

	if err != nil {
		return err
	}

	*o = LoadBalancer(varLoadBalancer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "name")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "ports")
		delete(additionalProperties, "pools")
		delete(additionalProperties, "ips")
		delete(additionalProperties, "location")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadBalancer struct {
	value *LoadBalancer
	isSet bool
}

func (v NullableLoadBalancer) Get() *LoadBalancer {
	return v.value
}

func (v *NullableLoadBalancer) Set(val *LoadBalancer) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancer) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancer(val *LoadBalancer) *NullableLoadBalancer {
	return &NullableLoadBalancer{value: val, isSet: true}
}

func (v NullableLoadBalancer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
