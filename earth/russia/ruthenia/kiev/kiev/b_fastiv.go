package kiev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法斯托夫FastivBarony struct {
	feud.BaseBarony
}

var BFastiv法斯托夫 feud.Barony = &法斯托夫FastivBarony{}

func init() {
	f := BFastiv法斯托夫.(*法斯托夫FastivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fastiv",
		TitleName: "法斯托夫",
		TitleCode: "b_fastiv",
	}
}
