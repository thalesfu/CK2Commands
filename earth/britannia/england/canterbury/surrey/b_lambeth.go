package surrey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗伯斯LambethBarony struct {
	feud.BaseBarony
}

var BLambeth朗伯斯 feud.Barony = &朗伯斯LambethBarony{}

func init() {
	f := BLambeth朗伯斯.(*朗伯斯LambethBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lambeth",
		TitleName: "朗伯斯",
		TitleCode: "b_lambeth",
	}
}
