package hesse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅伦贝格MerenbergBarony struct {
	feud.BaseBarony
}

var BMerenberg梅伦贝格 feud.Barony = &梅伦贝格MerenbergBarony{}

func init() {
	f := BMerenberg梅伦贝格.(*梅伦贝格MerenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merenberg",
		TitleName: "梅伦贝格",
		TitleCode: "b_merenberg",
	}
}
