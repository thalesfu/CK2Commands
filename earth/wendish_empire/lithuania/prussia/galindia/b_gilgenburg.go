package galindia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉尔根堡GilgenburgBarony struct {
	feud.BaseBarony
}

var BGilgenburg吉尔根堡 feud.Barony = &吉尔根堡GilgenburgBarony{}

func init() {
    f := BGilgenburg吉尔根堡.(*吉尔根堡GilgenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gilgenburg",
		TitleName: "吉尔根堡",
		TitleCode: "b_gilgenburg",
	}
}
