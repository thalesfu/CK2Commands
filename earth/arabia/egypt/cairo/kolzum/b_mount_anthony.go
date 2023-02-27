package kolzum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安东尼山Mount_anthonyBarony struct {
	feud.BaseBarony
}

var BMount_anthony安东尼山 feud.Barony = &安东尼山Mount_anthonyBarony{}

func init() {
    f := BMount_anthony安东尼山.(*安东尼山Mount_anthonyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mount_anthony",
		TitleName: "安东尼山",
		TitleCode: "b_mount_anthony",
	}
}
