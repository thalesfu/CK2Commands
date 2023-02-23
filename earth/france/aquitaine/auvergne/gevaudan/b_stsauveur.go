package gevaudan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣索弗尔StsauveurBarony struct {
	feud.BaseBarony
}

var BStsauveur圣索弗尔 feud.Barony = &圣索弗尔StsauveurBarony{}

func init() {
	f := BStsauveur圣索弗尔.(*圣索弗尔StsauveurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stsauveur",
		TitleName: "圣索弗尔",
		TitleCode: "b_stsauveur",
	}
}
