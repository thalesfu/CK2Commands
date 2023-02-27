package kyzikos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿德亚诺特莱AdrianutheraiBarony struct {
	feud.BaseBarony
}

var BAdrianutherai阿德亚诺特莱 feud.Barony = &阿德亚诺特莱AdrianutheraiBarony{}

func init() {
    f := BAdrianutherai阿德亚诺特莱.(*阿德亚诺特莱AdrianutheraiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adrianutherai",
		TitleName: "阿德亚诺特莱",
		TitleCode: "b_adrianutherai",
	}
}
