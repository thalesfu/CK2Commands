package mzab

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/mzab/arigh"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/mzab/mzab"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/mzab/tuggurt"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MzabDuke interface {
	feud.Duke
	CArigh阿里格() arigh.ArighCounty
	CMzab盖尔达耶() mzab.MzabCounty
	CTuggurt图古尔特() tuggurt.TuggurtCounty
}

type 姆扎卜MzabDuke struct {
	feud.BaseDuke
	阿里格Arigh    arigh.ArighCounty
	盖尔达耶Mzab    mzab.MzabCounty
	图古尔特Tuggurt tuggurt.TuggurtCounty
}

func (d *姆扎卜MzabDuke) CArigh阿里格() arigh.ArighCounty {
	return d.阿里格Arigh
}

func (d *姆扎卜MzabDuke) CMzab盖尔达耶() mzab.MzabCounty {
	return d.盖尔达耶Mzab
}

func (d *姆扎卜MzabDuke) CTuggurt图古尔特() tuggurt.TuggurtCounty {
	return d.图古尔特Tuggurt
}

var DMzab姆扎卜 MzabDuke = &姆扎卜MzabDuke{}

func init() {
	f := DMzab姆扎卜.(*姆扎卜MzabDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "mzab",
		TitleName: "姆扎卜",
		TitleCode: "d_mzab",
		Counties:  map[string]feud.County{},
	}

	f.阿里格Arigh = arigh.CArigh阿里格
	f.阿里格Arigh.SetParent(f)

	f.盖尔达耶Mzab = mzab.CMzab盖尔达耶
	f.盖尔达耶Mzab.SetParent(f)

	f.图古尔特Tuggurt = tuggurt.CTuggurt图古尔特
	f.图古尔特Tuggurt.SetParent(f)

}
