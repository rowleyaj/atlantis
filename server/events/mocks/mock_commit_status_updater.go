// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: CommitStatusUpdater)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockCommitStatusUpdater struct {
	fail func(message string, callerSkip ...int)
}

func NewMockCommitStatusUpdater(options ...pegomock.Option) *MockCommitStatusUpdater {
	mock := &MockCommitStatusUpdater{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockCommitStatusUpdater) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockCommitStatusUpdater) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockCommitStatusUpdater) UpdateCombined(repo models.Repo, pull models.PullRequest, status models.CommitStatus, command models.CommandName) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommitStatusUpdater().")
	}
	params := []pegomock.Param{repo, pull, status, command}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdateCombined", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockCommitStatusUpdater) UpdateCombinedCount(repo models.Repo, pull models.PullRequest, status models.CommitStatus, command models.CommandName, numSuccess int, numTotal int) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommitStatusUpdater().")
	}
	params := []pegomock.Param{repo, pull, status, command, numSuccess, numTotal}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdateCombinedCount", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockCommitStatusUpdater) UpdateProject(ctx models.ProjectCommandContext, cmdName models.CommandName, status models.CommitStatus, url string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommitStatusUpdater().")
	}
	params := []pegomock.Param{ctx, cmdName, status, url}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdateProject", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockCommitStatusUpdater) VerifyWasCalledOnce() *VerifierMockCommitStatusUpdater {
	return &VerifierMockCommitStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockCommitStatusUpdater) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockCommitStatusUpdater {
	return &VerifierMockCommitStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockCommitStatusUpdater) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockCommitStatusUpdater {
	return &VerifierMockCommitStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockCommitStatusUpdater) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockCommitStatusUpdater {
	return &VerifierMockCommitStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockCommitStatusUpdater struct {
	mock                   *MockCommitStatusUpdater
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockCommitStatusUpdater) UpdateCombined(repo models.Repo, pull models.PullRequest, status models.CommitStatus, command models.CommandName) *MockCommitStatusUpdater_UpdateCombined_OngoingVerification {
	params := []pegomock.Param{repo, pull, status, command}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdateCombined", params, verifier.timeout)
	return &MockCommitStatusUpdater_UpdateCombined_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommitStatusUpdater_UpdateCombined_OngoingVerification struct {
	mock              *MockCommitStatusUpdater
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommitStatusUpdater_UpdateCombined_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, models.CommitStatus, models.CommandName) {
	repo, pull, status, command := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1], status[len(status)-1], command[len(command)-1]
}

func (c *MockCommitStatusUpdater_UpdateCombined_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []models.CommitStatus, _param3 []models.CommandName) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]models.CommitStatus, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(models.CommitStatus)
		}
		_param3 = make([]models.CommandName, len(params[3]))
		for u, param := range params[3] {
			_param3[u] = param.(models.CommandName)
		}
	}
	return
}

func (verifier *VerifierMockCommitStatusUpdater) UpdateCombinedCount(repo models.Repo, pull models.PullRequest, status models.CommitStatus, command models.CommandName, numSuccess int, numTotal int) *MockCommitStatusUpdater_UpdateCombinedCount_OngoingVerification {
	params := []pegomock.Param{repo, pull, status, command, numSuccess, numTotal}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdateCombinedCount", params, verifier.timeout)
	return &MockCommitStatusUpdater_UpdateCombinedCount_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommitStatusUpdater_UpdateCombinedCount_OngoingVerification struct {
	mock              *MockCommitStatusUpdater
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommitStatusUpdater_UpdateCombinedCount_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, models.CommitStatus, models.CommandName, int, int) {
	repo, pull, status, command, numSuccess, numTotal := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1], status[len(status)-1], command[len(command)-1], numSuccess[len(numSuccess)-1], numTotal[len(numTotal)-1]
}

func (c *MockCommitStatusUpdater_UpdateCombinedCount_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []models.CommitStatus, _param3 []models.CommandName, _param4 []int, _param5 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]models.CommitStatus, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(models.CommitStatus)
		}
		_param3 = make([]models.CommandName, len(params[3]))
		for u, param := range params[3] {
			_param3[u] = param.(models.CommandName)
		}
		_param4 = make([]int, len(params[4]))
		for u, param := range params[4] {
			_param4[u] = param.(int)
		}
		_param5 = make([]int, len(params[5]))
		for u, param := range params[5] {
			_param5[u] = param.(int)
		}
	}
	return
}

func (verifier *VerifierMockCommitStatusUpdater) UpdateProject(ctx models.ProjectCommandContext, cmdName models.CommandName, status models.CommitStatus, url string) *MockCommitStatusUpdater_UpdateProject_OngoingVerification {
	params := []pegomock.Param{ctx, cmdName, status, url}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdateProject", params, verifier.timeout)
	return &MockCommitStatusUpdater_UpdateProject_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommitStatusUpdater_UpdateProject_OngoingVerification struct {
	mock              *MockCommitStatusUpdater
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommitStatusUpdater_UpdateProject_OngoingVerification) GetCapturedArguments() (models.ProjectCommandContext, models.CommandName, models.CommitStatus, string) {
	ctx, cmdName, status, url := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], cmdName[len(cmdName)-1], status[len(status)-1], url[len(url)-1]
}

func (c *MockCommitStatusUpdater_UpdateProject_OngoingVerification) GetAllCapturedArguments() (_param0 []models.ProjectCommandContext, _param1 []models.CommandName, _param2 []models.CommitStatus, _param3 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.ProjectCommandContext, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.ProjectCommandContext)
		}
		_param1 = make([]models.CommandName, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(models.CommandName)
		}
		_param2 = make([]models.CommitStatus, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(models.CommitStatus)
		}
		_param3 = make([]string, len(params[3]))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
	}
	return
}
