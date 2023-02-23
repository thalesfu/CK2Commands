package varadzin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦拉日丁VarazdinBarony struct {
	feud.BaseBarony
}

var BVarazdin瓦拉日丁 feud.Barony = &瓦拉日丁VarazdinBarony{}

func init() {
	f := BVarazdin瓦拉日丁.(*瓦拉日丁VarazdinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varazdin",
		TitleName: "瓦拉日丁",
		TitleCode: "b_varazdin",
	}
}
