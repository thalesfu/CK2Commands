package leeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 庞蒂弗拉克特PontefractBarony struct {
	feud.BaseBarony
}

var BPontefract庞蒂弗拉克特 feud.Barony = &庞蒂弗拉克特PontefractBarony{}

func init() {
	f := BPontefract庞蒂弗拉克特.(*庞蒂弗拉克特PontefractBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pontefract",
		TitleName: "庞蒂弗拉克特",
		TitleCode: "b_pontefract",
	}
}
