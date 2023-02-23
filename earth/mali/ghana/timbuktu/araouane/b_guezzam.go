package araouane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖赞GuezzamBarony struct {
	feud.BaseBarony
}

var BGuezzam盖赞 feud.Barony = &盖赞GuezzamBarony{}

func init() {
	f := BGuezzam盖赞.(*盖赞GuezzamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guezzam",
		TitleName: "盖赞",
		TitleCode: "b_guezzam",
	}
}
