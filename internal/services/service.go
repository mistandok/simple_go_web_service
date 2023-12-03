package services

import "errors"

type SimpleLogic struct {
	logger    Logger
	dataStore DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.logger.Log("in SayHello for " + userID)
	name, ok := sl.dataStore.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.logger.Log("in SayGoodbye for " + userID)
	name, ok := sl.dataStore.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

func NewSimpleLogic(logger Logger, dataStore DataStore) SimpleLogic {
	return SimpleLogic{
		logger:    logger,
		dataStore: dataStore,
	}
}
