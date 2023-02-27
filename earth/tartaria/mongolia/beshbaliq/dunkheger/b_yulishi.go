package dunkheger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 郁立师YulishiBarony struct {
	feud.BaseBarony
}

var BYulishi郁立师 feud.Barony = &郁立师YulishiBarony{}

func init() {
    f := BYulishi郁立师.(*郁立师YulishiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yulishi",
		TitleName: "郁立师",
		TitleCode: "b_yulishi",
	}
}
