package bari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 焦维纳佐GiovinazzoBarony struct {
	feud.BaseBarony
}

var BGiovinazzo焦维纳佐 feud.Barony = &焦维纳佐GiovinazzoBarony{}

func init() {
    f := BGiovinazzo焦维纳佐.(*焦维纳佐GiovinazzoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "giovinazzo",
		TitleName: "焦维纳佐",
		TitleCode: "b_giovinazzo",
	}
}
