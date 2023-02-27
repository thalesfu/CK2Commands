package candradvipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘罗那KhulnaBarony struct {
	feud.BaseBarony
}

var BKhulna丘罗那 feud.Barony = &丘罗那KhulnaBarony{}

func init() {
    f := BKhulna丘罗那.(*丘罗那KhulnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khulna",
		TitleName: "丘罗那",
		TitleCode: "b_khulna",
	}
}
