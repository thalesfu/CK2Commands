package ascalon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格伦AgelenBarony struct {
	feud.BaseBarony
}

var BAgelen阿格伦 feud.Barony = &阿格伦AgelenBarony{}

func init() {
	f := BAgelen阿格伦.(*阿格伦AgelenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agelen",
		TitleName: "阿格伦",
		TitleCode: "b_agelen",
	}
}
