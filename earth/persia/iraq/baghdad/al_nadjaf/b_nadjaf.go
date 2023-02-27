package al_nadjaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳杰夫NadjafBarony struct {
	feud.BaseBarony
}

var BNadjaf纳杰夫 feud.Barony = &纳杰夫NadjafBarony{}

func init() {
    f := BNadjaf纳杰夫.(*纳杰夫NadjafBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nadjaf",
		TitleName: "纳杰夫",
		TitleCode: "b_nadjaf",
	}
}
