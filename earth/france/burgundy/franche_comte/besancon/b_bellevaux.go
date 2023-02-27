package besancon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝勒沃BellevauxBarony struct {
	feud.BaseBarony
}

var BBellevaux贝勒沃 feud.Barony = &贝勒沃BellevauxBarony{}

func init() {
    f := BBellevaux贝勒沃.(*贝勒沃BellevauxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bellevaux",
		TitleName: "贝勒沃",
		TitleCode: "b_bellevaux",
	}
}
