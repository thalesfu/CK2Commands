package gevaudan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒皮LepuyBarony struct {
	feud.BaseBarony
}

var BLepuy勒皮 feud.Barony = &勒皮LepuyBarony{}

func init() {
	f := BLepuy勒皮.(*勒皮LepuyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lepuy",
		TitleName: "勒皮",
		TitleCode: "b_lepuy",
	}
}
