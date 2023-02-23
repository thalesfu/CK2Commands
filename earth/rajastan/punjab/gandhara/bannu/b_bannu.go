package bannu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 般奴BannuBarony struct {
	feud.BaseBarony
}

var BBannu般奴 feud.Barony = &般奴BannuBarony{}

func init() {
	f := BBannu般奴.(*般奴BannuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bannu",
		TitleName: "般奴",
		TitleCode: "b_bannu",
	}
}
