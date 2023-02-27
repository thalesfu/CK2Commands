package vijayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩诃末那揭罗MohammadnagarBarony struct {
	feud.BaseBarony
}

var BMohammadnagar摩诃末那揭罗 feud.Barony = &摩诃末那揭罗MohammadnagarBarony{}

func init() {
    f := BMohammadnagar摩诃末那揭罗.(*摩诃末那揭罗MohammadnagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mohammadnagar",
		TitleName: "摩诃末那揭罗",
		TitleCode: "b_mohammadnagar",
	}
}
