package shahrazur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑达SandaBarony struct {
	feud.BaseBarony
}

var BSanda桑达 feud.Barony = &桑达SandaBarony{}

func init() {
    f := BSanda桑达.(*桑达SandaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanda",
		TitleName: "桑达",
		TitleCode: "b_sanda",
	}
}
