package sambhal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希尔西SirsiBarony struct {
	feud.BaseBarony
}

var BSirsi希尔西 feud.Barony = &希尔西SirsiBarony{}

func init() {
    f := BSirsi希尔西.(*希尔西SirsiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sirsi",
		TitleName: "希尔西",
		TitleCode: "b_sirsi",
	}
}
