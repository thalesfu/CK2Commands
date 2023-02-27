package corfu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿金加诺鲁姆AcinganorumBarony struct {
	feud.BaseBarony
}

var BAcinganorum阿金加诺鲁姆 feud.Barony = &阿金加诺鲁姆AcinganorumBarony{}

func init() {
    f := BAcinganorum阿金加诺鲁姆.(*阿金加诺鲁姆AcinganorumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "acinganorum",
		TitleName: "阿金加诺鲁姆",
		TitleCode: "b_acinganorum",
	}
}
