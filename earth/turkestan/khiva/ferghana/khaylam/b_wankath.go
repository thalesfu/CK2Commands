package khaylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 万卡思WankathBarony struct {
	feud.BaseBarony
}

var BWankath万卡思 feud.Barony = &万卡思WankathBarony{}

func init() {
	f := BWankath万卡思.(*万卡思WankathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wankath",
		TitleName: "万卡思",
		TitleCode: "b_wankath",
	}
}
