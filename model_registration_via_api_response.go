/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.6.1-alpha.1
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RegistrationViaApiResponse The Response for Registration Flows via API
type RegistrationViaApiResponse struct {
	Identity Identity `json:"identity"`
	Session *Session `json:"session,omitempty"`
	// The Session Token  This field is only set when the session hook is configured as a post-registration hook.  A session token is equivalent to a session cookie, but it can be sent in the HTTP Authorization Header:  Authorization: bearer ${session-token}  The session token is only issued for API flows, not for Browser flows!
	SessionToken string `json:"session_token"`
}

// NewRegistrationViaApiResponse instantiates a new RegistrationViaApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationViaApiResponse(identity Identity, sessionToken string) *RegistrationViaApiResponse {
	this := RegistrationViaApiResponse{}
	this.Identity = identity
	this.SessionToken = sessionToken
	return &this
}

// NewRegistrationViaApiResponseWithDefaults instantiates a new RegistrationViaApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationViaApiResponseWithDefaults() *RegistrationViaApiResponse {
	this := RegistrationViaApiResponse{}
	return &this
}

// GetIdentity returns the Identity field value
func (o *RegistrationViaApiResponse) GetIdentity() Identity {
	if o == nil {
		var ret Identity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *RegistrationViaApiResponse) GetIdentityOk() (*Identity, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *RegistrationViaApiResponse) SetIdentity(v Identity) {
	o.Identity = v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *RegistrationViaApiResponse) GetSession() Session {
	if o == nil || o.Session == nil {
		var ret Session
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationViaApiResponse) GetSessionOk() (*Session, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *RegistrationViaApiResponse) HasSession() bool {
	if o != nil && o.Session != nil {
		return true
	}

	return false
}

// SetSession gets a reference to the given Session and assigns it to the Session field.
func (o *RegistrationViaApiResponse) SetSession(v Session) {
	o.Session = &v
}

// GetSessionToken returns the SessionToken field value
func (o *RegistrationViaApiResponse) GetSessionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value
// and a boolean to check if the value has been set.
func (o *RegistrationViaApiResponse) GetSessionTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SessionToken, true
}

// SetSessionToken sets field value
func (o *RegistrationViaApiResponse) SetSessionToken(v string) {
	o.SessionToken = v
}

func (o RegistrationViaApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identity"] = o.Identity
	}
	if o.Session != nil {
		toSerialize["session"] = o.Session
	}
	if true {
		toSerialize["session_token"] = o.SessionToken
	}
	return json.Marshal(toSerialize)
}

type NullableRegistrationViaApiResponse struct {
	value *RegistrationViaApiResponse
	isSet bool
}

func (v NullableRegistrationViaApiResponse) Get() *RegistrationViaApiResponse {
	return v.value
}

func (v *NullableRegistrationViaApiResponse) Set(val *RegistrationViaApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationViaApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationViaApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationViaApiResponse(val *RegistrationViaApiResponse) *NullableRegistrationViaApiResponse {
	return &NullableRegistrationViaApiResponse{value: val, isSet: true}
}

func (v NullableRegistrationViaApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationViaApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


