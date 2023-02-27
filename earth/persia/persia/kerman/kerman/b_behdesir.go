package kerman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比哈德斯尔BehdesirBarony struct {
	feud.BaseBarony
}

var BBehdesir比哈德斯尔 feud.Barony = &比哈德斯尔BehdesirBarony{}

func init() {
    f := BBehdesir比哈德斯尔.(*比哈德斯尔BehdesirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "behdesir",
		TitleName: "比哈德斯尔",
		TitleCode: "b_behdesir",
	}
}
