package lattalura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔加拉TagaraBarony struct {
	feud.BaseBarony
}

var BTagara塔加拉 feud.Barony = &塔加拉TagaraBarony{}

func init() {
	f := BTagara塔加拉.(*塔加拉TagaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tagara",
		TitleName: "塔加拉",
		TitleCode: "b_tagara",
	}
}
