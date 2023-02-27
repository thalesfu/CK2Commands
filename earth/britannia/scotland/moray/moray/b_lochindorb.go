package moray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛欣多布LochindorbBarony struct {
	feud.BaseBarony
}

var BLochindorb洛欣多布 feud.Barony = &洛欣多布LochindorbBarony{}

func init() {
    f := BLochindorb洛欣多布.(*洛欣多布LochindorbBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lochindorb",
		TitleName: "洛欣多布",
		TitleCode: "b_lochindorb",
	}
}
