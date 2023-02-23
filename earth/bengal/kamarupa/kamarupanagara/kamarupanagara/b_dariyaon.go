package kamarupanagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀利耶乌那DariyaonBarony struct {
	feud.BaseBarony
}

var BDariyaon陀利耶乌那 feud.Barony = &陀利耶乌那DariyaonBarony{}

func init() {
	f := BDariyaon陀利耶乌那.(*陀利耶乌那DariyaonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dariyaon",
		TitleName: "陀利耶乌那",
		TitleCode: "b_dariyaon",
	}
}
