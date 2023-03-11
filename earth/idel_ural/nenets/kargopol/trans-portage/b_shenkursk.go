package trans-portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 申库尔斯克ShenkurskBarony struct {
	feud.BaseBarony
}

var BShenkursk申库尔斯克 feud.Barony = &申库尔斯克ShenkurskBarony{}

func init() {
    f := BShenkursk申库尔斯克.(*申库尔斯克ShenkurskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shenkursk",
		TitleName: "申库尔斯克",
		TitleCode: "b_shenkursk",
	}
}
