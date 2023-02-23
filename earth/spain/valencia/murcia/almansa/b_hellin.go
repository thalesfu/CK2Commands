package almansa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃林HellinBarony struct {
	feud.BaseBarony
}

var BHellin埃林 feud.Barony = &埃林HellinBarony{}

func init() {
	f := BHellin埃林.(*埃林HellinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hellin",
		TitleName: "埃林",
		TitleCode: "b_hellin",
	}
}
