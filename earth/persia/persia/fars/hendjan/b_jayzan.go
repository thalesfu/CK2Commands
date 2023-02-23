package hendjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎耶赞JayzanBarony struct {
	feud.BaseBarony
}

var BJayzan扎耶赞 feud.Barony = &扎耶赞JayzanBarony{}

func init() {
	f := BJayzan扎耶赞.(*扎耶赞JayzanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jayzan",
		TitleName: "扎耶赞",
		TitleCode: "b_jayzan",
	}
}
