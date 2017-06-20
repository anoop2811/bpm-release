// This file was generated by counterfeiter
package runcadapterfakes

import (
	"crucible/config"
	"crucible/runcadapter"
	"io"
	"os"
	"sync"

	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type FakeRuncAdapter struct {
	CreateJobPrerequisitesStub        func(systemRoot, jobName string) (string, *os.File, *os.File, error)
	createJobPrerequisitesMutex       sync.RWMutex
	createJobPrerequisitesArgsForCall []struct {
		systemRoot string
		jobName    string
	}
	createJobPrerequisitesReturns struct {
		result1 string
		result2 *os.File
		result3 *os.File
		result4 error
	}
	createJobPrerequisitesReturnsOnCall map[int]struct {
		result1 string
		result2 *os.File
		result3 *os.File
		result4 error
	}
	BuildSpecStub        func(jobName string, jobConfig *config.CrucibleConfig) (specs.Spec, error)
	buildSpecMutex       sync.RWMutex
	buildSpecArgsForCall []struct {
		jobName   string
		jobConfig *config.CrucibleConfig
	}
	buildSpecReturns struct {
		result1 specs.Spec
		result2 error
	}
	buildSpecReturnsOnCall map[int]struct {
		result1 specs.Spec
		result2 error
	}
	CreateBundleStub        func(bundlesRoot, jobName string, jobSpec specs.Spec) (string, error)
	createBundleMutex       sync.RWMutex
	createBundleArgsForCall []struct {
		bundlesRoot string
		jobName     string
		jobSpec     specs.Spec
	}
	createBundleReturns struct {
		result1 string
		result2 error
	}
	createBundleReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	RunContainerStub        func(pidDir, bundlePath, jobName string, stdout, stderr io.Writer) error
	runContainerMutex       sync.RWMutex
	runContainerArgsForCall []struct {
		pidDir     string
		bundlePath string
		jobName    string
		stdout     io.Writer
		stderr     io.Writer
	}
	runContainerReturns struct {
		result1 error
	}
	runContainerReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteContainerStub        func(jobName string) error
	deleteContainerMutex       sync.RWMutex
	deleteContainerArgsForCall []struct {
		jobName string
	}
	deleteContainerReturns struct {
		result1 error
	}
	deleteContainerReturnsOnCall map[int]struct {
		result1 error
	}
	DestroyBundleStub        func(bundlesRoot, jobName string) error
	destroyBundleMutex       sync.RWMutex
	destroyBundleArgsForCall []struct {
		bundlesRoot string
		jobName     string
	}
	destroyBundleReturns struct {
		result1 error
	}
	destroyBundleReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRuncAdapter) CreateJobPrerequisites(systemRoot string, jobName string) (string, *os.File, *os.File, error) {
	fake.createJobPrerequisitesMutex.Lock()
	ret, specificReturn := fake.createJobPrerequisitesReturnsOnCall[len(fake.createJobPrerequisitesArgsForCall)]
	fake.createJobPrerequisitesArgsForCall = append(fake.createJobPrerequisitesArgsForCall, struct {
		systemRoot string
		jobName    string
	}{systemRoot, jobName})
	fake.recordInvocation("CreateJobPrerequisites", []interface{}{systemRoot, jobName})
	fake.createJobPrerequisitesMutex.Unlock()
	if fake.CreateJobPrerequisitesStub != nil {
		return fake.CreateJobPrerequisitesStub(systemRoot, jobName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.createJobPrerequisitesReturns.result1, fake.createJobPrerequisitesReturns.result2, fake.createJobPrerequisitesReturns.result3, fake.createJobPrerequisitesReturns.result4
}

func (fake *FakeRuncAdapter) CreateJobPrerequisitesCallCount() int {
	fake.createJobPrerequisitesMutex.RLock()
	defer fake.createJobPrerequisitesMutex.RUnlock()
	return len(fake.createJobPrerequisitesArgsForCall)
}

func (fake *FakeRuncAdapter) CreateJobPrerequisitesArgsForCall(i int) (string, string) {
	fake.createJobPrerequisitesMutex.RLock()
	defer fake.createJobPrerequisitesMutex.RUnlock()
	return fake.createJobPrerequisitesArgsForCall[i].systemRoot, fake.createJobPrerequisitesArgsForCall[i].jobName
}

func (fake *FakeRuncAdapter) CreateJobPrerequisitesReturns(result1 string, result2 *os.File, result3 *os.File, result4 error) {
	fake.CreateJobPrerequisitesStub = nil
	fake.createJobPrerequisitesReturns = struct {
		result1 string
		result2 *os.File
		result3 *os.File
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeRuncAdapter) CreateJobPrerequisitesReturnsOnCall(i int, result1 string, result2 *os.File, result3 *os.File, result4 error) {
	fake.CreateJobPrerequisitesStub = nil
	if fake.createJobPrerequisitesReturnsOnCall == nil {
		fake.createJobPrerequisitesReturnsOnCall = make(map[int]struct {
			result1 string
			result2 *os.File
			result3 *os.File
			result4 error
		})
	}
	fake.createJobPrerequisitesReturnsOnCall[i] = struct {
		result1 string
		result2 *os.File
		result3 *os.File
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeRuncAdapter) BuildSpec(jobName string, jobConfig *config.CrucibleConfig) (specs.Spec, error) {
	fake.buildSpecMutex.Lock()
	ret, specificReturn := fake.buildSpecReturnsOnCall[len(fake.buildSpecArgsForCall)]
	fake.buildSpecArgsForCall = append(fake.buildSpecArgsForCall, struct {
		jobName   string
		jobConfig *config.CrucibleConfig
	}{jobName, jobConfig})
	fake.recordInvocation("BuildSpec", []interface{}{jobName, jobConfig})
	fake.buildSpecMutex.Unlock()
	if fake.BuildSpecStub != nil {
		return fake.BuildSpecStub(jobName, jobConfig)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.buildSpecReturns.result1, fake.buildSpecReturns.result2
}

func (fake *FakeRuncAdapter) BuildSpecCallCount() int {
	fake.buildSpecMutex.RLock()
	defer fake.buildSpecMutex.RUnlock()
	return len(fake.buildSpecArgsForCall)
}

func (fake *FakeRuncAdapter) BuildSpecArgsForCall(i int) (string, *config.CrucibleConfig) {
	fake.buildSpecMutex.RLock()
	defer fake.buildSpecMutex.RUnlock()
	return fake.buildSpecArgsForCall[i].jobName, fake.buildSpecArgsForCall[i].jobConfig
}

func (fake *FakeRuncAdapter) BuildSpecReturns(result1 specs.Spec, result2 error) {
	fake.BuildSpecStub = nil
	fake.buildSpecReturns = struct {
		result1 specs.Spec
		result2 error
	}{result1, result2}
}

func (fake *FakeRuncAdapter) BuildSpecReturnsOnCall(i int, result1 specs.Spec, result2 error) {
	fake.BuildSpecStub = nil
	if fake.buildSpecReturnsOnCall == nil {
		fake.buildSpecReturnsOnCall = make(map[int]struct {
			result1 specs.Spec
			result2 error
		})
	}
	fake.buildSpecReturnsOnCall[i] = struct {
		result1 specs.Spec
		result2 error
	}{result1, result2}
}

func (fake *FakeRuncAdapter) CreateBundle(bundlesRoot string, jobName string, jobSpec specs.Spec) (string, error) {
	fake.createBundleMutex.Lock()
	ret, specificReturn := fake.createBundleReturnsOnCall[len(fake.createBundleArgsForCall)]
	fake.createBundleArgsForCall = append(fake.createBundleArgsForCall, struct {
		bundlesRoot string
		jobName     string
		jobSpec     specs.Spec
	}{bundlesRoot, jobName, jobSpec})
	fake.recordInvocation("CreateBundle", []interface{}{bundlesRoot, jobName, jobSpec})
	fake.createBundleMutex.Unlock()
	if fake.CreateBundleStub != nil {
		return fake.CreateBundleStub(bundlesRoot, jobName, jobSpec)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createBundleReturns.result1, fake.createBundleReturns.result2
}

func (fake *FakeRuncAdapter) CreateBundleCallCount() int {
	fake.createBundleMutex.RLock()
	defer fake.createBundleMutex.RUnlock()
	return len(fake.createBundleArgsForCall)
}

func (fake *FakeRuncAdapter) CreateBundleArgsForCall(i int) (string, string, specs.Spec) {
	fake.createBundleMutex.RLock()
	defer fake.createBundleMutex.RUnlock()
	return fake.createBundleArgsForCall[i].bundlesRoot, fake.createBundleArgsForCall[i].jobName, fake.createBundleArgsForCall[i].jobSpec
}

func (fake *FakeRuncAdapter) CreateBundleReturns(result1 string, result2 error) {
	fake.CreateBundleStub = nil
	fake.createBundleReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRuncAdapter) CreateBundleReturnsOnCall(i int, result1 string, result2 error) {
	fake.CreateBundleStub = nil
	if fake.createBundleReturnsOnCall == nil {
		fake.createBundleReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createBundleReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRuncAdapter) RunContainer(pidDir string, bundlePath string, jobName string, stdout io.Writer, stderr io.Writer) error {
	fake.runContainerMutex.Lock()
	ret, specificReturn := fake.runContainerReturnsOnCall[len(fake.runContainerArgsForCall)]
	fake.runContainerArgsForCall = append(fake.runContainerArgsForCall, struct {
		pidDir     string
		bundlePath string
		jobName    string
		stdout     io.Writer
		stderr     io.Writer
	}{pidDir, bundlePath, jobName, stdout, stderr})
	fake.recordInvocation("RunContainer", []interface{}{pidDir, bundlePath, jobName, stdout, stderr})
	fake.runContainerMutex.Unlock()
	if fake.RunContainerStub != nil {
		return fake.RunContainerStub(pidDir, bundlePath, jobName, stdout, stderr)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.runContainerReturns.result1
}

func (fake *FakeRuncAdapter) RunContainerCallCount() int {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return len(fake.runContainerArgsForCall)
}

func (fake *FakeRuncAdapter) RunContainerArgsForCall(i int) (string, string, string, io.Writer, io.Writer) {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return fake.runContainerArgsForCall[i].pidDir, fake.runContainerArgsForCall[i].bundlePath, fake.runContainerArgsForCall[i].jobName, fake.runContainerArgsForCall[i].stdout, fake.runContainerArgsForCall[i].stderr
}

func (fake *FakeRuncAdapter) RunContainerReturns(result1 error) {
	fake.RunContainerStub = nil
	fake.runContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuncAdapter) RunContainerReturnsOnCall(i int, result1 error) {
	fake.RunContainerStub = nil
	if fake.runContainerReturnsOnCall == nil {
		fake.runContainerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runContainerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuncAdapter) DeleteContainer(jobName string) error {
	fake.deleteContainerMutex.Lock()
	ret, specificReturn := fake.deleteContainerReturnsOnCall[len(fake.deleteContainerArgsForCall)]
	fake.deleteContainerArgsForCall = append(fake.deleteContainerArgsForCall, struct {
		jobName string
	}{jobName})
	fake.recordInvocation("DeleteContainer", []interface{}{jobName})
	fake.deleteContainerMutex.Unlock()
	if fake.DeleteContainerStub != nil {
		return fake.DeleteContainerStub(jobName)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteContainerReturns.result1
}

func (fake *FakeRuncAdapter) DeleteContainerCallCount() int {
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	return len(fake.deleteContainerArgsForCall)
}

func (fake *FakeRuncAdapter) DeleteContainerArgsForCall(i int) string {
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	return fake.deleteContainerArgsForCall[i].jobName
}

func (fake *FakeRuncAdapter) DeleteContainerReturns(result1 error) {
	fake.DeleteContainerStub = nil
	fake.deleteContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuncAdapter) DeleteContainerReturnsOnCall(i int, result1 error) {
	fake.DeleteContainerStub = nil
	if fake.deleteContainerReturnsOnCall == nil {
		fake.deleteContainerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteContainerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuncAdapter) DestroyBundle(bundlesRoot string, jobName string) error {
	fake.destroyBundleMutex.Lock()
	ret, specificReturn := fake.destroyBundleReturnsOnCall[len(fake.destroyBundleArgsForCall)]
	fake.destroyBundleArgsForCall = append(fake.destroyBundleArgsForCall, struct {
		bundlesRoot string
		jobName     string
	}{bundlesRoot, jobName})
	fake.recordInvocation("DestroyBundle", []interface{}{bundlesRoot, jobName})
	fake.destroyBundleMutex.Unlock()
	if fake.DestroyBundleStub != nil {
		return fake.DestroyBundleStub(bundlesRoot, jobName)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.destroyBundleReturns.result1
}

func (fake *FakeRuncAdapter) DestroyBundleCallCount() int {
	fake.destroyBundleMutex.RLock()
	defer fake.destroyBundleMutex.RUnlock()
	return len(fake.destroyBundleArgsForCall)
}

func (fake *FakeRuncAdapter) DestroyBundleArgsForCall(i int) (string, string) {
	fake.destroyBundleMutex.RLock()
	defer fake.destroyBundleMutex.RUnlock()
	return fake.destroyBundleArgsForCall[i].bundlesRoot, fake.destroyBundleArgsForCall[i].jobName
}

func (fake *FakeRuncAdapter) DestroyBundleReturns(result1 error) {
	fake.DestroyBundleStub = nil
	fake.destroyBundleReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuncAdapter) DestroyBundleReturnsOnCall(i int, result1 error) {
	fake.DestroyBundleStub = nil
	if fake.destroyBundleReturnsOnCall == nil {
		fake.destroyBundleReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyBundleReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuncAdapter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createJobPrerequisitesMutex.RLock()
	defer fake.createJobPrerequisitesMutex.RUnlock()
	fake.buildSpecMutex.RLock()
	defer fake.buildSpecMutex.RUnlock()
	fake.createBundleMutex.RLock()
	defer fake.createBundleMutex.RUnlock()
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	fake.destroyBundleMutex.RLock()
	defer fake.destroyBundleMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRuncAdapter) recordInvocation(key string, args []interface{}) {
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

var _ runcadapter.RuncAdapter = new(FakeRuncAdapter)
