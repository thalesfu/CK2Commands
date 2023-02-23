package kalpi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽勒菩乔GalleborgaoBarony struct {
	feud.BaseBarony
}

var BGalleborgao伽勒菩乔 feud.Barony = &伽勒菩乔GalleborgaoBarony{}

func init() {
	f := BGalleborgao伽勒菩乔.(*伽勒菩乔GalleborgaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galleborgao",
		TitleName: "伽勒菩乔",
		TitleCode: "b_galleborgao",
	}
}
