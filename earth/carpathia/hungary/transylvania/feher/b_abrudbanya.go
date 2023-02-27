package feher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥布鲁德巴尼奥AbrudbanyaBarony struct {
	feud.BaseBarony
}

var BAbrudbanya奥布鲁德巴尼奥 feud.Barony = &奥布鲁德巴尼奥AbrudbanyaBarony{}

func init() {
    f := BAbrudbanya奥布鲁德巴尼奥.(*奥布鲁德巴尼奥AbrudbanyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abrudbanya",
		TitleName: "奥布鲁德巴尼奥",
		TitleCode: "b_abrudbanya",
	}
}
