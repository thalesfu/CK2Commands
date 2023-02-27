package oppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利勒哈默尔LillehammerBarony struct {
	feud.BaseBarony
}

var BLillehammer利勒哈默尔 feud.Barony = &利勒哈默尔LillehammerBarony{}

func init() {
    f := BLillehammer利勒哈默尔.(*利勒哈默尔LillehammerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lillehammer",
		TitleName: "利勒哈默尔",
		TitleCode: "b_lillehammer",
	}
}
