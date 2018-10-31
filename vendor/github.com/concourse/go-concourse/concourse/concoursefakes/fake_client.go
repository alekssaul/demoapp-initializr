// Code generated by counterfeiter. DO NOT EDIT.
package concoursefakes

import (
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/concourse/atc"
	"github.com/concourse/go-concourse/concourse"
)

type FakeClient struct {
	URLStub        func() string
	uRLMutex       sync.RWMutex
	uRLArgsForCall []struct{}
	uRLReturns     struct {
		result1 string
	}
	uRLReturnsOnCall map[int]struct {
		result1 string
	}
	HTTPClientStub        func() *http.Client
	hTTPClientMutex       sync.RWMutex
	hTTPClientArgsForCall []struct{}
	hTTPClientReturns     struct {
		result1 *http.Client
	}
	hTTPClientReturnsOnCall map[int]struct {
		result1 *http.Client
	}
	BuildsStub        func(concourse.Page) ([]atc.Build, concourse.Pagination, error)
	buildsMutex       sync.RWMutex
	buildsArgsForCall []struct {
		arg1 concourse.Page
	}
	buildsReturns struct {
		result1 []atc.Build
		result2 concourse.Pagination
		result3 error
	}
	buildsReturnsOnCall map[int]struct {
		result1 []atc.Build
		result2 concourse.Pagination
		result3 error
	}
	BuildStub        func(buildID string) (atc.Build, bool, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		buildID string
	}
	buildReturns struct {
		result1 atc.Build
		result2 bool
		result3 error
	}
	buildReturnsOnCall map[int]struct {
		result1 atc.Build
		result2 bool
		result3 error
	}
	BuildEventsStub        func(buildID string) (concourse.Events, error)
	buildEventsMutex       sync.RWMutex
	buildEventsArgsForCall []struct {
		buildID string
	}
	buildEventsReturns struct {
		result1 concourse.Events
		result2 error
	}
	buildEventsReturnsOnCall map[int]struct {
		result1 concourse.Events
		result2 error
	}
	BuildResourcesStub        func(buildID int) (atc.BuildInputsOutputs, bool, error)
	buildResourcesMutex       sync.RWMutex
	buildResourcesArgsForCall []struct {
		buildID int
	}
	buildResourcesReturns struct {
		result1 atc.BuildInputsOutputs
		result2 bool
		result3 error
	}
	buildResourcesReturnsOnCall map[int]struct {
		result1 atc.BuildInputsOutputs
		result2 bool
		result3 error
	}
	AbortBuildStub        func(buildID string) error
	abortBuildMutex       sync.RWMutex
	abortBuildArgsForCall []struct {
		buildID string
	}
	abortBuildReturns struct {
		result1 error
	}
	abortBuildReturnsOnCall map[int]struct {
		result1 error
	}
	BuildPlanStub        func(buildID int) (atc.PublicBuildPlan, bool, error)
	buildPlanMutex       sync.RWMutex
	buildPlanArgsForCall []struct {
		buildID int
	}
	buildPlanReturns struct {
		result1 atc.PublicBuildPlan
		result2 bool
		result3 error
	}
	buildPlanReturnsOnCall map[int]struct {
		result1 atc.PublicBuildPlan
		result2 bool
		result3 error
	}
	SendInputToBuildPlanStub        func(buildID int, planID atc.PlanID, src io.Reader) (bool, error)
	sendInputToBuildPlanMutex       sync.RWMutex
	sendInputToBuildPlanArgsForCall []struct {
		buildID int
		planID  atc.PlanID
		src     io.Reader
	}
	sendInputToBuildPlanReturns struct {
		result1 bool
		result2 error
	}
	sendInputToBuildPlanReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ReadOutputFromBuildPlanStub        func(buildID int, planID atc.PlanID) (io.ReadCloser, bool, error)
	readOutputFromBuildPlanMutex       sync.RWMutex
	readOutputFromBuildPlanArgsForCall []struct {
		buildID int
		planID  atc.PlanID
	}
	readOutputFromBuildPlanReturns struct {
		result1 io.ReadCloser
		result2 bool
		result3 error
	}
	readOutputFromBuildPlanReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 bool
		result3 error
	}
	SaveWorkerStub        func(atc.Worker, *time.Duration) (*atc.Worker, error)
	saveWorkerMutex       sync.RWMutex
	saveWorkerArgsForCall []struct {
		arg1 atc.Worker
		arg2 *time.Duration
	}
	saveWorkerReturns struct {
		result1 *atc.Worker
		result2 error
	}
	saveWorkerReturnsOnCall map[int]struct {
		result1 *atc.Worker
		result2 error
	}
	ListWorkersStub        func() ([]atc.Worker, error)
	listWorkersMutex       sync.RWMutex
	listWorkersArgsForCall []struct{}
	listWorkersReturns     struct {
		result1 []atc.Worker
		result2 error
	}
	listWorkersReturnsOnCall map[int]struct {
		result1 []atc.Worker
		result2 error
	}
	PruneWorkerStub        func(workerName string) error
	pruneWorkerMutex       sync.RWMutex
	pruneWorkerArgsForCall []struct {
		workerName string
	}
	pruneWorkerReturns struct {
		result1 error
	}
	pruneWorkerReturnsOnCall map[int]struct {
		result1 error
	}
	GetInfoStub        func() (atc.Info, error)
	getInfoMutex       sync.RWMutex
	getInfoArgsForCall []struct{}
	getInfoReturns     struct {
		result1 atc.Info
		result2 error
	}
	getInfoReturnsOnCall map[int]struct {
		result1 atc.Info
		result2 error
	}
	GetCLIReaderStub        func(arch, platform string) (io.ReadCloser, http.Header, error)
	getCLIReaderMutex       sync.RWMutex
	getCLIReaderArgsForCall []struct {
		arch     string
		platform string
	}
	getCLIReaderReturns struct {
		result1 io.ReadCloser
		result2 http.Header
		result3 error
	}
	getCLIReaderReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 http.Header
		result3 error
	}
	ListPipelinesStub        func() ([]atc.Pipeline, error)
	listPipelinesMutex       sync.RWMutex
	listPipelinesArgsForCall []struct{}
	listPipelinesReturns     struct {
		result1 []atc.Pipeline
		result2 error
	}
	listPipelinesReturnsOnCall map[int]struct {
		result1 []atc.Pipeline
		result2 error
	}
	ListTeamsStub        func() ([]atc.Team, error)
	listTeamsMutex       sync.RWMutex
	listTeamsArgsForCall []struct{}
	listTeamsReturns     struct {
		result1 []atc.Team
		result2 error
	}
	listTeamsReturnsOnCall map[int]struct {
		result1 []atc.Team
		result2 error
	}
	TeamStub        func(teamName string) concourse.Team
	teamMutex       sync.RWMutex
	teamArgsForCall []struct {
		teamName string
	}
	teamReturns struct {
		result1 concourse.Team
	}
	teamReturnsOnCall map[int]struct {
		result1 concourse.Team
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) URL() string {
	fake.uRLMutex.Lock()
	ret, specificReturn := fake.uRLReturnsOnCall[len(fake.uRLArgsForCall)]
	fake.uRLArgsForCall = append(fake.uRLArgsForCall, struct{}{})
	fake.recordInvocation("URL", []interface{}{})
	fake.uRLMutex.Unlock()
	if fake.URLStub != nil {
		return fake.URLStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.uRLReturns.result1
}

func (fake *FakeClient) URLCallCount() int {
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	return len(fake.uRLArgsForCall)
}

func (fake *FakeClient) URLReturns(result1 string) {
	fake.URLStub = nil
	fake.uRLReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeClient) URLReturnsOnCall(i int, result1 string) {
	fake.URLStub = nil
	if fake.uRLReturnsOnCall == nil {
		fake.uRLReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.uRLReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeClient) HTTPClient() *http.Client {
	fake.hTTPClientMutex.Lock()
	ret, specificReturn := fake.hTTPClientReturnsOnCall[len(fake.hTTPClientArgsForCall)]
	fake.hTTPClientArgsForCall = append(fake.hTTPClientArgsForCall, struct{}{})
	fake.recordInvocation("HTTPClient", []interface{}{})
	fake.hTTPClientMutex.Unlock()
	if fake.HTTPClientStub != nil {
		return fake.HTTPClientStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.hTTPClientReturns.result1
}

func (fake *FakeClient) HTTPClientCallCount() int {
	fake.hTTPClientMutex.RLock()
	defer fake.hTTPClientMutex.RUnlock()
	return len(fake.hTTPClientArgsForCall)
}

func (fake *FakeClient) HTTPClientReturns(result1 *http.Client) {
	fake.HTTPClientStub = nil
	fake.hTTPClientReturns = struct {
		result1 *http.Client
	}{result1}
}

func (fake *FakeClient) HTTPClientReturnsOnCall(i int, result1 *http.Client) {
	fake.HTTPClientStub = nil
	if fake.hTTPClientReturnsOnCall == nil {
		fake.hTTPClientReturnsOnCall = make(map[int]struct {
			result1 *http.Client
		})
	}
	fake.hTTPClientReturnsOnCall[i] = struct {
		result1 *http.Client
	}{result1}
}

func (fake *FakeClient) Builds(arg1 concourse.Page) ([]atc.Build, concourse.Pagination, error) {
	fake.buildsMutex.Lock()
	ret, specificReturn := fake.buildsReturnsOnCall[len(fake.buildsArgsForCall)]
	fake.buildsArgsForCall = append(fake.buildsArgsForCall, struct {
		arg1 concourse.Page
	}{arg1})
	fake.recordInvocation("Builds", []interface{}{arg1})
	fake.buildsMutex.Unlock()
	if fake.BuildsStub != nil {
		return fake.BuildsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.buildsReturns.result1, fake.buildsReturns.result2, fake.buildsReturns.result3
}

func (fake *FakeClient) BuildsCallCount() int {
	fake.buildsMutex.RLock()
	defer fake.buildsMutex.RUnlock()
	return len(fake.buildsArgsForCall)
}

func (fake *FakeClient) BuildsArgsForCall(i int) concourse.Page {
	fake.buildsMutex.RLock()
	defer fake.buildsMutex.RUnlock()
	return fake.buildsArgsForCall[i].arg1
}

func (fake *FakeClient) BuildsReturns(result1 []atc.Build, result2 concourse.Pagination, result3 error) {
	fake.BuildsStub = nil
	fake.buildsReturns = struct {
		result1 []atc.Build
		result2 concourse.Pagination
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) BuildsReturnsOnCall(i int, result1 []atc.Build, result2 concourse.Pagination, result3 error) {
	fake.BuildsStub = nil
	if fake.buildsReturnsOnCall == nil {
		fake.buildsReturnsOnCall = make(map[int]struct {
			result1 []atc.Build
			result2 concourse.Pagination
			result3 error
		})
	}
	fake.buildsReturnsOnCall[i] = struct {
		result1 []atc.Build
		result2 concourse.Pagination
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) Build(buildID string) (atc.Build, bool, error) {
	fake.buildMutex.Lock()
	ret, specificReturn := fake.buildReturnsOnCall[len(fake.buildArgsForCall)]
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		buildID string
	}{buildID})
	fake.recordInvocation("Build", []interface{}{buildID})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub(buildID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.buildReturns.result1, fake.buildReturns.result2, fake.buildReturns.result3
}

func (fake *FakeClient) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakeClient) BuildArgsForCall(i int) string {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return fake.buildArgsForCall[i].buildID
}

func (fake *FakeClient) BuildReturns(result1 atc.Build, result2 bool, result3 error) {
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 atc.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) BuildReturnsOnCall(i int, result1 atc.Build, result2 bool, result3 error) {
	fake.BuildStub = nil
	if fake.buildReturnsOnCall == nil {
		fake.buildReturnsOnCall = make(map[int]struct {
			result1 atc.Build
			result2 bool
			result3 error
		})
	}
	fake.buildReturnsOnCall[i] = struct {
		result1 atc.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) BuildEvents(buildID string) (concourse.Events, error) {
	fake.buildEventsMutex.Lock()
	ret, specificReturn := fake.buildEventsReturnsOnCall[len(fake.buildEventsArgsForCall)]
	fake.buildEventsArgsForCall = append(fake.buildEventsArgsForCall, struct {
		buildID string
	}{buildID})
	fake.recordInvocation("BuildEvents", []interface{}{buildID})
	fake.buildEventsMutex.Unlock()
	if fake.BuildEventsStub != nil {
		return fake.BuildEventsStub(buildID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.buildEventsReturns.result1, fake.buildEventsReturns.result2
}

func (fake *FakeClient) BuildEventsCallCount() int {
	fake.buildEventsMutex.RLock()
	defer fake.buildEventsMutex.RUnlock()
	return len(fake.buildEventsArgsForCall)
}

func (fake *FakeClient) BuildEventsArgsForCall(i int) string {
	fake.buildEventsMutex.RLock()
	defer fake.buildEventsMutex.RUnlock()
	return fake.buildEventsArgsForCall[i].buildID
}

func (fake *FakeClient) BuildEventsReturns(result1 concourse.Events, result2 error) {
	fake.BuildEventsStub = nil
	fake.buildEventsReturns = struct {
		result1 concourse.Events
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) BuildEventsReturnsOnCall(i int, result1 concourse.Events, result2 error) {
	fake.BuildEventsStub = nil
	if fake.buildEventsReturnsOnCall == nil {
		fake.buildEventsReturnsOnCall = make(map[int]struct {
			result1 concourse.Events
			result2 error
		})
	}
	fake.buildEventsReturnsOnCall[i] = struct {
		result1 concourse.Events
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) BuildResources(buildID int) (atc.BuildInputsOutputs, bool, error) {
	fake.buildResourcesMutex.Lock()
	ret, specificReturn := fake.buildResourcesReturnsOnCall[len(fake.buildResourcesArgsForCall)]
	fake.buildResourcesArgsForCall = append(fake.buildResourcesArgsForCall, struct {
		buildID int
	}{buildID})
	fake.recordInvocation("BuildResources", []interface{}{buildID})
	fake.buildResourcesMutex.Unlock()
	if fake.BuildResourcesStub != nil {
		return fake.BuildResourcesStub(buildID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.buildResourcesReturns.result1, fake.buildResourcesReturns.result2, fake.buildResourcesReturns.result3
}

func (fake *FakeClient) BuildResourcesCallCount() int {
	fake.buildResourcesMutex.RLock()
	defer fake.buildResourcesMutex.RUnlock()
	return len(fake.buildResourcesArgsForCall)
}

func (fake *FakeClient) BuildResourcesArgsForCall(i int) int {
	fake.buildResourcesMutex.RLock()
	defer fake.buildResourcesMutex.RUnlock()
	return fake.buildResourcesArgsForCall[i].buildID
}

func (fake *FakeClient) BuildResourcesReturns(result1 atc.BuildInputsOutputs, result2 bool, result3 error) {
	fake.BuildResourcesStub = nil
	fake.buildResourcesReturns = struct {
		result1 atc.BuildInputsOutputs
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) BuildResourcesReturnsOnCall(i int, result1 atc.BuildInputsOutputs, result2 bool, result3 error) {
	fake.BuildResourcesStub = nil
	if fake.buildResourcesReturnsOnCall == nil {
		fake.buildResourcesReturnsOnCall = make(map[int]struct {
			result1 atc.BuildInputsOutputs
			result2 bool
			result3 error
		})
	}
	fake.buildResourcesReturnsOnCall[i] = struct {
		result1 atc.BuildInputsOutputs
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) AbortBuild(buildID string) error {
	fake.abortBuildMutex.Lock()
	ret, specificReturn := fake.abortBuildReturnsOnCall[len(fake.abortBuildArgsForCall)]
	fake.abortBuildArgsForCall = append(fake.abortBuildArgsForCall, struct {
		buildID string
	}{buildID})
	fake.recordInvocation("AbortBuild", []interface{}{buildID})
	fake.abortBuildMutex.Unlock()
	if fake.AbortBuildStub != nil {
		return fake.AbortBuildStub(buildID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.abortBuildReturns.result1
}

func (fake *FakeClient) AbortBuildCallCount() int {
	fake.abortBuildMutex.RLock()
	defer fake.abortBuildMutex.RUnlock()
	return len(fake.abortBuildArgsForCall)
}

func (fake *FakeClient) AbortBuildArgsForCall(i int) string {
	fake.abortBuildMutex.RLock()
	defer fake.abortBuildMutex.RUnlock()
	return fake.abortBuildArgsForCall[i].buildID
}

func (fake *FakeClient) AbortBuildReturns(result1 error) {
	fake.AbortBuildStub = nil
	fake.abortBuildReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) AbortBuildReturnsOnCall(i int, result1 error) {
	fake.AbortBuildStub = nil
	if fake.abortBuildReturnsOnCall == nil {
		fake.abortBuildReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.abortBuildReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) BuildPlan(buildID int) (atc.PublicBuildPlan, bool, error) {
	fake.buildPlanMutex.Lock()
	ret, specificReturn := fake.buildPlanReturnsOnCall[len(fake.buildPlanArgsForCall)]
	fake.buildPlanArgsForCall = append(fake.buildPlanArgsForCall, struct {
		buildID int
	}{buildID})
	fake.recordInvocation("BuildPlan", []interface{}{buildID})
	fake.buildPlanMutex.Unlock()
	if fake.BuildPlanStub != nil {
		return fake.BuildPlanStub(buildID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.buildPlanReturns.result1, fake.buildPlanReturns.result2, fake.buildPlanReturns.result3
}

func (fake *FakeClient) BuildPlanCallCount() int {
	fake.buildPlanMutex.RLock()
	defer fake.buildPlanMutex.RUnlock()
	return len(fake.buildPlanArgsForCall)
}

func (fake *FakeClient) BuildPlanArgsForCall(i int) int {
	fake.buildPlanMutex.RLock()
	defer fake.buildPlanMutex.RUnlock()
	return fake.buildPlanArgsForCall[i].buildID
}

func (fake *FakeClient) BuildPlanReturns(result1 atc.PublicBuildPlan, result2 bool, result3 error) {
	fake.BuildPlanStub = nil
	fake.buildPlanReturns = struct {
		result1 atc.PublicBuildPlan
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) BuildPlanReturnsOnCall(i int, result1 atc.PublicBuildPlan, result2 bool, result3 error) {
	fake.BuildPlanStub = nil
	if fake.buildPlanReturnsOnCall == nil {
		fake.buildPlanReturnsOnCall = make(map[int]struct {
			result1 atc.PublicBuildPlan
			result2 bool
			result3 error
		})
	}
	fake.buildPlanReturnsOnCall[i] = struct {
		result1 atc.PublicBuildPlan
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) SendInputToBuildPlan(buildID int, planID atc.PlanID, src io.Reader) (bool, error) {
	fake.sendInputToBuildPlanMutex.Lock()
	ret, specificReturn := fake.sendInputToBuildPlanReturnsOnCall[len(fake.sendInputToBuildPlanArgsForCall)]
	fake.sendInputToBuildPlanArgsForCall = append(fake.sendInputToBuildPlanArgsForCall, struct {
		buildID int
		planID  atc.PlanID
		src     io.Reader
	}{buildID, planID, src})
	fake.recordInvocation("SendInputToBuildPlan", []interface{}{buildID, planID, src})
	fake.sendInputToBuildPlanMutex.Unlock()
	if fake.SendInputToBuildPlanStub != nil {
		return fake.SendInputToBuildPlanStub(buildID, planID, src)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.sendInputToBuildPlanReturns.result1, fake.sendInputToBuildPlanReturns.result2
}

func (fake *FakeClient) SendInputToBuildPlanCallCount() int {
	fake.sendInputToBuildPlanMutex.RLock()
	defer fake.sendInputToBuildPlanMutex.RUnlock()
	return len(fake.sendInputToBuildPlanArgsForCall)
}

func (fake *FakeClient) SendInputToBuildPlanArgsForCall(i int) (int, atc.PlanID, io.Reader) {
	fake.sendInputToBuildPlanMutex.RLock()
	defer fake.sendInputToBuildPlanMutex.RUnlock()
	return fake.sendInputToBuildPlanArgsForCall[i].buildID, fake.sendInputToBuildPlanArgsForCall[i].planID, fake.sendInputToBuildPlanArgsForCall[i].src
}

func (fake *FakeClient) SendInputToBuildPlanReturns(result1 bool, result2 error) {
	fake.SendInputToBuildPlanStub = nil
	fake.sendInputToBuildPlanReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) SendInputToBuildPlanReturnsOnCall(i int, result1 bool, result2 error) {
	fake.SendInputToBuildPlanStub = nil
	if fake.sendInputToBuildPlanReturnsOnCall == nil {
		fake.sendInputToBuildPlanReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.sendInputToBuildPlanReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ReadOutputFromBuildPlan(buildID int, planID atc.PlanID) (io.ReadCloser, bool, error) {
	fake.readOutputFromBuildPlanMutex.Lock()
	ret, specificReturn := fake.readOutputFromBuildPlanReturnsOnCall[len(fake.readOutputFromBuildPlanArgsForCall)]
	fake.readOutputFromBuildPlanArgsForCall = append(fake.readOutputFromBuildPlanArgsForCall, struct {
		buildID int
		planID  atc.PlanID
	}{buildID, planID})
	fake.recordInvocation("ReadOutputFromBuildPlan", []interface{}{buildID, planID})
	fake.readOutputFromBuildPlanMutex.Unlock()
	if fake.ReadOutputFromBuildPlanStub != nil {
		return fake.ReadOutputFromBuildPlanStub(buildID, planID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.readOutputFromBuildPlanReturns.result1, fake.readOutputFromBuildPlanReturns.result2, fake.readOutputFromBuildPlanReturns.result3
}

func (fake *FakeClient) ReadOutputFromBuildPlanCallCount() int {
	fake.readOutputFromBuildPlanMutex.RLock()
	defer fake.readOutputFromBuildPlanMutex.RUnlock()
	return len(fake.readOutputFromBuildPlanArgsForCall)
}

func (fake *FakeClient) ReadOutputFromBuildPlanArgsForCall(i int) (int, atc.PlanID) {
	fake.readOutputFromBuildPlanMutex.RLock()
	defer fake.readOutputFromBuildPlanMutex.RUnlock()
	return fake.readOutputFromBuildPlanArgsForCall[i].buildID, fake.readOutputFromBuildPlanArgsForCall[i].planID
}

func (fake *FakeClient) ReadOutputFromBuildPlanReturns(result1 io.ReadCloser, result2 bool, result3 error) {
	fake.ReadOutputFromBuildPlanStub = nil
	fake.readOutputFromBuildPlanReturns = struct {
		result1 io.ReadCloser
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ReadOutputFromBuildPlanReturnsOnCall(i int, result1 io.ReadCloser, result2 bool, result3 error) {
	fake.ReadOutputFromBuildPlanStub = nil
	if fake.readOutputFromBuildPlanReturnsOnCall == nil {
		fake.readOutputFromBuildPlanReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 bool
			result3 error
		})
	}
	fake.readOutputFromBuildPlanReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) SaveWorker(arg1 atc.Worker, arg2 *time.Duration) (*atc.Worker, error) {
	fake.saveWorkerMutex.Lock()
	ret, specificReturn := fake.saveWorkerReturnsOnCall[len(fake.saveWorkerArgsForCall)]
	fake.saveWorkerArgsForCall = append(fake.saveWorkerArgsForCall, struct {
		arg1 atc.Worker
		arg2 *time.Duration
	}{arg1, arg2})
	fake.recordInvocation("SaveWorker", []interface{}{arg1, arg2})
	fake.saveWorkerMutex.Unlock()
	if fake.SaveWorkerStub != nil {
		return fake.SaveWorkerStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.saveWorkerReturns.result1, fake.saveWorkerReturns.result2
}

func (fake *FakeClient) SaveWorkerCallCount() int {
	fake.saveWorkerMutex.RLock()
	defer fake.saveWorkerMutex.RUnlock()
	return len(fake.saveWorkerArgsForCall)
}

func (fake *FakeClient) SaveWorkerArgsForCall(i int) (atc.Worker, *time.Duration) {
	fake.saveWorkerMutex.RLock()
	defer fake.saveWorkerMutex.RUnlock()
	return fake.saveWorkerArgsForCall[i].arg1, fake.saveWorkerArgsForCall[i].arg2
}

func (fake *FakeClient) SaveWorkerReturns(result1 *atc.Worker, result2 error) {
	fake.SaveWorkerStub = nil
	fake.saveWorkerReturns = struct {
		result1 *atc.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) SaveWorkerReturnsOnCall(i int, result1 *atc.Worker, result2 error) {
	fake.SaveWorkerStub = nil
	if fake.saveWorkerReturnsOnCall == nil {
		fake.saveWorkerReturnsOnCall = make(map[int]struct {
			result1 *atc.Worker
			result2 error
		})
	}
	fake.saveWorkerReturnsOnCall[i] = struct {
		result1 *atc.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListWorkers() ([]atc.Worker, error) {
	fake.listWorkersMutex.Lock()
	ret, specificReturn := fake.listWorkersReturnsOnCall[len(fake.listWorkersArgsForCall)]
	fake.listWorkersArgsForCall = append(fake.listWorkersArgsForCall, struct{}{})
	fake.recordInvocation("ListWorkers", []interface{}{})
	fake.listWorkersMutex.Unlock()
	if fake.ListWorkersStub != nil {
		return fake.ListWorkersStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listWorkersReturns.result1, fake.listWorkersReturns.result2
}

func (fake *FakeClient) ListWorkersCallCount() int {
	fake.listWorkersMutex.RLock()
	defer fake.listWorkersMutex.RUnlock()
	return len(fake.listWorkersArgsForCall)
}

func (fake *FakeClient) ListWorkersReturns(result1 []atc.Worker, result2 error) {
	fake.ListWorkersStub = nil
	fake.listWorkersReturns = struct {
		result1 []atc.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListWorkersReturnsOnCall(i int, result1 []atc.Worker, result2 error) {
	fake.ListWorkersStub = nil
	if fake.listWorkersReturnsOnCall == nil {
		fake.listWorkersReturnsOnCall = make(map[int]struct {
			result1 []atc.Worker
			result2 error
		})
	}
	fake.listWorkersReturnsOnCall[i] = struct {
		result1 []atc.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) PruneWorker(workerName string) error {
	fake.pruneWorkerMutex.Lock()
	ret, specificReturn := fake.pruneWorkerReturnsOnCall[len(fake.pruneWorkerArgsForCall)]
	fake.pruneWorkerArgsForCall = append(fake.pruneWorkerArgsForCall, struct {
		workerName string
	}{workerName})
	fake.recordInvocation("PruneWorker", []interface{}{workerName})
	fake.pruneWorkerMutex.Unlock()
	if fake.PruneWorkerStub != nil {
		return fake.PruneWorkerStub(workerName)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pruneWorkerReturns.result1
}

func (fake *FakeClient) PruneWorkerCallCount() int {
	fake.pruneWorkerMutex.RLock()
	defer fake.pruneWorkerMutex.RUnlock()
	return len(fake.pruneWorkerArgsForCall)
}

func (fake *FakeClient) PruneWorkerArgsForCall(i int) string {
	fake.pruneWorkerMutex.RLock()
	defer fake.pruneWorkerMutex.RUnlock()
	return fake.pruneWorkerArgsForCall[i].workerName
}

func (fake *FakeClient) PruneWorkerReturns(result1 error) {
	fake.PruneWorkerStub = nil
	fake.pruneWorkerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) PruneWorkerReturnsOnCall(i int, result1 error) {
	fake.PruneWorkerStub = nil
	if fake.pruneWorkerReturnsOnCall == nil {
		fake.pruneWorkerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pruneWorkerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) GetInfo() (atc.Info, error) {
	fake.getInfoMutex.Lock()
	ret, specificReturn := fake.getInfoReturnsOnCall[len(fake.getInfoArgsForCall)]
	fake.getInfoArgsForCall = append(fake.getInfoArgsForCall, struct{}{})
	fake.recordInvocation("GetInfo", []interface{}{})
	fake.getInfoMutex.Unlock()
	if fake.GetInfoStub != nil {
		return fake.GetInfoStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getInfoReturns.result1, fake.getInfoReturns.result2
}

func (fake *FakeClient) GetInfoCallCount() int {
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	return len(fake.getInfoArgsForCall)
}

func (fake *FakeClient) GetInfoReturns(result1 atc.Info, result2 error) {
	fake.GetInfoStub = nil
	fake.getInfoReturns = struct {
		result1 atc.Info
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetInfoReturnsOnCall(i int, result1 atc.Info, result2 error) {
	fake.GetInfoStub = nil
	if fake.getInfoReturnsOnCall == nil {
		fake.getInfoReturnsOnCall = make(map[int]struct {
			result1 atc.Info
			result2 error
		})
	}
	fake.getInfoReturnsOnCall[i] = struct {
		result1 atc.Info
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetCLIReader(arch string, platform string) (io.ReadCloser, http.Header, error) {
	fake.getCLIReaderMutex.Lock()
	ret, specificReturn := fake.getCLIReaderReturnsOnCall[len(fake.getCLIReaderArgsForCall)]
	fake.getCLIReaderArgsForCall = append(fake.getCLIReaderArgsForCall, struct {
		arch     string
		platform string
	}{arch, platform})
	fake.recordInvocation("GetCLIReader", []interface{}{arch, platform})
	fake.getCLIReaderMutex.Unlock()
	if fake.GetCLIReaderStub != nil {
		return fake.GetCLIReaderStub(arch, platform)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getCLIReaderReturns.result1, fake.getCLIReaderReturns.result2, fake.getCLIReaderReturns.result3
}

func (fake *FakeClient) GetCLIReaderCallCount() int {
	fake.getCLIReaderMutex.RLock()
	defer fake.getCLIReaderMutex.RUnlock()
	return len(fake.getCLIReaderArgsForCall)
}

func (fake *FakeClient) GetCLIReaderArgsForCall(i int) (string, string) {
	fake.getCLIReaderMutex.RLock()
	defer fake.getCLIReaderMutex.RUnlock()
	return fake.getCLIReaderArgsForCall[i].arch, fake.getCLIReaderArgsForCall[i].platform
}

func (fake *FakeClient) GetCLIReaderReturns(result1 io.ReadCloser, result2 http.Header, result3 error) {
	fake.GetCLIReaderStub = nil
	fake.getCLIReaderReturns = struct {
		result1 io.ReadCloser
		result2 http.Header
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) GetCLIReaderReturnsOnCall(i int, result1 io.ReadCloser, result2 http.Header, result3 error) {
	fake.GetCLIReaderStub = nil
	if fake.getCLIReaderReturnsOnCall == nil {
		fake.getCLIReaderReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 http.Header
			result3 error
		})
	}
	fake.getCLIReaderReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 http.Header
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListPipelines() ([]atc.Pipeline, error) {
	fake.listPipelinesMutex.Lock()
	ret, specificReturn := fake.listPipelinesReturnsOnCall[len(fake.listPipelinesArgsForCall)]
	fake.listPipelinesArgsForCall = append(fake.listPipelinesArgsForCall, struct{}{})
	fake.recordInvocation("ListPipelines", []interface{}{})
	fake.listPipelinesMutex.Unlock()
	if fake.ListPipelinesStub != nil {
		return fake.ListPipelinesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listPipelinesReturns.result1, fake.listPipelinesReturns.result2
}

func (fake *FakeClient) ListPipelinesCallCount() int {
	fake.listPipelinesMutex.RLock()
	defer fake.listPipelinesMutex.RUnlock()
	return len(fake.listPipelinesArgsForCall)
}

func (fake *FakeClient) ListPipelinesReturns(result1 []atc.Pipeline, result2 error) {
	fake.ListPipelinesStub = nil
	fake.listPipelinesReturns = struct {
		result1 []atc.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListPipelinesReturnsOnCall(i int, result1 []atc.Pipeline, result2 error) {
	fake.ListPipelinesStub = nil
	if fake.listPipelinesReturnsOnCall == nil {
		fake.listPipelinesReturnsOnCall = make(map[int]struct {
			result1 []atc.Pipeline
			result2 error
		})
	}
	fake.listPipelinesReturnsOnCall[i] = struct {
		result1 []atc.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListTeams() ([]atc.Team, error) {
	fake.listTeamsMutex.Lock()
	ret, specificReturn := fake.listTeamsReturnsOnCall[len(fake.listTeamsArgsForCall)]
	fake.listTeamsArgsForCall = append(fake.listTeamsArgsForCall, struct{}{})
	fake.recordInvocation("ListTeams", []interface{}{})
	fake.listTeamsMutex.Unlock()
	if fake.ListTeamsStub != nil {
		return fake.ListTeamsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listTeamsReturns.result1, fake.listTeamsReturns.result2
}

func (fake *FakeClient) ListTeamsCallCount() int {
	fake.listTeamsMutex.RLock()
	defer fake.listTeamsMutex.RUnlock()
	return len(fake.listTeamsArgsForCall)
}

func (fake *FakeClient) ListTeamsReturns(result1 []atc.Team, result2 error) {
	fake.ListTeamsStub = nil
	fake.listTeamsReturns = struct {
		result1 []atc.Team
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListTeamsReturnsOnCall(i int, result1 []atc.Team, result2 error) {
	fake.ListTeamsStub = nil
	if fake.listTeamsReturnsOnCall == nil {
		fake.listTeamsReturnsOnCall = make(map[int]struct {
			result1 []atc.Team
			result2 error
		})
	}
	fake.listTeamsReturnsOnCall[i] = struct {
		result1 []atc.Team
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Team(teamName string) concourse.Team {
	fake.teamMutex.Lock()
	ret, specificReturn := fake.teamReturnsOnCall[len(fake.teamArgsForCall)]
	fake.teamArgsForCall = append(fake.teamArgsForCall, struct {
		teamName string
	}{teamName})
	fake.recordInvocation("Team", []interface{}{teamName})
	fake.teamMutex.Unlock()
	if fake.TeamStub != nil {
		return fake.TeamStub(teamName)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.teamReturns.result1
}

func (fake *FakeClient) TeamCallCount() int {
	fake.teamMutex.RLock()
	defer fake.teamMutex.RUnlock()
	return len(fake.teamArgsForCall)
}

func (fake *FakeClient) TeamArgsForCall(i int) string {
	fake.teamMutex.RLock()
	defer fake.teamMutex.RUnlock()
	return fake.teamArgsForCall[i].teamName
}

func (fake *FakeClient) TeamReturns(result1 concourse.Team) {
	fake.TeamStub = nil
	fake.teamReturns = struct {
		result1 concourse.Team
	}{result1}
}

func (fake *FakeClient) TeamReturnsOnCall(i int, result1 concourse.Team) {
	fake.TeamStub = nil
	if fake.teamReturnsOnCall == nil {
		fake.teamReturnsOnCall = make(map[int]struct {
			result1 concourse.Team
		})
	}
	fake.teamReturnsOnCall[i] = struct {
		result1 concourse.Team
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	fake.hTTPClientMutex.RLock()
	defer fake.hTTPClientMutex.RUnlock()
	fake.buildsMutex.RLock()
	defer fake.buildsMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	fake.buildEventsMutex.RLock()
	defer fake.buildEventsMutex.RUnlock()
	fake.buildResourcesMutex.RLock()
	defer fake.buildResourcesMutex.RUnlock()
	fake.abortBuildMutex.RLock()
	defer fake.abortBuildMutex.RUnlock()
	fake.buildPlanMutex.RLock()
	defer fake.buildPlanMutex.RUnlock()
	fake.sendInputToBuildPlanMutex.RLock()
	defer fake.sendInputToBuildPlanMutex.RUnlock()
	fake.readOutputFromBuildPlanMutex.RLock()
	defer fake.readOutputFromBuildPlanMutex.RUnlock()
	fake.saveWorkerMutex.RLock()
	defer fake.saveWorkerMutex.RUnlock()
	fake.listWorkersMutex.RLock()
	defer fake.listWorkersMutex.RUnlock()
	fake.pruneWorkerMutex.RLock()
	defer fake.pruneWorkerMutex.RUnlock()
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	fake.getCLIReaderMutex.RLock()
	defer fake.getCLIReaderMutex.RUnlock()
	fake.listPipelinesMutex.RLock()
	defer fake.listPipelinesMutex.RUnlock()
	fake.listTeamsMutex.RLock()
	defer fake.listTeamsMutex.RUnlock()
	fake.teamMutex.RLock()
	defer fake.teamMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ concourse.Client = new(FakeClient)
