package vanema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔森TalsenBarony struct {
	feud.BaseBarony
}

var BTalsen塔尔森 feud.Barony = &塔尔森TalsenBarony{}

func init() {
    f := BTalsen塔尔森.(*塔尔森TalsenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talsen",
		TitleName: "塔尔森",
		TitleCode: "b_talsen",
	}
}
