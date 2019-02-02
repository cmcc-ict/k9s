// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/derailed/k9s/resource (interfaces: ClusterIfc)

package resource_test

import (
	pegomock "github.com/petergtz/pegomock"
	"reflect"
	"time"
)

type MockClusterIfc struct {
	fail func(message string, callerSkip ...int)
}

func NewMockClusterIfc() *MockClusterIfc {
	return &MockClusterIfc{fail: pegomock.GlobalFailHandler}
}

func (mock *MockClusterIfc) ClusterName() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterIfc().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ClusterName", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockClusterIfc) Version() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterIfc().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Version", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClusterIfc) VerifyWasCalledOnce() *VerifierClusterIfc {
	return &VerifierClusterIfc{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockClusterIfc) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierClusterIfc {
	return &VerifierClusterIfc{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockClusterIfc) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierClusterIfc {
	return &VerifierClusterIfc{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockClusterIfc) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierClusterIfc {
	return &VerifierClusterIfc{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierClusterIfc struct {
	mock                   *MockClusterIfc
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierClusterIfc) ClusterName() *ClusterIfc_ClusterName_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ClusterName", params, verifier.timeout)
	return &ClusterIfc_ClusterName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterIfc_ClusterName_OngoingVerification struct {
	mock              *MockClusterIfc
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterIfc_ClusterName_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterIfc_ClusterName_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierClusterIfc) Version() *ClusterIfc_Version_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Version", params, verifier.timeout)
	return &ClusterIfc_Version_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterIfc_Version_OngoingVerification struct {
	mock              *MockClusterIfc
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterIfc_Version_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterIfc_Version_OngoingVerification) GetAllCapturedArguments() {
}