package vermandois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉斯GuiseBarony struct {
	feud.BaseBarony
}

var BGuise吉斯 feud.Barony = &吉斯GuiseBarony{}

func init() {
	f := BGuise吉斯.(*吉斯GuiseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guise",
		TitleName: "吉斯",
		TitleCode: "b_guise",
	}
}
