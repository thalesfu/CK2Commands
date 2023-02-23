package messina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托雷格罗塔TorregrotaBarony struct {
	feud.BaseBarony
}

var BTorregrota托雷格罗塔 feud.Barony = &托雷格罗塔TorregrotaBarony{}

func init() {
	f := BTorregrota托雷格罗塔.(*托雷格罗塔TorregrotaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torregrota",
		TitleName: "托雷格罗塔",
		TitleCode: "b_torregrota",
	}
}
