package el_rif

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳祖尔NadorBarony struct {
	feud.BaseBarony
}

var BNador纳祖尔 feud.Barony = &纳祖尔NadorBarony{}

func init() {
    f := BNador纳祖尔.(*纳祖尔NadorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nador",
		TitleName: "纳祖尔",
		TitleCode: "b_nador",
	}
}
