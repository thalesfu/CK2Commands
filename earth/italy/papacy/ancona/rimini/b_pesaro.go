package rimini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩萨罗PesaroBarony struct {
	feud.BaseBarony
}

var BPesaro佩萨罗 feud.Barony = &佩萨罗PesaroBarony{}

func init() {
	f := BPesaro佩萨罗.(*佩萨罗PesaroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pesaro",
		TitleName: "佩萨罗",
		TitleCode: "b_pesaro",
	}
}
