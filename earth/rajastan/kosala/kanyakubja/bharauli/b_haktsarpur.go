package bharauli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃克娑补HaktsarpurBarony struct {
	feud.BaseBarony
}

var BHaktsarpur诃克娑补 feud.Barony = &诃克娑补HaktsarpurBarony{}

func init() {
    f := BHaktsarpur诃克娑补.(*诃克娑补HaktsarpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haktsarpur",
		TitleName: "诃克娑补",
		TitleCode: "b_haktsarpur",
	}
}
