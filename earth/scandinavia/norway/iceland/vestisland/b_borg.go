package vestisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔格BorgBarony struct {
	feud.BaseBarony
}

var BBorg博尔格 feud.Barony = &博尔格BorgBarony{}

func init() {
    f := BBorg博尔格.(*博尔格BorgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borg",
		TitleName: "博尔格",
		TitleCode: "b_borg",
	}
}
