package anglesey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍利黑德HolyheadBarony struct {
	feud.BaseBarony
}

var BHolyhead霍利黑德 feud.Barony = &霍利黑德HolyheadBarony{}

func init() {
	f := BHolyhead霍利黑德.(*霍利黑德HolyheadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "holyhead",
		TitleName: "霍利黑德",
		TitleCode: "b_holyhead",
	}
}
