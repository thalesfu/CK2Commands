package chlisselbourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅夫杜斯特罗伊NevdubstroyBarony struct {
	feud.BaseBarony
}

var BNevdubstroy涅夫杜斯特罗伊 feud.Barony = &涅夫杜斯特罗伊NevdubstroyBarony{}

func init() {
    f := BNevdubstroy涅夫杜斯特罗伊.(*涅夫杜斯特罗伊NevdubstroyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nevdubstroy",
		TitleName: "涅夫杜斯特罗伊",
		TitleCode: "b_nevdubstroy",
	}
}
