package tahoua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁鲁姆RurumBarony struct {
	feud.BaseBarony
}

var BRurum鲁鲁姆 feud.Barony = &鲁鲁姆RurumBarony{}

func init() {
    f := BRurum鲁鲁姆.(*鲁鲁姆RurumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rurum",
		TitleName: "鲁鲁姆",
		TitleCode: "b_rurum",
	}
}
