package wiltshire

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马姆斯伯里MalmesburyBarony struct {
	feud.BaseBarony
}

var BMalmesbury马姆斯伯里 feud.Barony = &马姆斯伯里MalmesburyBarony{}

func init() {
    f := BMalmesbury马姆斯伯里.(*马姆斯伯里MalmesburyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malmesbury",
		TitleName: "马姆斯伯里",
		TitleCode: "b_malmesbury",
	}
}
