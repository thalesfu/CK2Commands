package vastergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法尔雪平FalkopingBarony struct {
	feud.BaseBarony
}

var BFalkoping法尔雪平 feud.Barony = &法尔雪平FalkopingBarony{}

func init() {
    f := BFalkoping法尔雪平.(*法尔雪平FalkopingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "falkoping",
		TitleName: "法尔雪平",
		TitleCode: "b_falkoping",
	}
}
