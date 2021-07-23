package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNotification_Notify(t *testing.T) {
	sender := NewEmailMsgSender([]string{"test@test.com", "cris@163.com"})
	n := NewErrorNotification(sender)
	err := n.Notify("test msg 1")
	assert.Nil(t, err)

	sender2 := NewPhoneMsgSender([]string{"你饭吃了吗", "我吃了"})
	n2 := NewErrorNotification(sender2)
	err = n2.Notify("test msg 2")
	assert.Nil(t, err)
}
