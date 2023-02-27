package galich_mersky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡德KadyyBarony struct {
	feud.BaseBarony
}

var BKadyy卡德 feud.Barony = &卡德KadyyBarony{}

func init() {
    f := BKadyy卡德.(*卡德KadyyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kadyy",
		TitleName: "卡德",
		TitleCode: "b_kadyy",
	}
}
