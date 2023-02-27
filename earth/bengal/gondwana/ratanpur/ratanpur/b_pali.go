package ratanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尼PaliBarony struct {
	feud.BaseBarony
}

var BPali波尼 feud.Barony = &波尼PaliBarony{}

func init() {
    f := BPali波尼.(*波尼PaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pali",
		TitleName: "波尼",
		TitleCode: "b_pali",
	}
}
