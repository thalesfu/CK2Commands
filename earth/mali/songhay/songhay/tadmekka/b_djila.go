package tadmekka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉拉DjilaBarony struct {
	feud.BaseBarony
}

var BDjila吉拉 feud.Barony = &吉拉DjilaBarony{}

func init() {
    f := BDjila吉拉.(*吉拉DjilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "djila",
		TitleName: "吉拉",
		TitleCode: "b_djila",
	}
}
