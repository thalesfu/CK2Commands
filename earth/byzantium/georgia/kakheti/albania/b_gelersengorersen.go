package albania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格雷森格拉森GelersengorersenBarony struct {
	feud.BaseBarony
}

var BGelersengorersen格雷森格拉森 feud.Barony = &格雷森格拉森GelersengorersenBarony{}

func init() {
	f := BGelersengorersen格雷森格拉森.(*格雷森格拉森GelersengorersenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gelersengorersen",
		TitleName: "格雷森格拉森",
		TitleCode: "b_gelersengorersen",
	}
}
