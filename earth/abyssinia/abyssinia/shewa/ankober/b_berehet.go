package ankober

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔赫特BerehetBarony struct {
	feud.BaseBarony
}

var BBerehet贝尔赫特 feud.Barony = &贝尔赫特BerehetBarony{}

func init() {
	f := BBerehet贝尔赫特.(*贝尔赫特BerehetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berehet",
		TitleName: "贝尔赫特",
		TitleCode: "b_berehet",
	}
}
