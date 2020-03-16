package hellomock

import (
	"testing"

	mock_hellomock "mygotraining/cmd/review/gotraining/package/mock/hellomock/mock"

	"github.com/golang/mock/gomock"
)

func TestCompany_Meeting(t *testing.T) {
	person := NewPerson("王尼美")
	company := NewCompany(person)
	t.Log(company.Meeting("王尼玛"))
}

func TestCompany_Meeting2(t *testing.T) {
	ctl := gomock.NewController(t)

	mock_talker := mock_hellomock.NewMockTalker(ctl)
	mock_talker.EXPECT().SayHello(gomock.Eq("王尼玛")).Return("asdfsdfsdf")

	company := NewCompany(mock_talker)
	t.Log(company.Meeting("王尼玛"))
	//t.Log(company.Meeting("张全蛋"))
}
