package tudgha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图德盖TudghaBarony struct {
	feud.BaseBarony
}

var BTudgha图德盖 feud.Barony = &图德盖TudghaBarony{}

func init() {
	f := BTudgha图德盖.(*图德盖TudghaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tudgha",
		TitleName: "图德盖",
		TitleCode: "b_tudgha",
	}
}
