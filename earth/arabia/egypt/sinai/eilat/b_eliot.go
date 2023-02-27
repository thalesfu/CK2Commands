package eilat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 以禄EliotBarony struct {
	feud.BaseBarony
}

var BEliot以禄 feud.Barony = &以禄EliotBarony{}

func init() {
    f := BEliot以禄.(*以禄EliotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eliot",
		TitleName: "以禄",
		TitleCode: "b_eliot",
	}
}
