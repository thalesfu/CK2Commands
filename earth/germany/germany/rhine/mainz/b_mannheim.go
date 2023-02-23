package mainz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼海姆MannheimBarony struct {
	feud.BaseBarony
}

var BMannheim曼海姆 feud.Barony = &曼海姆MannheimBarony{}

func init() {
	f := BMannheim曼海姆.(*曼海姆MannheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mannheim",
		TitleName: "曼海姆",
		TitleCode: "b_mannheim",
	}
}
