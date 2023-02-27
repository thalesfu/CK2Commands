package dyfed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 滕比TenbyBarony struct {
	feud.BaseBarony
}

var BTenby滕比 feud.Barony = &滕比TenbyBarony{}

func init() {
    f := BTenby滕比.(*滕比TenbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tenby",
		TitleName: "滕比",
		TitleCode: "b_tenby",
	}
}
