package leicester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索斯韦尔SouthwellBarony struct {
	feud.BaseBarony
}

var BSouthwell索斯韦尔 feud.Barony = &索斯韦尔SouthwellBarony{}

func init() {
	f := BSouthwell索斯韦尔.(*索斯韦尔SouthwellBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "southwell",
		TitleName: "索斯韦尔",
		TitleCode: "b_southwell",
	}
}
