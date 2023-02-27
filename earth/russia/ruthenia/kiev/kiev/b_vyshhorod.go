package kiev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维什戈罗德VyshhorodBarony struct {
	feud.BaseBarony
}

var BVyshhorod维什戈罗德 feud.Barony = &维什戈罗德VyshhorodBarony{}

func init() {
    f := BVyshhorod维什戈罗德.(*维什戈罗德VyshhorodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vyshhorod",
		TitleName: "维什戈罗德",
		TitleCode: "b_vyshhorod",
	}
}
