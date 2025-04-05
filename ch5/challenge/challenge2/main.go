package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (dm directMessage) importance() int {
	if dm.isUrgent {
		return 50
	}
	return dm.priorityLevel
}

func (gm groupMessage) importance() int {
	score := gm.priorityLevel
	return score
}

func (sa systemAlert) importance() int {
	score := 100
	return score
}

func processNotification(n notification) (string, int) {
	switch n.(type) {
	case directMessage:
		dm := n.(directMessage)
		return dm.senderUsername, dm.importance()
	case groupMessage:
		gm := n.(groupMessage)
		return gm.groupName, gm.importance()
	case systemAlert:
		sa := n.(systemAlert)
		return sa.alertCode, sa.importance()
	default:
		return "", 0
	}
}
