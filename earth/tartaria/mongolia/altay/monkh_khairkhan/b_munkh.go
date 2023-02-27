package monkh_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙赫MunkhBarony struct {
	feud.BaseBarony
}

var BMunkh蒙赫 feud.Barony = &蒙赫MunkhBarony{}

func init() {
    f := BMunkh蒙赫.(*蒙赫MunkhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "munkh",
		TitleName: "蒙赫",
		TitleCode: "b_munkh",
	}
}
