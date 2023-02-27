package merv

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库什卡KushkaBarony struct {
	feud.BaseBarony
}

var BKushka库什卡 feud.Barony = &库什卡KushkaBarony{}

func init() {
    f := BKushka库什卡.(*库什卡KushkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kushka",
		TitleName: "库什卡",
		TitleCode: "b_kushka",
	}
}
