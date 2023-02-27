package faiyum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法尤姆FaiyumBarony struct {
	feud.BaseBarony
}

var BFaiyum法尤姆 feud.Barony = &法尤姆FaiyumBarony{}

func init() {
    f := BFaiyum法尤姆.(*法尤姆FaiyumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "faiyum",
		TitleName: "法尤姆",
		TitleCode: "b_faiyum",
	}
}
