package purang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贤柏林SimbilingBarony struct {
	feud.BaseBarony
}

var BSimbiling贤柏林 feud.Barony = &贤柏林SimbilingBarony{}

func init() {
    f := BSimbiling贤柏林.(*贤柏林SimbilingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "simbiling",
		TitleName: "贤柏林",
		TitleCode: "b_simbiling",
	}
}
