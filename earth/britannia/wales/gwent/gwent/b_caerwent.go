package gwent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯尔文特CaerwentBarony struct {
	feud.BaseBarony
}

var BCaerwent凯尔文特 feud.Barony = &凯尔文特CaerwentBarony{}

func init() {
	f := BCaerwent凯尔文特.(*凯尔文特CaerwentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caerwent",
		TitleName: "凯尔文特",
		TitleCode: "b_caerwent",
	}
}
