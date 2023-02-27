package damascus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜马DumaBarony struct {
	feud.BaseBarony
}

var BDuma杜马 feud.Barony = &杜马DumaBarony{}

func init() {
    f := BDuma杜马.(*杜马DumaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "duma",
		TitleName: "杜马",
		TitleCode: "b_duma",
	}
}
