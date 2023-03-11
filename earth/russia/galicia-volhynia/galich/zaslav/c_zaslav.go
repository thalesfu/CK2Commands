package zaslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZaslavCounty interface {
    feud.County
    BBuzhesk布热斯克() 	feud.Barony
    BChepetivka舍佩蒂夫卡() 	feud.Barony
    BKremyanets克列梅涅茨() 	feud.Barony
    BOstroh奥斯特罗赫() 	feud.Barony
    BPlisnesk普列斯涅斯克() 	feud.Barony
    BSlavuta斯拉武塔() 	feud.Barony
    BZaslav扎斯拉夫() 	feud.Barony
}

type 扎斯拉夫ZaslavCounty struct {
	feud.BaseCounty
	布热斯克Buzhesk 	feud.Barony
	舍佩蒂夫卡Chepetivka 	feud.Barony
	克列梅涅茨Kremyanets 	feud.Barony
	奥斯特罗赫Ostroh 	feud.Barony
	普列斯涅斯克Plisnesk 	feud.Barony
	斯拉武塔Slavuta 	feud.Barony
	扎斯拉夫Zaslav 	feud.Barony
}

func (c *扎斯拉夫ZaslavCounty) BBuzhesk布热斯克() feud.Barony {
	return c.布热斯克Buzhesk
}
    
func (c *扎斯拉夫ZaslavCounty) BChepetivka舍佩蒂夫卡() feud.Barony {
	return c.舍佩蒂夫卡Chepetivka
}
    
func (c *扎斯拉夫ZaslavCounty) BKremyanets克列梅涅茨() feud.Barony {
	return c.克列梅涅茨Kremyanets
}
    
func (c *扎斯拉夫ZaslavCounty) BOstroh奥斯特罗赫() feud.Barony {
	return c.奥斯特罗赫Ostroh
}
    
func (c *扎斯拉夫ZaslavCounty) BPlisnesk普列斯涅斯克() feud.Barony {
	return c.普列斯涅斯克Plisnesk
}
    
func (c *扎斯拉夫ZaslavCounty) BSlavuta斯拉武塔() feud.Barony {
	return c.斯拉武塔Slavuta
}
    
func (c *扎斯拉夫ZaslavCounty) BZaslav扎斯拉夫() feud.Barony {
	return c.扎斯拉夫Zaslav
}
    
var CZaslav扎斯拉夫 ZaslavCounty = &扎斯拉夫ZaslavCounty{}

func init() {
	f := CZaslav扎斯拉夫.(*扎斯拉夫ZaslavCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1631",
		Title:     "zaslav",
		TitleName: "扎斯拉夫",
		TitleCode: "c_zaslav",
		Baronies:  map[string]feud.Barony{},
	}

	f.布热斯克Buzhesk = BBuzhesk布热斯克
	f.布热斯克Buzhesk.SetParent(f)
	
	f.舍佩蒂夫卡Chepetivka = BChepetivka舍佩蒂夫卡
	f.舍佩蒂夫卡Chepetivka.SetParent(f)
	
	f.克列梅涅茨Kremyanets = BKremyanets克列梅涅茨
	f.克列梅涅茨Kremyanets.SetParent(f)
	
	f.奥斯特罗赫Ostroh = BOstroh奥斯特罗赫
	f.奥斯特罗赫Ostroh.SetParent(f)
	
	f.普列斯涅斯克Plisnesk = BPlisnesk普列斯涅斯克
	f.普列斯涅斯克Plisnesk.SetParent(f)
	
	f.斯拉武塔Slavuta = BSlavuta斯拉武塔
	f.斯拉武塔Slavuta.SetParent(f)
	
	f.扎斯拉夫Zaslav = BZaslav扎斯拉夫
	f.扎斯拉夫Zaslav.SetParent(f)
	
}
