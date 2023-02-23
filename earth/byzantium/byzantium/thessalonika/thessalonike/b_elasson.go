package thessalonike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃拉索纳ElassonBarony struct {
	feud.BaseBarony
}

var BElasson埃拉索纳 feud.Barony = &埃拉索纳ElassonBarony{}

func init() {
	f := BElasson埃拉索纳.(*埃拉索纳ElassonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elasson",
		TitleName: "埃拉索纳",
		TitleCode: "b_elasson",
	}
}
