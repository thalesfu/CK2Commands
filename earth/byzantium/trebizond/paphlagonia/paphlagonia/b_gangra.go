package paphlagonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 冈格拉GangraBarony struct {
	feud.BaseBarony
}

var BGangra冈格拉 feud.Barony = &冈格拉GangraBarony{}

func init() {
	f := BGangra冈格拉.(*冈格拉GangraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gangra",
		TitleName: "冈格拉",
		TitleCode: "b_gangra",
	}
}
