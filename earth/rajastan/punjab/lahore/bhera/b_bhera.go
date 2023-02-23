package bhera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠罗BheraBarony struct {
	feud.BaseBarony
}

var BBhera吠罗 feud.Barony = &吠罗BheraBarony{}

func init() {
	f := BBhera吠罗.(*吠罗BheraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhera",
		TitleName: "吠罗",
		TitleCode: "b_bhera",
	}
}
