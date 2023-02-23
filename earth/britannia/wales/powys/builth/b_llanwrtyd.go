package builth

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉努蒂德LlanwrtydBarony struct {
	feud.BaseBarony
}

var BLlanwrtyd拉努蒂德 feud.Barony = &拉努蒂德LlanwrtydBarony{}

func init() {
	f := BLlanwrtyd拉努蒂德.(*拉努蒂德LlanwrtydBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llanwrtyd",
		TitleName: "拉努蒂德",
		TitleCode: "b_llanwrtyd",
	}
}
