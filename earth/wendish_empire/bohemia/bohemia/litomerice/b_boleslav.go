package litomerice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博莱斯拉夫BoleslavBarony struct {
	feud.BaseBarony
}

var BBoleslav博莱斯拉夫 feud.Barony = &博莱斯拉夫BoleslavBarony{}

func init() {
    f := BBoleslav博莱斯拉夫.(*博莱斯拉夫BoleslavBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boleslav",
		TitleName: "博莱斯拉夫",
		TitleCode: "b_boleslav",
	}
}
