package zahedan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZahedanCounty interface {
    feud.County
    BBuk布克() 	feud.Barony
    BGolchah格拉察() 	feud.Barony
    BHamun达姆恩() 	feud.Barony
    BJahangir贾汉吉尔() 	feud.Barony
    BKacharud卡察鲁德() 	feud.Barony
    BKhajeh克哈吉() 	feud.Barony
    BKuhtaftan毾㲪山() 	feud.Barony
    BZahedan扎黑丹() 	feud.Barony
}

type 扎黑丹ZahedanCounty struct {
	feud.BaseCounty
	布克Buk 	feud.Barony
	格拉察Golchah 	feud.Barony
	达姆恩Hamun 	feud.Barony
	贾汉吉尔Jahangir 	feud.Barony
	卡察鲁德Kacharud 	feud.Barony
	克哈吉Khajeh 	feud.Barony
	毾㲪山Kuhtaftan 	feud.Barony
	扎黑丹Zahedan 	feud.Barony
}

func (c *扎黑丹ZahedanCounty) BBuk布克() feud.Barony {
	return c.布克Buk
}
    
func (c *扎黑丹ZahedanCounty) BGolchah格拉察() feud.Barony {
	return c.格拉察Golchah
}
    
func (c *扎黑丹ZahedanCounty) BHamun达姆恩() feud.Barony {
	return c.达姆恩Hamun
}
    
func (c *扎黑丹ZahedanCounty) BJahangir贾汉吉尔() feud.Barony {
	return c.贾汉吉尔Jahangir
}
    
func (c *扎黑丹ZahedanCounty) BKacharud卡察鲁德() feud.Barony {
	return c.卡察鲁德Kacharud
}
    
func (c *扎黑丹ZahedanCounty) BKhajeh克哈吉() feud.Barony {
	return c.克哈吉Khajeh
}
    
func (c *扎黑丹ZahedanCounty) BKuhtaftan毾㲪山() feud.Barony {
	return c.毾㲪山Kuhtaftan
}
    
func (c *扎黑丹ZahedanCounty) BZahedan扎黑丹() feud.Barony {
	return c.扎黑丹Zahedan
}
    
var CZahedan扎黑丹 ZahedanCounty = &扎黑丹ZahedanCounty{}

func init() {
	f := CZahedan扎黑丹.(*扎黑丹ZahedanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "852",
		Title:     "zahedan",
		TitleName: "扎黑丹",
		TitleCode: "c_zahedan",
		Baronies:  map[string]feud.Barony{},
	}

	f.布克Buk = BBuk布克
	f.布克Buk.SetParent(f)
	
	f.格拉察Golchah = BGolchah格拉察
	f.格拉察Golchah.SetParent(f)
	
	f.达姆恩Hamun = BHamun达姆恩
	f.达姆恩Hamun.SetParent(f)
	
	f.贾汉吉尔Jahangir = BJahangir贾汉吉尔
	f.贾汉吉尔Jahangir.SetParent(f)
	
	f.卡察鲁德Kacharud = BKacharud卡察鲁德
	f.卡察鲁德Kacharud.SetParent(f)
	
	f.克哈吉Khajeh = BKhajeh克哈吉
	f.克哈吉Khajeh.SetParent(f)
	
	f.毾㲪山Kuhtaftan = BKuhtaftan毾㲪山
	f.毾㲪山Kuhtaftan.SetParent(f)
	
	f.扎黑丹Zahedan = BZahedan扎黑丹
	f.扎黑丹Zahedan.SetParent(f)
	
}
