package sinope

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 忒弥斯库拉ThemiscyraBarony struct {
	feud.BaseBarony
}

var BThemiscyra忒弥斯库拉 feud.Barony = &忒弥斯库拉ThemiscyraBarony{}

func init() {
    f := BThemiscyra忒弥斯库拉.(*忒弥斯库拉ThemiscyraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "themiscyra",
		TitleName: "忒弥斯库拉",
		TitleCode: "b_themiscyra",
	}
}
