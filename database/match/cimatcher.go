package match

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/q231950/alethea/ci"
)

type ciType struct{ kind ci.CI }

func CIType(kind ci.CI) gomock.Matcher {
	return &ciType{kind}
}

func (o *ciType) Matches(x interface{}) bool {
	return reflect.TypeOf(x).String() == o.String()
}

func (o *ciType) String() string {
	return "is expected to be " + o.kind.String()
}
