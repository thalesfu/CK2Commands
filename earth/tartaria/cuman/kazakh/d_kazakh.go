package kazakh

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/kazakh/balkhash"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/kazakh/betpaqdala"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/kazakh/kazakh"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/kazakh/qarazhyrya"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/kazakh/sarysu"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KazakhDuke interface {
    feud.Duke
    CBalkhash巴尔喀什() 	balkhash.BalkhashCounty
    CBetpaqdala别特帕克达拉() 	betpaqdala.BetpaqdalaCounty
    CKazakh哈萨克() 	kazakh.KazakhCounty
    CQarazhyrya喀拉哲拉() 	qarazhyrya.QarazhyryaCounty
    CSarysu萨雷苏() 	sarysu.SarysuCounty
}

type 哈萨克KazakhDuke struct {
	feud.BaseDuke
	巴尔喀什Balkhash 	balkhash.BalkhashCounty
	别特帕克达拉Betpaqdala 	betpaqdala.BetpaqdalaCounty
	哈萨克Kazakh 	kazakh.KazakhCounty
	喀拉哲拉Qarazhyrya 	qarazhyrya.QarazhyryaCounty
	萨雷苏Sarysu 	sarysu.SarysuCounty
}

func (d *哈萨克KazakhDuke) CBalkhash巴尔喀什() balkhash.BalkhashCounty {
	return d.巴尔喀什Balkhash
}
    
func (d *哈萨克KazakhDuke) CBetpaqdala别特帕克达拉() betpaqdala.BetpaqdalaCounty {
	return d.别特帕克达拉Betpaqdala
}
    
func (d *哈萨克KazakhDuke) CKazakh哈萨克() kazakh.KazakhCounty {
	return d.哈萨克Kazakh
}
    
func (d *哈萨克KazakhDuke) CQarazhyrya喀拉哲拉() qarazhyrya.QarazhyryaCounty {
	return d.喀拉哲拉Qarazhyrya
}
    
func (d *哈萨克KazakhDuke) CSarysu萨雷苏() sarysu.SarysuCounty {
	return d.萨雷苏Sarysu
}
    
var DKazakh哈萨克 KazakhDuke = &哈萨克KazakhDuke{}

func init() {
	f := DKazakh哈萨克.(*哈萨克KazakhDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kazakh",
		TitleName: "哈萨克",
		TitleCode: "d_kazakh",
		Counties:  map[string]feud.County{},
	}

	f.巴尔喀什Balkhash = balkhash.CBalkhash巴尔喀什
	f.巴尔喀什Balkhash.SetParent(f)
	
	f.别特帕克达拉Betpaqdala = betpaqdala.CBetpaqdala别特帕克达拉
	f.别特帕克达拉Betpaqdala.SetParent(f)
	
	f.哈萨克Kazakh = kazakh.CKazakh哈萨克
	f.哈萨克Kazakh.SetParent(f)
	
	f.喀拉哲拉Qarazhyrya = qarazhyrya.CQarazhyrya喀拉哲拉
	f.喀拉哲拉Qarazhyrya.SetParent(f)
	
	f.萨雷苏Sarysu = sarysu.CSarysu萨雷苏
	f.萨雷苏Sarysu.SetParent(f)
	
}
