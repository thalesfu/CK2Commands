package yamalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌连戈伊UrengoiBarony struct {
	feud.BaseBarony
}

var BUrengoi乌连戈伊 feud.Barony = &乌连戈伊UrengoiBarony{}

func init() {
    f := BUrengoi乌连戈伊.(*乌连戈伊UrengoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urengoi",
		TitleName: "乌连戈伊",
		TitleCode: "b_urengoi",
	}
}
