package korchev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎维特尼ZavitneBarony struct {
	feud.BaseBarony
}

var BZavitne扎维特尼 feud.Barony = &扎维特尼ZavitneBarony{}

func init() {
    f := BZavitne扎维特尼.(*扎维特尼ZavitneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zavitne",
		TitleName: "扎维特尼",
		TitleCode: "b_zavitne",
	}
}
