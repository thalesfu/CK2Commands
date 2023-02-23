package montpellier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞兹河畔巴尼奥勒BagnolssurcezeBarony struct {
	feud.BaseBarony
}

var BBagnolssurceze塞兹河畔巴尼奥勒 feud.Barony = &塞兹河畔巴尼奥勒BagnolssurcezeBarony{}

func init() {
	f := BBagnolssurceze塞兹河畔巴尼奥勒.(*塞兹河畔巴尼奥勒BagnolssurcezeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bagnolssurceze",
		TitleName: "塞兹河畔巴尼奥勒",
		TitleCode: "b_bagnolssurceze",
	}
}
