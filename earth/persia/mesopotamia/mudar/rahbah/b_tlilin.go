package rahbah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特利林TlilinBarony struct {
	feud.BaseBarony
}

var BTlilin特利林 feud.Barony = &特利林TlilinBarony{}

func init() {
    f := BTlilin特利林.(*特利林TlilinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tlilin",
		TitleName: "特利林",
		TitleCode: "b_tlilin",
	}
}
