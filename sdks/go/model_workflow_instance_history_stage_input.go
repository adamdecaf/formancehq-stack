/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: develop
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the WorkflowInstanceHistoryStageInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowInstanceHistoryStageInput{}

// WorkflowInstanceHistoryStageInput struct for WorkflowInstanceHistoryStageInput
type WorkflowInstanceHistoryStageInput struct {
	GetAccount *ActivityGetAccount `json:"GetAccount,omitempty"`
	CreateTransaction *ActivityCreateTransaction `json:"CreateTransaction,omitempty"`
	RevertTransaction *ActivityRevertTransaction `json:"RevertTransaction,omitempty"`
	StripeTransfer *StripeTransferRequest `json:"StripeTransfer,omitempty"`
	GetPayment *ActivityGetPayment `json:"GetPayment,omitempty"`
	ConfirmHold *ActivityConfirmHold `json:"ConfirmHold,omitempty"`
	CreditWallet *ActivityCreditWallet `json:"CreditWallet,omitempty"`
	DebitWallet *ActivityDebitWallet `json:"DebitWallet,omitempty"`
	GetWallet *ActivityGetWallet `json:"GetWallet,omitempty"`
	VoidHold *ActivityVoidHold `json:"VoidHold,omitempty"`
}

// NewWorkflowInstanceHistoryStageInput instantiates a new WorkflowInstanceHistoryStageInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowInstanceHistoryStageInput() *WorkflowInstanceHistoryStageInput {
	this := WorkflowInstanceHistoryStageInput{}
	return &this
}

// NewWorkflowInstanceHistoryStageInputWithDefaults instantiates a new WorkflowInstanceHistoryStageInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowInstanceHistoryStageInputWithDefaults() *WorkflowInstanceHistoryStageInput {
	this := WorkflowInstanceHistoryStageInput{}
	return &this
}

// GetGetAccount returns the GetAccount field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageInput) GetGetAccount() ActivityGetAccount {
	if o == nil || IsNil(o.GetAccount) {
		var ret ActivityGetAccount
		return ret
	}
	return *o.GetAccount
}

// GetGetAccountOk returns a tuple with the GetAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageInput) GetGetAccountOk() (*ActivityGetAccount, bool) {
	if o == nil || IsNil(o.GetAccount) {
		return nil, false
	}
	return o.GetAccount, true
}

// HasGetAccount returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageInput) HasGetAccount() bool {
	if o != nil && !IsNil(o.GetAccount) {
		return true
	}

	return false
}

// SetGetAccount gets a reference to the given ActivityGetAccount and assigns it to the GetAccount field.
func (o *WorkflowInstanceHistoryStageInput) SetGetAccount(v ActivityGetAccount) {
	o.GetAccount = &v
}

// GetCreateTransaction returns the CreateTransaction field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageInput) GetCreateTransaction() ActivityCreateTransaction {
	if o == nil || IsNil(o.CreateTransaction) {
		var ret ActivityCreateTransaction
		return ret
	}
	return *o.CreateTransaction
}

// GetCreateTransactionOk returns a tuple with the CreateTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageInput) GetCreateTransactionOk() (*ActivityCreateTransaction, bool) {
	if o == nil || IsNil(o.CreateTransaction) {
		return nil, false
	}
	return o.CreateTransaction, true
}

// HasCreateTransaction returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageInput) HasCreateTransaction() bool {
	if o != nil && !IsNil(o.CreateTransaction) {
		return true
	}

	return false
}

// SetCreateTransaction gets a reference to the given ActivityCreateTransaction and assigns it to the CreateTransaction field.
func (o *WorkflowInstanceHistoryStageInput) SetCreateTransaction(v ActivityCreateTransaction) {
	o.CreateTransaction = &v
}

// GetRevertTransaction returns the RevertTransaction field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageInput) GetRevertTransaction() ActivityRevertTransaction {
	if o == nil || IsNil(o.RevertTransaction) {
		var ret ActivityRevertTransaction
		return ret
	}
	return *o.RevertTransaction
}

// GetRevertTransactionOk returns a tuple with the RevertTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageInput) GetRevertTransactionOk() (*ActivityRevertTransaction, bool) {
	if o == nil || IsNil(o.RevertTransaction) {
		return nil, false
	}
	return o.RevertTransaction, true
}

// HasRevertTransaction returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageInput) HasRevertTransaction() bool {
	if o != nil && !IsNil(o.RevertTransaction) {
		return true
	}

	return false
}

// SetRevertTransaction gets a reference to the given ActivityRevertTransaction and assigns it to the RevertTransaction field.
func (o *WorkflowInstanceHistoryStageInput) SetRevertTransaction(v ActivityRevertTransaction) {
	o.RevertTransaction = &v
}

// GetStripeTransfer returns the StripeTransfer field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageInput) GetStripeTransfer() StripeTransferRequest {
	if o == nil || IsNil(o.StripeTransfer) {
		var ret StripeTransferRequest
		return ret
	}
	return *o.StripeTransfer
}

// GetStripeTransferOk returns a tuple with the StripeTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageInput) GetStripeTransferOk() (*StripeTransferRequest, bool) {
	if o == nil || IsNil(o.StripeTransfer) {
		return nil, false
	}
	return o.StripeTransfer, true
}

// HasStripeTransfer returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageInput) HasStripeTransfer() bool {
	if o != nil && !IsNil(o.StripeTransfer) {
		return true
	}

	return false
}

// SetStripeTransfer gets a reference to the given StripeTransferRequest and assigns it to the StripeTransfer field.
func (o *WorkflowInstanceHistoryStageInput) SetStripeTransfer(v StripeTransferRequest) {
	o.StripeTransfer = &v
}

// GetGetPayment returns the GetPayment field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageInput) GetGetPayment() ActivityGetPayment {
	if o == nil || IsNil(o.GetPayment) {
		var ret ActivityGetPayment
		return ret
	}
	return *o.GetPayment
}

// GetGetPaymentOk returns a tuple with the GetPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageInput) GetGetPaymentOk() (*ActivityGetPayment, bool) {
	if o == nil || IsNil(o.GetPayment) {
		return nil, false
	}
	return o.GetPayment, true
}

// HasGetPayment returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageInput) HasGetPayment() bool {
	if o != nil && !IsNil(o.GetPayment) {
		return true
	}

	return false
}

// SetGetPayment gets a reference to the given ActivityGetPayment and assigns it to the GetPayment field.
func (o *WorkflowInstanceHistoryStageInput) SetGetPayment(v ActivityGetPayment) {
	o.GetPayment = &v
}

// GetConfirmHold returns the ConfirmHold field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageInput) GetConfirmHold() ActivityConfirmHold {
	if o == nil || IsNil(o.ConfirmHold) {
		var ret ActivityConfirmHold
		return ret
	}
	return *o.ConfirmHold
}

// GetConfirmHoldOk returns a tuple with the ConfirmHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageInput) GetConfirmHoldOk() (*ActivityConfirmHold, bool) {
	if o == nil || IsNil(o.ConfirmHold) {
		return nil, false
	}
	return o.ConfirmHold, true
}

// HasConfirmHold returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageInput) HasConfirmHold() bool {
	if o != nil && !IsNil(o.ConfirmHold) {
		return true
	}

	return false
}

// SetConfirmHold gets a reference to the given ActivityConfirmHold and assigns it to the ConfirmHold field.
func (o *WorkflowInstanceHistoryStageInput) SetConfirmHold(v ActivityConfirmHold) {
	o.ConfirmHold = &v
}

// GetCreditWallet returns the CreditWallet field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageInput) GetCreditWallet() ActivityCreditWallet {
	if o == nil || IsNil(o.CreditWallet) {
		var ret ActivityCreditWallet
		return ret
	}
	return *o.CreditWallet
}

// GetCreditWalletOk returns a tuple with the CreditWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageInput) GetCreditWalletOk() (*ActivityCreditWallet, bool) {
	if o == nil || IsNil(o.CreditWallet) {
		return nil, false
	}
	return o.CreditWallet, true
}

// HasCreditWallet returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageInput) HasCreditWallet() bool {
	if o != nil && !IsNil(o.CreditWallet) {
		return true
	}

	return false
}

// SetCreditWallet gets a reference to the given ActivityCreditWallet and assigns it to the CreditWallet field.
func (o *WorkflowInstanceHistoryStageInput) SetCreditWallet(v ActivityCreditWallet) {
	o.CreditWallet = &v
}

// GetDebitWallet returns the DebitWallet field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageInput) GetDebitWallet() ActivityDebitWallet {
	if o == nil || IsNil(o.DebitWallet) {
		var ret ActivityDebitWallet
		return ret
	}
	return *o.DebitWallet
}

// GetDebitWalletOk returns a tuple with the DebitWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageInput) GetDebitWalletOk() (*ActivityDebitWallet, bool) {
	if o == nil || IsNil(o.DebitWallet) {
		return nil, false
	}
	return o.DebitWallet, true
}

// HasDebitWallet returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageInput) HasDebitWallet() bool {
	if o != nil && !IsNil(o.DebitWallet) {
		return true
	}

	return false
}

// SetDebitWallet gets a reference to the given ActivityDebitWallet and assigns it to the DebitWallet field.
func (o *WorkflowInstanceHistoryStageInput) SetDebitWallet(v ActivityDebitWallet) {
	o.DebitWallet = &v
}

// GetGetWallet returns the GetWallet field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageInput) GetGetWallet() ActivityGetWallet {
	if o == nil || IsNil(o.GetWallet) {
		var ret ActivityGetWallet
		return ret
	}
	return *o.GetWallet
}

// GetGetWalletOk returns a tuple with the GetWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageInput) GetGetWalletOk() (*ActivityGetWallet, bool) {
	if o == nil || IsNil(o.GetWallet) {
		return nil, false
	}
	return o.GetWallet, true
}

// HasGetWallet returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageInput) HasGetWallet() bool {
	if o != nil && !IsNil(o.GetWallet) {
		return true
	}

	return false
}

// SetGetWallet gets a reference to the given ActivityGetWallet and assigns it to the GetWallet field.
func (o *WorkflowInstanceHistoryStageInput) SetGetWallet(v ActivityGetWallet) {
	o.GetWallet = &v
}

// GetVoidHold returns the VoidHold field value if set, zero value otherwise.
func (o *WorkflowInstanceHistoryStageInput) GetVoidHold() ActivityVoidHold {
	if o == nil || IsNil(o.VoidHold) {
		var ret ActivityVoidHold
		return ret
	}
	return *o.VoidHold
}

// GetVoidHoldOk returns a tuple with the VoidHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceHistoryStageInput) GetVoidHoldOk() (*ActivityVoidHold, bool) {
	if o == nil || IsNil(o.VoidHold) {
		return nil, false
	}
	return o.VoidHold, true
}

// HasVoidHold returns a boolean if a field has been set.
func (o *WorkflowInstanceHistoryStageInput) HasVoidHold() bool {
	if o != nil && !IsNil(o.VoidHold) {
		return true
	}

	return false
}

// SetVoidHold gets a reference to the given ActivityVoidHold and assigns it to the VoidHold field.
func (o *WorkflowInstanceHistoryStageInput) SetVoidHold(v ActivityVoidHold) {
	o.VoidHold = &v
}

func (o WorkflowInstanceHistoryStageInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowInstanceHistoryStageInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GetAccount) {
		toSerialize["GetAccount"] = o.GetAccount
	}
	if !IsNil(o.CreateTransaction) {
		toSerialize["CreateTransaction"] = o.CreateTransaction
	}
	if !IsNil(o.RevertTransaction) {
		toSerialize["RevertTransaction"] = o.RevertTransaction
	}
	if !IsNil(o.StripeTransfer) {
		toSerialize["StripeTransfer"] = o.StripeTransfer
	}
	if !IsNil(o.GetPayment) {
		toSerialize["GetPayment"] = o.GetPayment
	}
	if !IsNil(o.ConfirmHold) {
		toSerialize["ConfirmHold"] = o.ConfirmHold
	}
	if !IsNil(o.CreditWallet) {
		toSerialize["CreditWallet"] = o.CreditWallet
	}
	if !IsNil(o.DebitWallet) {
		toSerialize["DebitWallet"] = o.DebitWallet
	}
	if !IsNil(o.GetWallet) {
		toSerialize["GetWallet"] = o.GetWallet
	}
	if !IsNil(o.VoidHold) {
		toSerialize["VoidHold"] = o.VoidHold
	}
	return toSerialize, nil
}

type NullableWorkflowInstanceHistoryStageInput struct {
	value *WorkflowInstanceHistoryStageInput
	isSet bool
}

func (v NullableWorkflowInstanceHistoryStageInput) Get() *WorkflowInstanceHistoryStageInput {
	return v.value
}

func (v *NullableWorkflowInstanceHistoryStageInput) Set(val *WorkflowInstanceHistoryStageInput) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowInstanceHistoryStageInput) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowInstanceHistoryStageInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowInstanceHistoryStageInput(val *WorkflowInstanceHistoryStageInput) *NullableWorkflowInstanceHistoryStageInput {
	return &NullableWorkflowInstanceHistoryStageInput{value: val, isSet: true}
}

func (v NullableWorkflowInstanceHistoryStageInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowInstanceHistoryStageInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


