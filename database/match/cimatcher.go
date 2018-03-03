package match

import (
	"github.com/golang/mock/gomock"
	"github.com/q231950/alethea/ci"
	"github.com/q231950/alethea/model"
)

type ciType struct{ kind ci.CI }

func CIType(kind ci.CI) gomock.Matcher {
	return &ciType{kind}
}

func (o *ciType) Matches(x interface{}) bool {
	incident := x.(model.Incident)

	return incident.String() == o.kind.String()
}

func (o *ciType) String() string {
	return "Kind of CI is expected to be:`" + o.kind.String() + "`"
}
