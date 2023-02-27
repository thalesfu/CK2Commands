package rostock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 居斯特罗GustrowBarony struct {
	feud.BaseBarony
}

var BGustrow居斯特罗 feud.Barony = &居斯特罗GustrowBarony{}

func init() {
    f := BGustrow居斯特罗.(*居斯特罗GustrowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gustrow",
		TitleName: "居斯特罗",
		TitleCode: "b_gustrow",
	}
}
