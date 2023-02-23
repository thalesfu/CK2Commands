package cadota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安得悦EndereBarony struct {
	feud.BaseBarony
}

var BEndere安得悦 feud.Barony = &安得悦EndereBarony{}

func init() {
	f := BEndere安得悦.(*安得悦EndereBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "endere",
		TitleName: "安得悦",
		TitleCode: "b_endere",
	}
}
