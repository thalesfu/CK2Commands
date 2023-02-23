package berbera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢赫ShiikhBarony struct {
	feud.BaseBarony
}

var BShiikh谢赫 feud.Barony = &谢赫ShiikhBarony{}

func init() {
	f := BShiikh谢赫.(*谢赫ShiikhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shiikh",
		TitleName: "谢赫",
		TitleCode: "b_shiikh",
	}
}
