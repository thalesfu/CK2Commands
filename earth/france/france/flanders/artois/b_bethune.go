package artois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝蒂讷BethuneBarony struct {
	feud.BaseBarony
}

var BBethune贝蒂讷 feud.Barony = &贝蒂讷BethuneBarony{}

func init() {
    f := BBethune贝蒂讷.(*贝蒂讷BethuneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bethune",
		TitleName: "贝蒂讷",
		TitleCode: "b_bethune",
	}
}
