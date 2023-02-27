package khiva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基克肯KizketkenBarony struct {
	feud.BaseBarony
}

var BKizketken基克肯 feud.Barony = &基克肯KizketkenBarony{}

func init() {
    f := BKizketken基克肯.(*基克肯KizketkenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kizketken",
		TitleName: "基克肯",
		TitleCode: "b_kizketken",
	}
}
