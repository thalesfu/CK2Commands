package baluchistan

import (
	"github.com/thalesfu/CK2Commands/earth/persia/baluchistan/baluchistan/armail"
	"github.com/thalesfu/CK2Commands/earth/persia/baluchistan/baluchistan/chagai"
	"github.com/thalesfu/CK2Commands/earth/persia/baluchistan/baluchistan/makran"
	"github.com/thalesfu/CK2Commands/earth/persia/baluchistan/baluchistan/saravan"
	"github.com/thalesfu/CK2Commands/earth/persia/baluchistan/baluchistan/tis"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BaluchistanDuke interface {
    feud.Duke
    CArmail阿尔梅尔() 	armail.ArmailCounty
    CChagai查盖() 	chagai.ChagaiCounty
    CMakran弥兰() 	makran.MakranCounty
    CSaravan萨拉万() 	saravan.SaravanCounty
    CTis蒂斯() 	tis.TisCounty
}

type 弥兰BaluchistanDuke struct {
	feud.BaseDuke
	阿尔梅尔Armail 	armail.ArmailCounty
	查盖Chagai 	chagai.ChagaiCounty
	弥兰Makran 	makran.MakranCounty
	萨拉万Saravan 	saravan.SaravanCounty
	蒂斯Tis 	tis.TisCounty
}

func (d *弥兰BaluchistanDuke) CArmail阿尔梅尔() armail.ArmailCounty {
	return d.阿尔梅尔Armail
}
    
func (d *弥兰BaluchistanDuke) CChagai查盖() chagai.ChagaiCounty {
	return d.查盖Chagai
}
    
func (d *弥兰BaluchistanDuke) CMakran弥兰() makran.MakranCounty {
	return d.弥兰Makran
}
    
func (d *弥兰BaluchistanDuke) CSaravan萨拉万() saravan.SaravanCounty {
	return d.萨拉万Saravan
}
    
func (d *弥兰BaluchistanDuke) CTis蒂斯() tis.TisCounty {
	return d.蒂斯Tis
}
    
var DBaluchistan弥兰 BaluchistanDuke = &弥兰BaluchistanDuke{}

func init() {
	f := DBaluchistan弥兰.(*弥兰BaluchistanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "baluchistan",
		TitleName: "弥兰",
		TitleCode: "d_baluchistan",
		Counties:  map[string]feud.County{},
	}

	f.阿尔梅尔Armail = armail.CArmail阿尔梅尔
	f.阿尔梅尔Armail.SetParent(f)
	
	f.查盖Chagai = chagai.CChagai查盖
	f.查盖Chagai.SetParent(f)
	
	f.弥兰Makran = makran.CMakran弥兰
	f.弥兰Makran.SetParent(f)
	
	f.萨拉万Saravan = saravan.CSaravan萨拉万
	f.萨拉万Saravan.SetParent(f)
	
	f.蒂斯Tis = tis.CTis蒂斯
	f.蒂斯Tis.SetParent(f)
	
}
