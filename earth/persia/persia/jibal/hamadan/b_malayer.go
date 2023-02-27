package hamadan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马拉耶尔MalayerBarony struct {
	feud.BaseBarony
}

var BMalayer马拉耶尔 feud.Barony = &马拉耶尔MalayerBarony{}

func init() {
    f := BMalayer马拉耶尔.(*马拉耶尔MalayerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malayer",
		TitleName: "马拉耶尔",
		TitleCode: "b_malayer",
	}
}
