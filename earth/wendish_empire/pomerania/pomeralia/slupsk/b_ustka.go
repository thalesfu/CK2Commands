package slupsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯特卡UstkaBarony struct {
	feud.BaseBarony
}

var BUstka乌斯特卡 feud.Barony = &乌斯特卡UstkaBarony{}

func init() {
    f := BUstka乌斯特卡.(*乌斯特卡UstkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ustka",
		TitleName: "乌斯特卡",
		TitleCode: "b_ustka",
	}
}
