package consenza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切伦齐亚CerenziaBarony struct {
	feud.BaseBarony
}

var BCerenzia切伦齐亚 feud.Barony = &切伦齐亚CerenziaBarony{}

func init() {
    f := BCerenzia切伦齐亚.(*切伦齐亚CerenziaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cerenzia",
		TitleName: "切伦齐亚",
		TitleCode: "b_cerenzia",
	}
}
