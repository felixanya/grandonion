package hellomock

import (
	"github.com/aaronize/grandonion/examples/mock-demo/hellomock/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCompany_Meeting(t *testing.T) {
	person := NewPerson("Jeseling")
	company := NewCompany(person)
	t.Log(company.Meeting("Aaron"))
}

func TestCompany_Meeting2(t *testing.T) {
	ctl := gomock.NewController(t)
	mock_talker := mock_hellomock.NewMockTalker(ctl)
	mock_talker.EXPECT().SayHello(gomock.Eq("Stephen")).Return("Custom return info.")

	company := NewCompany(mock_talker)
	t.Log(company.Meeting("Stephen"))
}
