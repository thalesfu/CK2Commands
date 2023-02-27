package ungvar

import (
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/ungvar/abauj"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/ungvar/bereg"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/ungvar/marmaros"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/ungvar/saris"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UngvarDuke interface {
    feud.Duke
    CAbauj奥鲍乌伊() 	abauj.AbaujCounty
    CBereg拜赖格() 	bereg.BeregCounty
    CMarmaros马尔毛罗什() 	marmaros.MarmarosCounty
    CSaris斯皮什() 	saris.SarisCounty
}

type 翁格瓦尔UngvarDuke struct {
	feud.BaseDuke
	奥鲍乌伊Abauj 	abauj.AbaujCounty
	拜赖格Bereg 	bereg.BeregCounty
	马尔毛罗什Marmaros 	marmaros.MarmarosCounty
	斯皮什Saris 	saris.SarisCounty
}

func (d *翁格瓦尔UngvarDuke) CAbauj奥鲍乌伊() abauj.AbaujCounty {
	return d.奥鲍乌伊Abauj
}
    
func (d *翁格瓦尔UngvarDuke) CBereg拜赖格() bereg.BeregCounty {
	return d.拜赖格Bereg
}
    
func (d *翁格瓦尔UngvarDuke) CMarmaros马尔毛罗什() marmaros.MarmarosCounty {
	return d.马尔毛罗什Marmaros
}
    
func (d *翁格瓦尔UngvarDuke) CSaris斯皮什() saris.SarisCounty {
	return d.斯皮什Saris
}
    
var DUngvar翁格瓦尔 UngvarDuke = &翁格瓦尔UngvarDuke{}

func init() {
	f := DUngvar翁格瓦尔.(*翁格瓦尔UngvarDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ungvar",
		TitleName: "翁格瓦尔",
		TitleCode: "d_ungvar",
		Counties:  map[string]feud.County{},
	}

	f.奥鲍乌伊Abauj = abauj.CAbauj奥鲍乌伊
	f.奥鲍乌伊Abauj.SetParent(f)
	
	f.拜赖格Bereg = bereg.CBereg拜赖格
	f.拜赖格Bereg.SetParent(f)
	
	f.马尔毛罗什Marmaros = marmaros.CMarmaros马尔毛罗什
	f.马尔毛罗什Marmaros.SetParent(f)
	
	f.斯皮什Saris = saris.CSaris斯皮什
	f.斯皮什Saris.SetParent(f)
	
}
