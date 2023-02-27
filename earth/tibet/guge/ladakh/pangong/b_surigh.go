package pangong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨利吉勒SurighBarony struct {
	feud.BaseBarony
}

var BSurigh萨利吉勒 feud.Barony = &萨利吉勒SurighBarony{}

func init() {
    f := BSurigh萨利吉勒.(*萨利吉勒SurighBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "surigh",
		TitleName: "萨利吉勒",
		TitleCode: "b_surigh",
	}
}
