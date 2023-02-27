package santiago

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝林VerinBarony struct {
	feud.BaseBarony
}

var BVerin贝林 feud.Barony = &贝林VerinBarony{}

func init() {
    f := BVerin贝林.(*贝林VerinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "verin",
		TitleName: "贝林",
		TitleCode: "b_verin",
	}
}
