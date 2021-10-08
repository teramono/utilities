package broker

import "fmt"

func Namespace(version uint, action string, object string) string {
	return fmt.Sprintf("v%d.%s.%s", version, action, object)
}

func SubjectAny(version uint, action string, object string) string {
	return Namespace(version, action, object) + ".*"
}

func SubjectOf(version uint, action string, object string, id interface{}) string {
	return fmt.Sprintf("%s.%v", Namespace(version, action, object), id)
}
