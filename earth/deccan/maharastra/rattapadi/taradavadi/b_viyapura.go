package taradavadi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗耶补罗ViyapuraBarony struct {
	feud.BaseBarony
}

var BViyapura毗耶补罗 feud.Barony = &毗耶补罗ViyapuraBarony{}

func init() {
	f := BViyapura毗耶补罗.(*毗耶补罗ViyapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "viyapura",
		TitleName: "毗耶补罗",
		TitleCode: "b_viyapura",
	}
}
