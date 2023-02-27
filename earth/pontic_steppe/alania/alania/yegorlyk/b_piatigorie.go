package yegorlyk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮亚蒂戈里耶PiatigorieBarony struct {
	feud.BaseBarony
}

var BPiatigorie皮亚蒂戈里耶 feud.Barony = &皮亚蒂戈里耶PiatigorieBarony{}

func init() {
    f := BPiatigorie皮亚蒂戈里耶.(*皮亚蒂戈里耶PiatigorieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "piatigorie",
		TitleName: "皮亚蒂戈里耶",
		TitleCode: "b_piatigorie",
	}
}
