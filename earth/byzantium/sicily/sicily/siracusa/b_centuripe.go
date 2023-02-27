package siracusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 琴图里佩CenturipeBarony struct {
	feud.BaseBarony
}

var BCenturipe琴图里佩 feud.Barony = &琴图里佩CenturipeBarony{}

func init() {
    f := BCenturipe琴图里佩.(*琴图里佩CenturipeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "centuripe",
		TitleName: "琴图里佩",
		TitleCode: "b_centuripe",
	}
}
