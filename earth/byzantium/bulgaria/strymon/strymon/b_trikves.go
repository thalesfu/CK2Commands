package strymon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂克韦什TrikvesBarony struct {
	feud.BaseBarony
}

var BTrikves蒂克韦什 feud.Barony = &蒂克韦什TrikvesBarony{}

func init() {
	f := BTrikves蒂克韦什.(*蒂克韦什TrikvesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trikves",
		TitleName: "蒂克韦什",
		TitleCode: "b_trikves",
	}
}
