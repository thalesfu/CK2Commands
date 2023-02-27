package xainza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佳琼JakhyungBarony struct {
	feud.BaseBarony
}

var BJakhyung佳琼 feud.Barony = &佳琼JakhyungBarony{}

func init() {
    f := BJakhyung佳琼.(*佳琼JakhyungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jakhyung",
		TitleName: "佳琼",
		TitleCode: "b_jakhyung",
	}
}
