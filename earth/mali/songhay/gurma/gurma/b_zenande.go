package gurma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽南德ZenandeBarony struct {
	feud.BaseBarony
}

var BZenande泽南德 feud.Barony = &泽南德ZenandeBarony{}

func init() {
    f := BZenande泽南德.(*泽南德ZenandeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zenande",
		TitleName: "泽南德",
		TitleCode: "b_zenande",
	}
}
