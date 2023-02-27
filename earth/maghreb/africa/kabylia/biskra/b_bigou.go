package biskra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比勾BigouBarony struct {
	feud.BaseBarony
}

var BBigou比勾 feud.Barony = &比勾BigouBarony{}

func init() {
    f := BBigou比勾.(*比勾BigouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bigou",
		TitleName: "比勾",
		TitleCode: "b_bigou",
	}
}
