package ravenna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘贝托拉GambettolaBarony struct {
	feud.BaseBarony
}

var BGambettola甘贝托拉 feud.Barony = &甘贝托拉GambettolaBarony{}

func init() {
    f := BGambettola甘贝托拉.(*甘贝托拉GambettolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gambettola",
		TitleName: "甘贝托拉",
		TitleCode: "b_gambettola",
	}
}
