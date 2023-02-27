package khinjali_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 契那阇梨KhinjaliBarony struct {
	feud.BaseBarony
}

var BKhinjali契那阇梨 feud.Barony = &契那阇梨KhinjaliBarony{}

func init() {
    f := BKhinjali契那阇梨.(*契那阇梨KhinjaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khinjali",
		TitleName: "契那阇梨",
		TitleCode: "b_khinjali",
	}
}
