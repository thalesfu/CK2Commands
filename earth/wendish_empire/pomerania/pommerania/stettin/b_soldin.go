package stettin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索尔丁SoldinBarony struct {
	feud.BaseBarony
}

var BSoldin索尔丁 feud.Barony = &索尔丁SoldinBarony{}

func init() {
    f := BSoldin索尔丁.(*索尔丁SoldinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soldin",
		TitleName: "索尔丁",
		TitleCode: "b_soldin",
	}
}
