package tikhvine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰戈达ChagodaBarony struct {
	feud.BaseBarony
}

var BChagoda恰戈达 feud.Barony = &恰戈达ChagodaBarony{}

func init() {
    f := BChagoda恰戈达.(*恰戈达ChagodaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chagoda",
		TitleName: "恰戈达",
		TitleCode: "b_chagoda",
	}
}
