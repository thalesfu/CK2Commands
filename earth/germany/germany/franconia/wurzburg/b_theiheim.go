package wurzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰尔海姆TheiheimBarony struct {
	feud.BaseBarony
}

var BTheiheim泰尔海姆 feud.Barony = &泰尔海姆TheiheimBarony{}

func init() {
	f := BTheiheim泰尔海姆.(*泰尔海姆TheiheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "theiheim",
		TitleName: "泰尔海姆",
		TitleCode: "b_theiheim",
	}
}
