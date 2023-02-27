package chunar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达乌德布尔DaudpurBarony struct {
	feud.BaseBarony
}

var BDaudpur达乌德布尔 feud.Barony = &达乌德布尔DaudpurBarony{}

func init() {
    f := BDaudpur达乌德布尔.(*达乌德布尔DaudpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daudpur",
		TitleName: "达乌德布尔",
		TitleCode: "b_daudpur",
	}
}
