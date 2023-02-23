package barcelona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌孙纳OsonaBarony struct {
	feud.BaseBarony
}

var BOsona乌孙纳 feud.Barony = &乌孙纳OsonaBarony{}

func init() {
	f := BOsona乌孙纳.(*乌孙纳OsonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osona",
		TitleName: "乌孙纳",
		TitleCode: "b_osona",
	}
}
