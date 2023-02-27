package livs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔特纳MartnaBarony struct {
	feud.BaseBarony
}

var BMartna马尔特纳 feud.Barony = &马尔特纳MartnaBarony{}

func init() {
    f := BMartna马尔特纳.(*马尔特纳MartnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "martna",
		TitleName: "马尔特纳",
		TitleCode: "b_martna",
	}
}
