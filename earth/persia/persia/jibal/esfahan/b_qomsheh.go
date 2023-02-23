package esfahan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰什哈QomshehBarony struct {
	feud.BaseBarony
}

var BQomsheh恰什哈 feud.Barony = &恰什哈QomshehBarony{}

func init() {
	f := BQomsheh恰什哈.(*恰什哈QomshehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qomsheh",
		TitleName: "恰什哈",
		TitleCode: "b_qomsheh",
	}
}
