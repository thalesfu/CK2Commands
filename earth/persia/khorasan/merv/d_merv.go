package merv

import (
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/merv/amol"
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/merv/konjikala"
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/merv/merv"
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/merv/sarakhs"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MervDuke interface {
    feud.Duke
    CAmol阿莫勒() 	amol.AmolCounty
    CKonjikala孔吉卡拉() 	konjikala.KonjikalaCounty
    CMerv木鹿() 	merv.MervCounty
    CSarakhs撒剌哈夕() 	sarakhs.SarakhsCounty
}

type 木鹿MervDuke struct {
	feud.BaseDuke
	阿莫勒Amol 	amol.AmolCounty
	孔吉卡拉Konjikala 	konjikala.KonjikalaCounty
	木鹿Merv 	merv.MervCounty
	撒剌哈夕Sarakhs 	sarakhs.SarakhsCounty
}

func (d *木鹿MervDuke) CAmol阿莫勒() amol.AmolCounty {
	return d.阿莫勒Amol
}
    
func (d *木鹿MervDuke) CKonjikala孔吉卡拉() konjikala.KonjikalaCounty {
	return d.孔吉卡拉Konjikala
}
    
func (d *木鹿MervDuke) CMerv木鹿() merv.MervCounty {
	return d.木鹿Merv
}
    
func (d *木鹿MervDuke) CSarakhs撒剌哈夕() sarakhs.SarakhsCounty {
	return d.撒剌哈夕Sarakhs
}
    
var DMerv木鹿 MervDuke = &木鹿MervDuke{}

func init() {
	f := DMerv木鹿.(*木鹿MervDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "merv",
		TitleName: "木鹿",
		TitleCode: "d_merv",
		Counties:  map[string]feud.County{},
	}

	f.阿莫勒Amol = amol.CAmol阿莫勒
	f.阿莫勒Amol.SetParent(f)
	
	f.孔吉卡拉Konjikala = konjikala.CKonjikala孔吉卡拉
	f.孔吉卡拉Konjikala.SetParent(f)
	
	f.木鹿Merv = merv.CMerv木鹿
	f.木鹿Merv.SetParent(f)
	
	f.撒剌哈夕Sarakhs = sarakhs.CSarakhs撒剌哈夕
	f.撒剌哈夕Sarakhs.SetParent(f)
	
}
