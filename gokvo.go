package gokvo

import "errors"

// Observer unified registration notification data type
type Observer interface {
}

var (
	// ErrInvalidEvent will be returned when invokers don't Registration issue
	ErrInvalidEvent = errors.New("invalid event")
	// ErrInvalidPool will be returned when invokers don't Initialize the observer pool
	ErrInvalidPool = errors.New("invalid oberPool")
	// ErrInvalidSubscriber will be returned when have not sub
	ErrInvalidSubscriber = errors.New("invalid Subscriber")
)

// DefaultNotificationCenter is the default center used by Serve.
// Singleton pattern
var DefaultNotificationCenter = &defaultNotificationCenter

var defaultNotificationCenter NotificationCenter

// AddObserver Binding keys and notifications
func AddObserver(ober Observer, name EventName, action NotifiAction) {
	DefaultNotificationCenter.addObserver(ober, name, action)
}

// Post Send a notification to deliver a message
func Post(name EventName, info interface{}) error {
	return DefaultNotificationCenter.post(name, info)
}

// RemoveObserver Remove object from observation pool
func RemoveObserver(ober Observer, name EventName) {
	DefaultNotificationCenter.removeObserver(ober, name)
}
