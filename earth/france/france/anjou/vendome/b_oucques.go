package vendome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌克OucquesBarony struct {
	feud.BaseBarony
}

var BOucques乌克 feud.Barony = &乌克OucquesBarony{}

func init() {
	f := BOucques乌克.(*乌克OucquesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oucques",
		TitleName: "乌克",
		TitleCode: "b_oucques",
	}
}
