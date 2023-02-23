package tashkurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰特QiateBarony struct {
	feud.BaseBarony
}

var BQiate恰特 feud.Barony = &恰特QiateBarony{}

func init() {
	f := BQiate恰特.(*恰特QiateBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qiate",
		TitleName: "恰特",
		TitleCode: "b_qiate",
	}
}
