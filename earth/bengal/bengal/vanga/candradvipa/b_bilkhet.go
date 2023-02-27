package candradvipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗吉多BilkhetBarony struct {
	feud.BaseBarony
}

var BBilkhet毗吉多 feud.Barony = &毗吉多BilkhetBarony{}

func init() {
    f := BBilkhet毗吉多.(*毗吉多BilkhetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilkhet",
		TitleName: "毗吉多",
		TitleCode: "b_bilkhet",
	}
}
