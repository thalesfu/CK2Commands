package duqm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜库姆DuqmBarony struct {
	feud.BaseBarony
}

var BDuqm杜库姆 feud.Barony = &杜库姆DuqmBarony{}

func init() {
    f := BDuqm杜库姆.(*杜库姆DuqmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "duqm",
		TitleName: "杜库姆",
		TitleCode: "b_duqm",
	}
}
