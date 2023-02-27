package tirgoviste

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒姆尼库瑟拉特Ramnicu_saratBarony struct {
	feud.BaseBarony
}

var BRamnicu_sarat勒姆尼库瑟拉特 feud.Barony = &勒姆尼库瑟拉特Ramnicu_saratBarony{}

func init() {
    f := BRamnicu_sarat勒姆尼库瑟拉特.(*勒姆尼库瑟拉特Ramnicu_saratBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramnicu_sarat",
		TitleName: "勒姆尼库瑟拉特",
		TitleCode: "b_ramnicu_sarat",
	}
}
