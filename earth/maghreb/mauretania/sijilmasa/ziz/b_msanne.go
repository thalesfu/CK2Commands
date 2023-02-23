package ziz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆桑MsanneBarony struct {
	feud.BaseBarony
}

var BMsanne姆桑 feud.Barony = &姆桑MsanneBarony{}

func init() {
	f := BMsanne姆桑.(*姆桑MsanneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "msanne",
		TitleName: "姆桑",
		TitleCode: "b_msanne",
	}
}
