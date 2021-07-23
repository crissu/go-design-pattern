package bridge

import "log"

// IMsgSender IMsgSender
type IMsgSender interface {
	Send(msg string) error
}

/*
********** 可以变化的 维度 ***********
 */

// 维度 1
// EmailMsgSender 发送邮件
// 可能还有 电话、短信等各种实现
type EmailMsgSender struct {
	emails []string
}

// NewEmailMsgSender NewEmailMsgSender
func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

// Send Send
func (s *EmailMsgSender) Send(msg string) error {
	// 这里去发送消息
	log.Println("EmailMsgSender emails:", s.emails)
	log.Println("EmailMsgSender send:", msg)
	return nil
}

// 维度 2
type PhoneMsgSender struct {
	talkMsgs []string
}

func NewPhoneMsgSender(talkMsgs []string) *PhoneMsgSender {
	return &PhoneMsgSender{talkMsgs: talkMsgs}
}

func (p *PhoneMsgSender) Send(msg string) error {
	log.Println("PhoneMsgSender talkMsgs:", p.talkMsgs)
	log.Println("PhoneMsgSender send:", msg)
	return nil
}


// INotification 通知接口
type INotification interface {
	Notify(msg string) error
}

// ErrorNotification 错误通知
// 后面可能还有 warning 各种级别
type ErrorNotification struct {
	sender IMsgSender
}

// NewErrorNotification NewErrorNotification
func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

// Notify 发送通知
func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}
