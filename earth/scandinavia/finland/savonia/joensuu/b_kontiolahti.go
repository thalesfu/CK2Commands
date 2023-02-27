package joensuu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔蒂奥拉赫蒂KontiolahtiBarony struct {
	feud.BaseBarony
}

var BKontiolahti孔蒂奥拉赫蒂 feud.Barony = &孔蒂奥拉赫蒂KontiolahtiBarony{}

func init() {
    f := BKontiolahti孔蒂奥拉赫蒂.(*孔蒂奥拉赫蒂KontiolahtiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kontiolahti",
		TitleName: "孔蒂奥拉赫蒂",
		TitleCode: "b_kontiolahti",
	}
}
