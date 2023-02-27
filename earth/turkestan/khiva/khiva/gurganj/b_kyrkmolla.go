package gurganj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔克莫拉KyrkmollaBarony struct {
	feud.BaseBarony
}

var BKyrkmolla基尔克莫拉 feud.Barony = &基尔克莫拉KyrkmollaBarony{}

func init() {
    f := BKyrkmolla基尔克莫拉.(*基尔克莫拉KyrkmollaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyrkmolla",
		TitleName: "基尔克莫拉",
		TitleCode: "b_kyrkmolla",
	}
}
