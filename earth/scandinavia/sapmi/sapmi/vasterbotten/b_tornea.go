package vasterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔尼奥TorneaBarony struct {
	feud.BaseBarony
}

var BTornea托尔尼奥 feud.Barony = &托尔尼奥TorneaBarony{}

func init() {
    f := BTornea托尔尼奥.(*托尔尼奥TorneaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tornea",
		TitleName: "托尔尼奥",
		TitleCode: "b_tornea",
	}
}
