package quattara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代尔DayrBarony struct {
	feud.BaseBarony
}

var BDayr代尔 feud.Barony = &代尔DayrBarony{}

func init() {
    f := BDayr代尔.(*代尔DayrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dayr",
		TitleName: "代尔",
		TitleCode: "b_dayr",
	}
}
