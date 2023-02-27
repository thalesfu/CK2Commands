package ossory

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格雷南GrennanBarony struct {
	feud.BaseBarony
}

var BGrennan格雷南 feud.Barony = &格雷南GrennanBarony{}

func init() {
    f := BGrennan格雷南.(*格雷南GrennanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grennan",
		TitleName: "格雷南",
		TitleCode: "b_grennan",
	}
}
