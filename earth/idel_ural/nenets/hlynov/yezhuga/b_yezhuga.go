package yezhuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶茹加YezhugaBarony struct {
	feud.BaseBarony
}

var BYezhuga叶茹加 feud.Barony = &叶茹加YezhugaBarony{}

func init() {
    f := BYezhuga叶茹加.(*叶茹加YezhugaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yezhuga",
		TitleName: "叶茹加",
		TitleCode: "b_yezhuga",
	}
}
