// Code generated by Validator v0.1.4. DO NOT EDIT.

package douyinmessage

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *BaseResp) IsValid() error {
	return nil
}
func (p *Message) IsValid() error {
	return nil
}
func (p *ChatRecordRequest) IsValid() error {
	if p.ToUserId <= int64(0) {
		return fmt.Errorf("field ToUserId gt rule failed, current value: %v", p.ToUserId)
	}
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	return nil
}
func (p *ChatRecordResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *SendMessageRequest) IsValid() error {
	if p.ToUserId <= int64(0) {
		return fmt.Errorf("field ToUserId gt rule failed, current value: %v", p.ToUserId)
	}
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	_src := []int64{int64(1)}
	var _exist bool
	for _, src := range _src {
		if p.ActionType == int64(src) {
			_exist = true
			break
		}
	}
	if !_exist {
		return fmt.Errorf("field ActionType in rule failed, current value: %v", p.ActionType)
	}
	return nil
}
func (p *SendMessageResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *GetFirstMessagesRequest) IsValid() error {
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	return nil
}
func (p *GetFirstMessagesResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}