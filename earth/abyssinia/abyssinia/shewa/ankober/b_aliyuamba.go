package ankober

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿利尤安巴AliyuambaBarony struct {
	feud.BaseBarony
}

var BAliyuamba阿利尤安巴 feud.Barony = &阿利尤安巴AliyuambaBarony{}

func init() {
    f := BAliyuamba阿利尤安巴.(*阿利尤安巴AliyuambaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aliyuamba",
		TitleName: "阿利尤安巴",
		TitleCode: "b_aliyuamba",
	}
}
