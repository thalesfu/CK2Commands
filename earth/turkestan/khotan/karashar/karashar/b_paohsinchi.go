package karashar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲍兴集PaohsinchiBarony struct {
	feud.BaseBarony
}

var BPaohsinchi鲍兴集 feud.Barony = &鲍兴集PaohsinchiBarony{}

func init() {
	f := BPaohsinchi鲍兴集.(*鲍兴集PaohsinchiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paohsinchi",
		TitleName: "鲍兴集",
		TitleCode: "b_paohsinchi",
	}
}
