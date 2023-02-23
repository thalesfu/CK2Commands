package orania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈贾德MerselhadjadBarony struct {
	feud.BaseBarony
}

var BMerselhadjad哈贾德 feud.Barony = &哈贾德MerselhadjadBarony{}

func init() {
	f := BMerselhadjad哈贾德.(*哈贾德MerselhadjadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merselhadjad",
		TitleName: "哈贾德",
		TitleCode: "b_merselhadjad",
	}
}
