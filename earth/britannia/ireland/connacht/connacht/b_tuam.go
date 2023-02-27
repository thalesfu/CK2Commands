package connacht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂厄姆TuamBarony struct {
	feud.BaseBarony
}

var BTuam蒂厄姆 feud.Barony = &蒂厄姆TuamBarony{}

func init() {
    f := BTuam蒂厄姆.(*蒂厄姆TuamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tuam",
		TitleName: "蒂厄姆",
		TitleCode: "b_tuam",
	}
}
