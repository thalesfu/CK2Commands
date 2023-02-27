package desmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 顿那瑟德DunaseadBarony struct {
	feud.BaseBarony
}

var BDunasead顿那瑟德 feud.Barony = &顿那瑟德DunaseadBarony{}

func init() {
    f := BDunasead顿那瑟德.(*顿那瑟德DunaseadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunasead",
		TitleName: "顿那瑟德",
		TitleCode: "b_dunasead",
	}
}
