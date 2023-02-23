package belgrade

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯梅代雷沃SmederevoBarony struct {
	feud.BaseBarony
}

var BSmederevo斯梅代雷沃 feud.Barony = &斯梅代雷沃SmederevoBarony{}

func init() {
	f := BSmederevo斯梅代雷沃.(*斯梅代雷沃SmederevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "smederevo",
		TitleName: "斯梅代雷沃",
		TitleCode: "b_smederevo",
	}
}
