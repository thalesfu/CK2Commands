package hajr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉赞JizanBarony struct {
	feud.BaseBarony
}

var BJizan吉赞 feud.Barony = &吉赞JizanBarony{}

func init() {
	f := BJizan吉赞.(*吉赞JizanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jizan",
		TitleName: "吉赞",
		TitleCode: "b_jizan",
	}
}
