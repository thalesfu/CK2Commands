package pelusia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞拉比尤姆SerapeumBarony struct {
	feud.BaseBarony
}

var BSerapeum塞拉比尤姆 feud.Barony = &塞拉比尤姆SerapeumBarony{}

func init() {
    f := BSerapeum塞拉比尤姆.(*塞拉比尤姆SerapeumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "serapeum",
		TitleName: "塞拉比尤姆",
		TitleCode: "b_serapeum",
	}
}
