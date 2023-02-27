package gerze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 改则GerzeBarony struct {
	feud.BaseBarony
}

var BGerze改则 feud.Barony = &改则GerzeBarony{}

func init() {
    f := BGerze改则.(*改则GerzeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gerze",
		TitleName: "改则",
		TitleCode: "b_gerze",
	}
}
