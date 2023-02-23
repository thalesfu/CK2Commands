package naimisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 郡荼GondaBarony struct {
	feud.BaseBarony
}

var BGonda郡荼 feud.Barony = &郡荼GondaBarony{}

func init() {
	f := BGonda郡荼.(*郡荼GondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gonda",
		TitleName: "郡荼",
		TitleCode: "b_gonda",
	}
}
