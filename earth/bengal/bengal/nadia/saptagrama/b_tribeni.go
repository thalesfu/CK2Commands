package saptagrama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帝利吠尼TribeniBarony struct {
	feud.BaseBarony
}

var BTribeni帝利吠尼 feud.Barony = &帝利吠尼TribeniBarony{}

func init() {
    f := BTribeni帝利吠尼.(*帝利吠尼TribeniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tribeni",
		TitleName: "帝利吠尼",
		TitleCode: "b_tribeni",
	}
}
