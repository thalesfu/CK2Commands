package litomerice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 慕尼黑城堡Mnichovo_hradisteBarony struct {
	feud.BaseBarony
}

var BMnichovo_hradiste慕尼黑城堡 feud.Barony = &慕尼黑城堡Mnichovo_hradisteBarony{}

func init() {
    f := BMnichovo_hradiste慕尼黑城堡.(*慕尼黑城堡Mnichovo_hradisteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mnichovo_hradiste",
		TitleName: "慕尼黑城堡",
		TitleCode: "b_mnichovo_hradiste",
	}
}
