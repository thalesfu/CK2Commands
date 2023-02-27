package belaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌伦达UrundaBarony struct {
	feud.BaseBarony
}

var BUrunda乌伦达 feud.Barony = &乌伦达UrundaBarony{}

func init() {
    f := BUrunda乌伦达.(*乌伦达UrundaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urunda",
		TitleName: "乌伦达",
		TitleCode: "b_urunda",
	}
}
