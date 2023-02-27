package irtek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔萨诺沃KirsanovoBarony struct {
	feud.BaseBarony
}

var BKirsanovo基尔萨诺沃 feud.Barony = &基尔萨诺沃KirsanovoBarony{}

func init() {
    f := BKirsanovo基尔萨诺沃.(*基尔萨诺沃KirsanovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirsanovo",
		TitleName: "基尔萨诺沃",
		TitleCode: "b_kirsanovo",
	}
}
