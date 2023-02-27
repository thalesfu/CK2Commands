package znojmo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛乌卡LoukaBarony struct {
	feud.BaseBarony
}

var BLouka洛乌卡 feud.Barony = &洛乌卡LoukaBarony{}

func init() {
    f := BLouka洛乌卡.(*洛乌卡LoukaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "louka",
		TitleName: "洛乌卡",
		TitleCode: "b_louka",
	}
}
