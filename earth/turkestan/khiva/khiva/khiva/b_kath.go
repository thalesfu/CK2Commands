package khiva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯KathBarony struct {
	feud.BaseBarony
}

var BKath卡斯 feud.Barony = &卡斯KathBarony{}

func init() {
    f := BKath卡斯.(*卡斯KathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kath",
		TitleName: "卡斯",
		TitleCode: "b_kath",
	}
}
