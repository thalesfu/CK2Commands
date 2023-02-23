package saptagrama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 般娑吠利阿BansberiaBarony struct {
	feud.BaseBarony
}

var BBansberia般娑吠利阿 feud.Barony = &般娑吠利阿BansberiaBarony{}

func init() {
	f := BBansberia般娑吠利阿.(*般娑吠利阿BansberiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bansberia",
		TitleName: "般娑吠利阿",
		TitleCode: "b_bansberia",
	}
}
