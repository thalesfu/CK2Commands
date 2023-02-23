package karashar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 危须WeixuBarony struct {
	feud.BaseBarony
}

var BWeixu危须 feud.Barony = &危须WeixuBarony{}

func init() {
	f := BWeixu危须.(*危须WeixuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "weixu",
		TitleName: "危须",
		TitleCode: "b_weixu",
	}
}
