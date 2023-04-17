package control

import (
	"gihub.com/darkmagic66/a-tour-of-go/model"
)

type MemberData struct {
	data []model.Member
}

func NewMember() *MemberData {
	return &MemberData{}
}

func (m *MemberData) GetAllData() []model.Member {
	return m.data
}

func (m *MemberData) PostAddData(data model.Member) model.Member {
	m.data = append(m.data, data)
	return data
}
