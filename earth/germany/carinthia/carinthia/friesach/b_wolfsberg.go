package friesach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔夫斯堡WolfsbergBarony struct {
	feud.BaseBarony
}

var BWolfsberg沃尔夫斯堡 feud.Barony = &沃尔夫斯堡WolfsbergBarony{}

func init() {
	f := BWolfsberg沃尔夫斯堡.(*沃尔夫斯堡WolfsbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wolfsberg",
		TitleName: "沃尔夫斯堡",
		TitleCode: "b_wolfsberg",
	}
}
