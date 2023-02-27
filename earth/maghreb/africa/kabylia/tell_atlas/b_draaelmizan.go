package tell_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德拉米赞DraaelmizanBarony struct {
	feud.BaseBarony
}

var BDraaelmizan德拉米赞 feud.Barony = &德拉米赞DraaelmizanBarony{}

func init() {
    f := BDraaelmizan德拉米赞.(*德拉米赞DraaelmizanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "draaelmizan",
		TitleName: "德拉米赞",
		TitleCode: "b_draaelmizan",
	}
}
