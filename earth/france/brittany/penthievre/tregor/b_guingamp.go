package tregor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘冈GuingampBarony struct {
	feud.BaseBarony
}

var BGuingamp甘冈 feud.Barony = &甘冈GuingampBarony{}

func init() {
    f := BGuingamp甘冈.(*甘冈GuingampBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guingamp",
		TitleName: "甘冈",
		TitleCode: "b_guingamp",
	}
}
