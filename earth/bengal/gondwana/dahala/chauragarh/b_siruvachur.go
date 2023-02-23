package chauragarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯卢伐邱SiruvachurBarony struct {
	feud.BaseBarony
}

var BSiruvachur斯卢伐邱 feud.Barony = &斯卢伐邱SiruvachurBarony{}

func init() {
	f := BSiruvachur斯卢伐邱.(*斯卢伐邱SiruvachurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siruvachur",
		TitleName: "斯卢伐邱",
		TitleCode: "b_siruvachur",
	}
}
