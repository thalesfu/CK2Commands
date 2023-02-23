package polotsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多尔扎DouzaBarony struct {
	feud.BaseBarony
}

var BDouza多尔扎 feud.Barony = &多尔扎DouzaBarony{}

func init() {
	f := BDouza多尔扎.(*多尔扎DouzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "douza",
		TitleName: "多尔扎",
		TitleCode: "b_douza",
	}
}
