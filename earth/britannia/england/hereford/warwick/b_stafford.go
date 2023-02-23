package warwick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯塔福德StaffordBarony struct {
	feud.BaseBarony
}

var BStafford斯塔福德 feud.Barony = &斯塔福德StaffordBarony{}

func init() {
	f := BStafford斯塔福德.(*斯塔福德StaffordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stafford",
		TitleName: "斯塔福德",
		TitleCode: "b_stafford",
	}
}
