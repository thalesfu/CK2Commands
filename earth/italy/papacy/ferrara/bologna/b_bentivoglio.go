package bologna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 本蒂沃利奥BentivoglioBarony struct {
	feud.BaseBarony
}

var BBentivoglio本蒂沃利奥 feud.Barony = &本蒂沃利奥BentivoglioBarony{}

func init() {
	f := BBentivoglio本蒂沃利奥.(*本蒂沃利奥BentivoglioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bentivoglio",
		TitleName: "本蒂沃利奥",
		TitleCode: "b_bentivoglio",
	}
}
