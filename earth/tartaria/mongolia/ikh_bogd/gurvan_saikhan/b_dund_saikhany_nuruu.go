package gurvan_saikhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 敦德Dund_saikhany_nuruuBarony struct {
	feud.BaseBarony
}

var BDund_saikhany_nuruu敦德 feud.Barony = &敦德Dund_saikhany_nuruuBarony{}

func init() {
    f := BDund_saikhany_nuruu敦德.(*敦德Dund_saikhany_nuruuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dund_saikhany_nuruu",
		TitleName: "敦德",
		TitleCode: "b_dund_saikhany_nuruu",
	}
}
