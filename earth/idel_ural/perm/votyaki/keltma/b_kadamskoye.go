package keltma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡达姆斯科耶KadamskoyeBarony struct {
	feud.BaseBarony
}

var BKadamskoye卡达姆斯科耶 feud.Barony = &卡达姆斯科耶KadamskoyeBarony{}

func init() {
    f := BKadamskoye卡达姆斯科耶.(*卡达姆斯科耶KadamskoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kadamskoye",
		TitleName: "卡达姆斯科耶",
		TitleCode: "b_kadamskoye",
	}
}
