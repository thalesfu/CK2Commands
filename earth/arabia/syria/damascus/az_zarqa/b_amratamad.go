package az_zarqa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿木拉特玛德AmratamadBarony struct {
	feud.BaseBarony
}

var BAmratamad阿木拉特玛德 feud.Barony = &阿木拉特玛德AmratamadBarony{}

func init() {
    f := BAmratamad阿木拉特玛德.(*阿木拉特玛德AmratamadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amratamad",
		TitleName: "阿木拉特玛德",
		TitleCode: "b_amratamad",
	}
}
