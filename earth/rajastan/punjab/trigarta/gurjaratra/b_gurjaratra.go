package gurjaratra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿折罗多罗GurjaratraBarony struct {
	feud.BaseBarony
}

var BGurjaratra瞿折罗多罗 feud.Barony = &瞿折罗多罗GurjaratraBarony{}

func init() {
    f := BGurjaratra瞿折罗多罗.(*瞿折罗多罗GurjaratraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gurjaratra",
		TitleName: "瞿折罗多罗",
		TitleCode: "b_gurjaratra",
	}
}
