package odessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤日涅YuzhneBarony struct {
	feud.BaseBarony
}

var BYuzhne尤日涅 feud.Barony = &尤日涅YuzhneBarony{}

func init() {
    f := BYuzhne尤日涅.(*尤日涅YuzhneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yuzhne",
		TitleName: "尤日涅",
		TitleCode: "b_yuzhne",
	}
}
