package neuchatel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔盖尔ErguelBarony struct {
	feud.BaseBarony
}

var BErguel埃尔盖尔 feud.Barony = &埃尔盖尔ErguelBarony{}

func init() {
    f := BErguel埃尔盖尔.(*埃尔盖尔ErguelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erguel",
		TitleName: "埃尔盖尔",
		TitleCode: "b_erguel",
	}
}
