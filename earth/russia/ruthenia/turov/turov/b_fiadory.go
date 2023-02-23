package turov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费多雷FiadoryBarony struct {
	feud.BaseBarony
}

var BFiadory费多雷 feud.Barony = &费多雷FiadoryBarony{}

func init() {
	f := BFiadory费多雷.(*费多雷FiadoryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fiadory",
		TitleName: "费多雷",
		TitleCode: "b_fiadory",
	}
}
