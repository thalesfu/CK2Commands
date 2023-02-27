package samtho

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朝阳TraoyangBarony struct {
	feud.BaseBarony
}

var BTraoyang朝阳 feud.Barony = &朝阳TraoyangBarony{}

func init() {
    f := BTraoyang朝阳.(*朝阳TraoyangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "traoyang",
		TitleName: "朝阳",
		TitleCode: "b_traoyang",
	}
}
