package ouadane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿祖吉AzuggiBarony struct {
	feud.BaseBarony
}

var BAzuggi阿祖吉 feud.Barony = &阿祖吉AzuggiBarony{}

func init() {
    f := BAzuggi阿祖吉.(*阿祖吉AzuggiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azuggi",
		TitleName: "阿祖吉",
		TitleCode: "b_azuggi",
	}
}
