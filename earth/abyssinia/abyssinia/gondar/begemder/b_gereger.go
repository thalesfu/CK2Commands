package begemder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格雷格尔GeregerBarony struct {
	feud.BaseBarony
}

var BGereger格雷格尔 feud.Barony = &格雷格尔GeregerBarony{}

func init() {
	f := BGereger格雷格尔.(*格雷格尔GeregerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gereger",
		TitleName: "格雷格尔",
		TitleCode: "b_gereger",
	}
}
