package sedyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢季尤SedyuBarony struct {
	feud.BaseBarony
}

var BSedyu谢季尤 feud.Barony = &谢季尤SedyuBarony{}

func init() {
    f := BSedyu谢季尤.(*谢季尤SedyuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sedyu",
		TitleName: "谢季尤",
		TitleCode: "b_sedyu",
	}
}
