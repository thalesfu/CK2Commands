package jumla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 钱普尔ChainpurBarony struct {
	feud.BaseBarony
}

var BChainpur钱普尔 feud.Barony = &钱普尔ChainpurBarony{}

func init() {
    f := BChainpur钱普尔.(*钱普尔ChainpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chainpur",
		TitleName: "钱普尔",
		TitleCode: "b_chainpur",
	}
}
