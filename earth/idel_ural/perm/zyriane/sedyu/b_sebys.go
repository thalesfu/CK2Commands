package sedyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢贝西SebysBarony struct {
	feud.BaseBarony
}

var BSebys谢贝西 feud.Barony = &谢贝西SebysBarony{}

func init() {
    f := BSebys谢贝西.(*谢贝西SebysBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sebys",
		TitleName: "谢贝西",
		TitleCode: "b_sebys",
	}
}
