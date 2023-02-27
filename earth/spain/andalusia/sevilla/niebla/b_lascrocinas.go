package niebla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克罗西纳斯LascrocinasBarony struct {
	feud.BaseBarony
}

var BLascrocinas克罗西纳斯 feud.Barony = &克罗西纳斯LascrocinasBarony{}

func init() {
    f := BLascrocinas克罗西纳斯.(*克罗西纳斯LascrocinasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lascrocinas",
		TitleName: "克罗西纳斯",
		TitleCode: "b_lascrocinas",
	}
}
