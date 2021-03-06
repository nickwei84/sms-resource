// This file was generated by counterfeiter
package applicationfakes

import (
	"sync"

	"github.com/nickwei84/sms-resource/out/application"
)

type FakeSMSService struct {
	CreateTopicStub        func(topic string) (string, error)
	createTopicMutex       sync.RWMutex
	createTopicArgsForCall []struct {
		topic string
	}
	createTopicReturns struct {
		result1 string
		result2 error
	}
	GetExistingSubscribersStub        func(topicID string) ([]string, error)
	getExistingSubscribersMutex       sync.RWMutex
	getExistingSubscribersArgsForCall []struct {
		topicID string
	}
	getExistingSubscribersReturns struct {
		result1 []string
		result2 error
	}
	CreateNewSubscriptionsStub        func(topicID string, newSubscribers []string) error
	createNewSubscriptionsMutex       sync.RWMutex
	createNewSubscriptionsArgsForCall []struct {
		topicID        string
		newSubscribers []string
	}
	createNewSubscriptionsReturns struct {
		result1 error
	}
	PublishMessageStub        func(topicID string, message string) error
	publishMessageMutex       sync.RWMutex
	publishMessageArgsForCall []struct {
		topicID string
		message string
	}
	publishMessageReturns struct {
		result1 error
	}
	invocations map[string][][]interface{}
}

func (fake *FakeSMSService) CreateTopic(topic string) (string, error) {
	fake.createTopicMutex.Lock()
	fake.createTopicArgsForCall = append(fake.createTopicArgsForCall, struct {
		topic string
	}{topic})
	fake.guard("CreateTopic")
	fake.invocations["CreateTopic"] = append(fake.invocations["CreateTopic"], []interface{}{topic})
	fake.createTopicMutex.Unlock()
	if fake.CreateTopicStub != nil {
		return fake.CreateTopicStub(topic)
	} else {
		return fake.createTopicReturns.result1, fake.createTopicReturns.result2
	}
}

func (fake *FakeSMSService) CreateTopicCallCount() int {
	fake.createTopicMutex.RLock()
	defer fake.createTopicMutex.RUnlock()
	return len(fake.createTopicArgsForCall)
}

func (fake *FakeSMSService) CreateTopicArgsForCall(i int) string {
	fake.createTopicMutex.RLock()
	defer fake.createTopicMutex.RUnlock()
	return fake.createTopicArgsForCall[i].topic
}

func (fake *FakeSMSService) CreateTopicReturns(result1 string, result2 error) {
	fake.CreateTopicStub = nil
	fake.createTopicReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSMSService) GetExistingSubscribers(topicID string) ([]string, error) {
	fake.getExistingSubscribersMutex.Lock()
	fake.getExistingSubscribersArgsForCall = append(fake.getExistingSubscribersArgsForCall, struct {
		topicID string
	}{topicID})
	fake.guard("GetExistingSubscribers")
	fake.invocations["GetExistingSubscribers"] = append(fake.invocations["GetExistingSubscribers"], []interface{}{topicID})
	fake.getExistingSubscribersMutex.Unlock()
	if fake.GetExistingSubscribersStub != nil {
		return fake.GetExistingSubscribersStub(topicID)
	} else {
		return fake.getExistingSubscribersReturns.result1, fake.getExistingSubscribersReturns.result2
	}
}

func (fake *FakeSMSService) GetExistingSubscribersCallCount() int {
	fake.getExistingSubscribersMutex.RLock()
	defer fake.getExistingSubscribersMutex.RUnlock()
	return len(fake.getExistingSubscribersArgsForCall)
}

func (fake *FakeSMSService) GetExistingSubscribersArgsForCall(i int) string {
	fake.getExistingSubscribersMutex.RLock()
	defer fake.getExistingSubscribersMutex.RUnlock()
	return fake.getExistingSubscribersArgsForCall[i].topicID
}

func (fake *FakeSMSService) GetExistingSubscribersReturns(result1 []string, result2 error) {
	fake.GetExistingSubscribersStub = nil
	fake.getExistingSubscribersReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeSMSService) CreateNewSubscriptions(topicID string, newSubscribers []string) error {
	var newSubscribersCopy []string
	if newSubscribers != nil {
		newSubscribersCopy = make([]string, len(newSubscribers))
		copy(newSubscribersCopy, newSubscribers)
	}
	fake.createNewSubscriptionsMutex.Lock()
	fake.createNewSubscriptionsArgsForCall = append(fake.createNewSubscriptionsArgsForCall, struct {
		topicID        string
		newSubscribers []string
	}{topicID, newSubscribersCopy})
	fake.guard("CreateNewSubscriptions")
	fake.invocations["CreateNewSubscriptions"] = append(fake.invocations["CreateNewSubscriptions"], []interface{}{topicID, newSubscribersCopy})
	fake.createNewSubscriptionsMutex.Unlock()
	if fake.CreateNewSubscriptionsStub != nil {
		return fake.CreateNewSubscriptionsStub(topicID, newSubscribers)
	} else {
		return fake.createNewSubscriptionsReturns.result1
	}
}

func (fake *FakeSMSService) CreateNewSubscriptionsCallCount() int {
	fake.createNewSubscriptionsMutex.RLock()
	defer fake.createNewSubscriptionsMutex.RUnlock()
	return len(fake.createNewSubscriptionsArgsForCall)
}

func (fake *FakeSMSService) CreateNewSubscriptionsArgsForCall(i int) (string, []string) {
	fake.createNewSubscriptionsMutex.RLock()
	defer fake.createNewSubscriptionsMutex.RUnlock()
	return fake.createNewSubscriptionsArgsForCall[i].topicID, fake.createNewSubscriptionsArgsForCall[i].newSubscribers
}

func (fake *FakeSMSService) CreateNewSubscriptionsReturns(result1 error) {
	fake.CreateNewSubscriptionsStub = nil
	fake.createNewSubscriptionsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSMSService) PublishMessage(topicID string, message string) error {
	fake.publishMessageMutex.Lock()
	fake.publishMessageArgsForCall = append(fake.publishMessageArgsForCall, struct {
		topicID string
		message string
	}{topicID, message})
	fake.guard("PublishMessage")
	fake.invocations["PublishMessage"] = append(fake.invocations["PublishMessage"], []interface{}{topicID, message})
	fake.publishMessageMutex.Unlock()
	if fake.PublishMessageStub != nil {
		return fake.PublishMessageStub(topicID, message)
	} else {
		return fake.publishMessageReturns.result1
	}
}

func (fake *FakeSMSService) PublishMessageCallCount() int {
	fake.publishMessageMutex.RLock()
	defer fake.publishMessageMutex.RUnlock()
	return len(fake.publishMessageArgsForCall)
}

func (fake *FakeSMSService) PublishMessageArgsForCall(i int) (string, string) {
	fake.publishMessageMutex.RLock()
	defer fake.publishMessageMutex.RUnlock()
	return fake.publishMessageArgsForCall[i].topicID, fake.publishMessageArgsForCall[i].message
}

func (fake *FakeSMSService) PublishMessageReturns(result1 error) {
	fake.PublishMessageStub = nil
	fake.publishMessageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSMSService) Invocations() map[string][][]interface{} {
	return fake.invocations
}

func (fake *FakeSMSService) guard(key string) {
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
}

var _ application.SMSService = new(FakeSMSService)
