package gurma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔玛GurmaBarony struct {
	feud.BaseBarony
}

var BGurma古尔玛 feud.Barony = &古尔玛GurmaBarony{}

func init() {
	f := BGurma古尔玛.(*古尔玛GurmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gurma",
		TitleName: "古尔玛",
		TitleCode: "b_gurma",
	}
}
