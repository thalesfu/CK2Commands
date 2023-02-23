package hendjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝赫贝汉BehbahanBarony struct {
	feud.BaseBarony
}

var BBehbahan贝赫贝汉 feud.Barony = &贝赫贝汉BehbahanBarony{}

func init() {
	f := BBehbahan贝赫贝汉.(*贝赫贝汉BehbahanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "behbahan",
		TitleName: "贝赫贝汉",
		TitleCode: "b_behbahan",
	}
}
