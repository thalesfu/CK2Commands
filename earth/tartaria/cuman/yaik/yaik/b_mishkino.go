package yaik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米什基诺MishkinoBarony struct {
	feud.BaseBarony
}

var BMishkino米什基诺 feud.Barony = &米什基诺MishkinoBarony{}

func init() {
	f := BMishkino米什基诺.(*米什基诺MishkinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mishkino",
		TitleName: "米什基诺",
		TitleCode: "b_mishkino",
	}
}
