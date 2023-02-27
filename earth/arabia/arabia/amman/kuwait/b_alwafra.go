package kuwait

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃夫拉AlwafraBarony struct {
	feud.BaseBarony
}

var BAlwafra沃夫拉 feud.Barony = &沃夫拉AlwafraBarony{}

func init() {
    f := BAlwafra沃夫拉.(*沃夫拉AlwafraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alwafra",
		TitleName: "沃夫拉",
		TitleCode: "b_alwafra",
	}
}
