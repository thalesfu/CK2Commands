package surrey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃金WokingBarony struct {
	feud.BaseBarony
}

var BWoking沃金 feud.Barony = &沃金WokingBarony{}

func init() {
	f := BWoking沃金.(*沃金WokingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "woking",
		TitleName: "沃金",
		TitleCode: "b_woking",
	}
}
