package galich_mersky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加利奇梅尔斯基GalichmerskyBarony struct {
	feud.BaseBarony
}

var BGalichmersky加利奇梅尔斯基 feud.Barony = &加利奇梅尔斯基GalichmerskyBarony{}

func init() {
    f := BGalichmersky加利奇梅尔斯基.(*加利奇梅尔斯基GalichmerskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galichmersky",
		TitleName: "加利奇梅尔斯基",
		TitleCode: "b_galichmersky",
	}
}
