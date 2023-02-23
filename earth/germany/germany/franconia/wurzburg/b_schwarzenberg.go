package wurzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施瓦岑贝格SchwarzenbergBarony struct {
	feud.BaseBarony
}

var BSchwarzenberg施瓦岑贝格 feud.Barony = &施瓦岑贝格SchwarzenbergBarony{}

func init() {
	f := BSchwarzenberg施瓦岑贝格.(*施瓦岑贝格SchwarzenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schwarzenberg",
		TitleName: "施瓦岑贝格",
		TitleCode: "b_schwarzenberg",
	}
}
