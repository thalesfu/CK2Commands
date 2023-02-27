package makuria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MakuriaCounty interface {
    feud.County
    BAffat阿夫法特() 	feud.Barony
    BDongola栋古拉() 	feud.Barony
    BGabriya伽比亚() 	feud.Barony
    BKawa凯沃() 	feud.Barony
    BKerat凯拉特() 	feud.Barony
    BKerma科尔玛() 	feud.Barony
    BQubban古班() 	feud.Barony
}

type 马库里亚MakuriaCounty struct {
	feud.BaseCounty
	阿夫法特Affat 	feud.Barony
	栋古拉Dongola 	feud.Barony
	伽比亚Gabriya 	feud.Barony
	凯沃Kawa 	feud.Barony
	凯拉特Kerat 	feud.Barony
	科尔玛Kerma 	feud.Barony
	古班Qubban 	feud.Barony
}

func (c *马库里亚MakuriaCounty) BAffat阿夫法特() feud.Barony {
	return c.阿夫法特Affat
}
    
func (c *马库里亚MakuriaCounty) BDongola栋古拉() feud.Barony {
	return c.栋古拉Dongola
}
    
func (c *马库里亚MakuriaCounty) BGabriya伽比亚() feud.Barony {
	return c.伽比亚Gabriya
}
    
func (c *马库里亚MakuriaCounty) BKawa凯沃() feud.Barony {
	return c.凯沃Kawa
}
    
func (c *马库里亚MakuriaCounty) BKerat凯拉特() feud.Barony {
	return c.凯拉特Kerat
}
    
func (c *马库里亚MakuriaCounty) BKerma科尔玛() feud.Barony {
	return c.科尔玛Kerma
}
    
func (c *马库里亚MakuriaCounty) BQubban古班() feud.Barony {
	return c.古班Qubban
}
    
var CMakuria马库里亚 MakuriaCounty = &马库里亚MakuriaCounty{}

func init() {
	f := CMakuria马库里亚.(*马库里亚MakuriaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "793",
		Title:     "makuria",
		TitleName: "马库里亚",
		TitleCode: "c_makuria",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿夫法特Affat = BAffat阿夫法特
	f.阿夫法特Affat.SetParent(f)
	
	f.栋古拉Dongola = BDongola栋古拉
	f.栋古拉Dongola.SetParent(f)
	
	f.伽比亚Gabriya = BGabriya伽比亚
	f.伽比亚Gabriya.SetParent(f)
	
	f.凯沃Kawa = BKawa凯沃
	f.凯沃Kawa.SetParent(f)
	
	f.凯拉特Kerat = BKerat凯拉特
	f.凯拉特Kerat.SetParent(f)
	
	f.科尔玛Kerma = BKerma科尔玛
	f.科尔玛Kerma.SetParent(f)
	
	f.古班Qubban = BQubban古班
	f.古班Qubban.SetParent(f)
	
}
