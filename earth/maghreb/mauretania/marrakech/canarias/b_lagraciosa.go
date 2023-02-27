package canarias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉西奥萨LagraciosaBarony struct {
	feud.BaseBarony
}

var BLagraciosa格拉西奥萨 feud.Barony = &格拉西奥萨LagraciosaBarony{}

func init() {
    f := BLagraciosa格拉西奥萨.(*格拉西奥萨LagraciosaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lagraciosa",
		TitleName: "格拉西奥萨",
		TitleCode: "b_lagraciosa",
	}
}
