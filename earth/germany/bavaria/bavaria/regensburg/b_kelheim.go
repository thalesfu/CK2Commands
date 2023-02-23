package regensburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯尔海姆KelheimBarony struct {
	feud.BaseBarony
}

var BKelheim凯尔海姆 feud.Barony = &凯尔海姆KelheimBarony{}

func init() {
	f := BKelheim凯尔海姆.(*凯尔海姆KelheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kelheim",
		TitleName: "凯尔海姆",
		TitleCode: "b_kelheim",
	}
}
