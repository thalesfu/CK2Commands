package sedyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克德瓦KedvaBarony struct {
	feud.BaseBarony
}

var BKedva克德瓦 feud.Barony = &克德瓦KedvaBarony{}

func init() {
    f := BKedva克德瓦.(*克德瓦KedvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kedva",
		TitleName: "克德瓦",
		TitleCode: "b_kedva",
	}
}
