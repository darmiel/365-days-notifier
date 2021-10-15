package main

import "encoding/json"

func WrapRawToTypedEvent(val interface{}) (resp *TypedEvent, data []byte, err error) {
	// what. the. fuck.
	// there HAS to be a better way
	if data, err = json.Marshal(val); err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

func ParseActivities(raw []interface{}) (resp []CodeActivity, err error) {
	for _, r := range raw {
		var (
			typed *TypedEvent
			data  []byte
		)
		if typed, data, err = WrapRawToTypedEvent(r); err != nil {
			return
		}
		switch typed.Type {
		case TypeCreate:
			var a *EventCreate
			if err = json.Unmarshal(data, &a); err != nil {
				return
			}
			resp = append(resp, a)
		case TypePush:
			var a *EventPush
			if err = json.Unmarshal(data, &a); err != nil {
				return
			}
			resp = append(resp, a)
		case TypeIssue:
			var a *EventIssue
			if err = json.Unmarshal(data, &a); err != nil {
				return
			}
			resp = append(resp, a)
		case TypePullRequest:
			var a *EventPullRequest
			if err = json.Unmarshal(data, &a); err != nil {
				return
			}
			resp = append(resp, a)
		case TypeWatch:
			var a *EventWatch
			if err = json.Unmarshal(data, &a); err != nil {
				return
			}
			resp = append(resp, a)
		}
	}
	return
}
