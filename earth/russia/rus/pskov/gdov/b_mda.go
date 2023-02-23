package gdov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆达MdaBarony struct {
	feud.BaseBarony
}

var BMda姆达 feud.Barony = &姆达MdaBarony{}

func init() {
	f := BMda姆达.(*姆达MdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mda",
		TitleName: "姆达",
		TitleCode: "b_mda",
	}
}
