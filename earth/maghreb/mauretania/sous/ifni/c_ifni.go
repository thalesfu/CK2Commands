package ifni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IfniCounty interface {
    feud.County
    BAkhfennir阿克费尼() 	feud.Barony
    BAmot奥穆特() 	feud.Barony
    BChbika舍比凯() 	feud.Barony
    BIfni伊夫尼() 	feud.Barony
    BTantan坦坦() 	feud.Barony
    BTartaya塔尔法亚() 	feud.Barony
    BTharasset唐图勒() 	feud.Barony
    BTidergit提德基特() 	feud.Barony
    BTiglit梯格利特() 	feud.Barony
}

type 伊夫尼IfniCounty struct {
	feud.BaseCounty
	阿克费尼Akhfennir 	feud.Barony
	奥穆特Amot 	feud.Barony
	舍比凯Chbika 	feud.Barony
	伊夫尼Ifni 	feud.Barony
	坦坦Tantan 	feud.Barony
	塔尔法亚Tartaya 	feud.Barony
	唐图勒Tharasset 	feud.Barony
	提德基特Tidergit 	feud.Barony
	梯格利特Tiglit 	feud.Barony
}

func (c *伊夫尼IfniCounty) BAkhfennir阿克费尼() feud.Barony {
	return c.阿克费尼Akhfennir
}
    
func (c *伊夫尼IfniCounty) BAmot奥穆特() feud.Barony {
	return c.奥穆特Amot
}
    
func (c *伊夫尼IfniCounty) BChbika舍比凯() feud.Barony {
	return c.舍比凯Chbika
}
    
func (c *伊夫尼IfniCounty) BIfni伊夫尼() feud.Barony {
	return c.伊夫尼Ifni
}
    
func (c *伊夫尼IfniCounty) BTantan坦坦() feud.Barony {
	return c.坦坦Tantan
}
    
func (c *伊夫尼IfniCounty) BTartaya塔尔法亚() feud.Barony {
	return c.塔尔法亚Tartaya
}
    
func (c *伊夫尼IfniCounty) BTharasset唐图勒() feud.Barony {
	return c.唐图勒Tharasset
}
    
func (c *伊夫尼IfniCounty) BTidergit提德基特() feud.Barony {
	return c.提德基特Tidergit
}
    
func (c *伊夫尼IfniCounty) BTiglit梯格利特() feud.Barony {
	return c.梯格利特Tiglit
}
    
var CIfni伊夫尼 IfniCounty = &伊夫尼IfniCounty{}

func init() {
	f := CIfni伊夫尼.(*伊夫尼IfniCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "846",
		Title:     "ifni",
		TitleName: "伊夫尼",
		TitleCode: "c_ifni",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克费尼Akhfennir = BAkhfennir阿克费尼
	f.阿克费尼Akhfennir.SetParent(f)
	
	f.奥穆特Amot = BAmot奥穆特
	f.奥穆特Amot.SetParent(f)
	
	f.舍比凯Chbika = BChbika舍比凯
	f.舍比凯Chbika.SetParent(f)
	
	f.伊夫尼Ifni = BIfni伊夫尼
	f.伊夫尼Ifni.SetParent(f)
	
	f.坦坦Tantan = BTantan坦坦
	f.坦坦Tantan.SetParent(f)
	
	f.塔尔法亚Tartaya = BTartaya塔尔法亚
	f.塔尔法亚Tartaya.SetParent(f)
	
	f.唐图勒Tharasset = BTharasset唐图勒
	f.唐图勒Tharasset.SetParent(f)
	
	f.提德基特Tidergit = BTidergit提德基特
	f.提德基特Tidergit.SetParent(f)
	
	f.梯格利特Tiglit = BTiglit梯格利特
	f.梯格利特Tiglit.SetParent(f)
	
}
