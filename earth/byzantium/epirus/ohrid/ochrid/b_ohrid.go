package ochrid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥赫里德OhridBarony struct {
	feud.BaseBarony
}

var BOhrid奥赫里德 feud.Barony = &奥赫里德OhridBarony{}

func init() {
    f := BOhrid奥赫里德.(*奥赫里德OhridBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ohrid",
		TitleName: "奥赫里德",
		TitleCode: "b_ohrid",
	}
}
