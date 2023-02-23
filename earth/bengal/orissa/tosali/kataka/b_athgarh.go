package kataka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 八舍AthgarhBarony struct {
	feud.BaseBarony
}

var BAthgarh八舍 feud.Barony = &八舍AthgarhBarony{}

func init() {
	f := BAthgarh八舍.(*八舍AthgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "athgarh",
		TitleName: "八舍",
		TitleCode: "b_athgarh",
	}
}
