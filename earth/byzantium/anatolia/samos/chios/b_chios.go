package chios

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希俄斯ChiosBarony struct {
	feud.BaseBarony
}

var BChios希俄斯 feud.Barony = &希俄斯ChiosBarony{}

func init() {
    f := BChios希俄斯.(*希俄斯ChiosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chios",
		TitleName: "希俄斯",
		TitleCode: "b_chios",
	}
}
