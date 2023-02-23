package kashgar

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/kashgar/aksu"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/kashgar/artux"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/kashgar/kashgar"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/kashgar/uchturpan"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/kashgar/yopurga"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KashgarDuke interface {
	feud.Duke
	CAksu阿克苏() aksu.AksuCounty
	CArtux阿图什() artux.ArtuxCounty
	CKashgar可失合儿() kashgar.KashgarCounty
	CUchturpan乌什吐鲁番() uchturpan.UchturpanCounty
	CYopurga岳普湖() yopurga.YopurgaCounty
}

type 可失合儿KashgarDuke struct {
	feud.BaseDuke
	阿克苏Aksu        aksu.AksuCounty
	阿图什Artux       artux.ArtuxCounty
	可失合儿Kashgar    kashgar.KashgarCounty
	乌什吐鲁番Uchturpan uchturpan.UchturpanCounty
	岳普湖Yopurga     yopurga.YopurgaCounty
}

func (d *可失合儿KashgarDuke) CAksu阿克苏() aksu.AksuCounty {
	return d.阿克苏Aksu
}

func (d *可失合儿KashgarDuke) CArtux阿图什() artux.ArtuxCounty {
	return d.阿图什Artux
}

func (d *可失合儿KashgarDuke) CKashgar可失合儿() kashgar.KashgarCounty {
	return d.可失合儿Kashgar
}

func (d *可失合儿KashgarDuke) CUchturpan乌什吐鲁番() uchturpan.UchturpanCounty {
	return d.乌什吐鲁番Uchturpan
}

func (d *可失合儿KashgarDuke) CYopurga岳普湖() yopurga.YopurgaCounty {
	return d.岳普湖Yopurga
}

var DKashgar可失合儿 KashgarDuke = &可失合儿KashgarDuke{}

func init() {
	f := DKashgar可失合儿.(*可失合儿KashgarDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kashgar",
		TitleName: "可失合儿",
		TitleCode: "d_kashgar",
		Counties:  map[string]feud.County{},
	}

	f.阿克苏Aksu = aksu.CAksu阿克苏
	f.阿克苏Aksu.SetParent(f)

	f.阿图什Artux = artux.CArtux阿图什
	f.阿图什Artux.SetParent(f)

	f.可失合儿Kashgar = kashgar.CKashgar可失合儿
	f.可失合儿Kashgar.SetParent(f)

	f.乌什吐鲁番Uchturpan = uchturpan.CUchturpan乌什吐鲁番
	f.乌什吐鲁番Uchturpan.SetParent(f)

	f.岳普湖Yopurga = yopurga.CYopurga岳普湖
	f.岳普湖Yopurga.SetParent(f)

}
