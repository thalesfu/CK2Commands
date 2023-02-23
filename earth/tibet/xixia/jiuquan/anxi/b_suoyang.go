package anxi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锁阳SuoyangBarony struct {
	feud.BaseBarony
}

var BSuoyang锁阳 feud.Barony = &锁阳SuoyangBarony{}

func init() {
	f := BSuoyang锁阳.(*锁阳SuoyangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suoyang",
		TitleName: "锁阳",
		TitleCode: "b_suoyang",
	}
}
