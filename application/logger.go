package application

import (
	"encoding/json"
	"fmt"
	loglib "log"
)

type Log interface {
	WithFields(map[string]interface{}) Log
	Warn(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}

type log struct {
	fields map[string]interface{}
}

func (l *log) WithFields(fields map[string]interface{}) Log {

	mergedFields := make(map[string]interface{})

	for k, v := range l.fields {
		mergedFields[k] = v
	}

	for k, v := range fields {
		mergedFields[k] = v
	}

	return &log{
		fields: mergedFields,
	}
}

func (l *log) printf(lvl string, args ...interface{}) {

	stringArgs := ""
	for _, arg := range args {
		switch t := arg.(type) {
		case string:
			stringArgs += t
		case []Match:
			stringArgs += fmt.Sprintf("%#v", goStringableCombination(t))
		case [][]Match:
			stringArgs += fmt.Sprintf("%#v", goStringableCombinationSet(t))
		case []Movement:
			stringArgs += fmt.Sprintf("%#v", goStringableMovementSet(t))
		default:
			stringArgs += fmt.Sprintf("%#v", t)

		}
	}
	loglib.Printf(" [%s] fields: %+v \r\n message: %s\r\n ", lvl, l.fields, stringArgs)
}

func (l *log) Warn(args ...interface{}) {
	lvl := "WARN"
	l.printf(lvl, args...)
}
func (l *log) Info(args ...interface{}) {
	lvl := "INFO"
	l.printf(lvl, args...)
}
func (l *log) Debug(args ...interface{}) {
	lvl := "DEBU"
	l.printf(lvl, args...)
}

func (mvmt Movement) GoString() string {
	b, _ := json.Marshal(mvmt)
	return string(b)
}

func (match Match) GoString() string {

	result := "\r\nMovements:\r\n"

	for mvNo, mvmt := range match.Movements {
		result += fmt.Sprintf("%d : %#v\r\n", mvNo, mvmt)
	}
	result += fmt.Sprintf("ContractCondition:\r\n%#v\r\nScore: %d\r\n", match.ContractCondition, match.Score)
	return result
}

type goStringableCombination []Match

func (combi goStringableCombination) GoString() string {

	result := "\r\nMatch Combination:\r\n"

	for mNo, m := range combi {
		result += fmt.Sprintf("%d : %#v\r\n", mNo, m)
	}

	return result
}

type goStringableCombinationSet [][]Match

func (combiSet goStringableCombinationSet) GoString() string {

	result := "\r\nCombination Set:\r\n"

	for mNo, m := range combiSet {
		result += fmt.Sprintf("%d : %#v\r\n", mNo, m)
	}

	return result
}

type goStringableMovementSet []Movement

func (mvmtSet goStringableMovementSet) GoString() string {

	result := "\r\nMovement Set:\r\n"

	for mNo, m := range mvmtSet {
		result += fmt.Sprintf("%d : %#v\r\n", mNo, m)
	}

	return result
}
