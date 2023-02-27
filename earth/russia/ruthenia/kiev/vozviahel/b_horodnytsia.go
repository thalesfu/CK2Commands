package vozviahel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈罗德尼齐亚HorodnytsiaBarony struct {
	feud.BaseBarony
}

var BHorodnytsia戈罗德尼齐亚 feud.Barony = &戈罗德尼齐亚HorodnytsiaBarony{}

func init() {
    f := BHorodnytsia戈罗德尼齐亚.(*戈罗德尼齐亚HorodnytsiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "horodnytsia",
		TitleName: "戈罗德尼齐亚",
		TitleCode: "b_horodnytsia",
	}
}
