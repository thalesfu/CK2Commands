package halsingland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 松德海德SundhedBarony struct {
	feud.BaseBarony
}

var BSundhed松德海德 feud.Barony = &松德海德SundhedBarony{}

func init() {
	f := BSundhed松德海德.(*松德海德SundhedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sundhed",
		TitleName: "松德海德",
		TitleCode: "b_sundhed",
	}
}
