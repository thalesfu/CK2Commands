package gotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海姆瑟HemseBarony struct {
	feud.BaseBarony
}

var BHemse海姆瑟 feud.Barony = &海姆瑟HemseBarony{}

func init() {
	f := BHemse海姆瑟.(*海姆瑟HemseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hemse",
		TitleName: "海姆瑟",
		TitleCode: "b_hemse",
	}
}
