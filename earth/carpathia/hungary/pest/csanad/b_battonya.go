package csanad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲍托尼奥BattonyaBarony struct {
	feud.BaseBarony
}

var BBattonya鲍托尼奥 feud.Barony = &鲍托尼奥BattonyaBarony{}

func init() {
    f := BBattonya鲍托尼奥.(*鲍托尼奥BattonyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "battonya",
		TitleName: "鲍托尼奥",
		TitleCode: "b_battonya",
	}
}
