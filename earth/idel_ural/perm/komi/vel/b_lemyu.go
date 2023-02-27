package vel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列米尤LemyuBarony struct {
	feud.BaseBarony
}

var BLemyu列米尤 feud.Barony = &列米尤LemyuBarony{}

func init() {
    f := BLemyu列米尤.(*列米尤LemyuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lemyu",
		TitleName: "列米尤",
		TitleCode: "b_lemyu",
	}
}
