package zamora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝纳文特BenaventeBarony struct {
	feud.BaseBarony
}

var BBenavente贝纳文特 feud.Barony = &贝纳文特BenaventeBarony{}

func init() {
	f := BBenavente贝纳文特.(*贝纳文特BenaventeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "benavente",
		TitleName: "贝纳文特",
		TitleCode: "b_benavente",
	}
}
