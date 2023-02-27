package sundgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔特基什AltkirchBarony struct {
	feud.BaseBarony
}

var BAltkirch阿尔特基什 feud.Barony = &阿尔特基什AltkirchBarony{}

func init() {
    f := BAltkirch阿尔特基什.(*阿尔特基什AltkirchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "altkirch",
		TitleName: "阿尔特基什",
		TitleCode: "b_altkirch",
	}
}
