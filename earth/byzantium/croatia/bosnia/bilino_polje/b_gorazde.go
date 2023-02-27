package bilino_polje

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈拉日代GorazdeBarony struct {
	feud.BaseBarony
}

var BGorazde戈拉日代 feud.Barony = &戈拉日代GorazdeBarony{}

func init() {
    f := BGorazde戈拉日代.(*戈拉日代GorazdeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gorazde",
		TitleName: "戈拉日代",
		TitleCode: "b_gorazde",
	}
}
