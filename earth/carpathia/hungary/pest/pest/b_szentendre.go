package pest

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣安德烈SzentendreBarony struct {
	feud.BaseBarony
}

var BSzentendre圣安德烈 feud.Barony = &圣安德烈SzentendreBarony{}

func init() {
    f := BSzentendre圣安德烈.(*圣安德烈SzentendreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szentendre",
		TitleName: "圣安德烈",
		TitleCode: "b_szentendre",
	}
}
