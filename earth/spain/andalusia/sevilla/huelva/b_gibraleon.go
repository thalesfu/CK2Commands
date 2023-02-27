package huelva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希夫拉莱昂GibraleonBarony struct {
	feud.BaseBarony
}

var BGibraleon希夫拉莱昂 feud.Barony = &希夫拉莱昂GibraleonBarony{}

func init() {
    f := BGibraleon希夫拉莱昂.(*希夫拉莱昂GibraleonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gibraleon",
		TitleName: "希夫拉莱昂",
		TitleCode: "b_gibraleon",
	}
}
