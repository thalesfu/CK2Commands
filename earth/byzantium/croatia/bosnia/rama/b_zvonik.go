package rama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹沃尼克ZvonikBarony struct {
	feud.BaseBarony
}

var BZvonik兹沃尼克 feud.Barony = &兹沃尼克ZvonikBarony{}

func init() {
	f := BZvonik兹沃尼克.(*兹沃尼克ZvonikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zvonik",
		TitleName: "兹沃尼克",
		TitleCode: "b_zvonik",
	}
}
