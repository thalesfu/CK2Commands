package kotthasara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耆厘建荼迦GirikandakaBarony struct {
	feud.BaseBarony
}

var BGirikandaka耆厘建荼迦 feud.Barony = &耆厘建荼迦GirikandakaBarony{}

func init() {
    f := BGirikandaka耆厘建荼迦.(*耆厘建荼迦GirikandakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "girikandaka",
		TitleName: "耆厘建荼迦",
		TitleCode: "b_girikandaka",
	}
}
