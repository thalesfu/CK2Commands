package galindia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特罗伊堡TreuburgBarony struct {
	feud.BaseBarony
}

var BTreuburg特罗伊堡 feud.Barony = &特罗伊堡TreuburgBarony{}

func init() {
    f := BTreuburg特罗伊堡.(*特罗伊堡TreuburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "treuburg",
		TitleName: "特罗伊堡",
		TitleCode: "b_treuburg",
	}
}
