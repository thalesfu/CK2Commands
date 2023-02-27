package djerba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米都恩MidounBarony struct {
	feud.BaseBarony
}

var BMidoun米都恩 feud.Barony = &米都恩MidounBarony{}

func init() {
    f := BMidoun米都恩.(*米都恩MidounBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "midoun",
		TitleName: "米都恩",
		TitleCode: "b_midoun",
	}
}
