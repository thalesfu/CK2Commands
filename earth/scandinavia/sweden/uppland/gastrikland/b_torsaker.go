package gastrikland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托绍克尔TorsakerBarony struct {
	feud.BaseBarony
}

var BTorsaker托绍克尔 feud.Barony = &托绍克尔TorsakerBarony{}

func init() {
    f := BTorsaker托绍克尔.(*托绍克尔TorsakerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torsaker",
		TitleName: "托绍克尔",
		TitleCode: "b_torsaker",
	}
}
