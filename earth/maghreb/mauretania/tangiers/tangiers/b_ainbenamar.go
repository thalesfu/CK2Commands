package tangiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因本阿马尔AinbenamarBarony struct {
	feud.BaseBarony
}

var BAinbenamar艾因本阿马尔 feud.Barony = &艾因本阿马尔AinbenamarBarony{}

func init() {
	f := BAinbenamar艾因本阿马尔.(*艾因本阿马尔AinbenamarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ainbenamar",
		TitleName: "艾因本阿马尔",
		TitleCode: "b_ainbenamar",
	}
}
