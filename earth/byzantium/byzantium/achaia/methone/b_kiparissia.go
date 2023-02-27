package methone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基帕里西亚KiparissiaBarony struct {
	feud.BaseBarony
}

var BKiparissia基帕里西亚 feud.Barony = &基帕里西亚KiparissiaBarony{}

func init() {
    f := BKiparissia基帕里西亚.(*基帕里西亚KiparissiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiparissia",
		TitleName: "基帕里西亚",
		TitleCode: "b_kiparissia",
	}
}
