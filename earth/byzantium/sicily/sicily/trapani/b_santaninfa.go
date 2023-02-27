package trapani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣宁法SantaninfaBarony struct {
	feud.BaseBarony
}

var BSantaninfa圣宁法 feud.Barony = &圣宁法SantaninfaBarony{}

func init() {
    f := BSantaninfa圣宁法.(*圣宁法SantaninfaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santaninfa",
		TitleName: "圣宁法",
		TitleCode: "b_santaninfa",
	}
}
