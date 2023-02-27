package qalqut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QalqutCounty interface {
    feud.County
    BCalicut卡利卡特() 	feud.Barony
    BChaliyam伽利耶() 	feud.Barony
    BKannur甘努尔() 	feud.Barony
    BNediyiruppu内迪耶卢补() 	feud.Barony
    BParappanangadi波罗波那迦迪() 	feud.Barony
    BUrakam尤罗甘() 	feud.Barony
    BVatakara伐多伽罗() 	feud.Barony
}

type 古里QalqutCounty struct {
	feud.BaseCounty
	卡利卡特Calicut 	feud.Barony
	伽利耶Chaliyam 	feud.Barony
	甘努尔Kannur 	feud.Barony
	内迪耶卢补Nediyiruppu 	feud.Barony
	波罗波那迦迪Parappanangadi 	feud.Barony
	尤罗甘Urakam 	feud.Barony
	伐多伽罗Vatakara 	feud.Barony
}

func (c *古里QalqutCounty) BCalicut卡利卡特() feud.Barony {
	return c.卡利卡特Calicut
}
    
func (c *古里QalqutCounty) BChaliyam伽利耶() feud.Barony {
	return c.伽利耶Chaliyam
}
    
func (c *古里QalqutCounty) BKannur甘努尔() feud.Barony {
	return c.甘努尔Kannur
}
    
func (c *古里QalqutCounty) BNediyiruppu内迪耶卢补() feud.Barony {
	return c.内迪耶卢补Nediyiruppu
}
    
func (c *古里QalqutCounty) BParappanangadi波罗波那迦迪() feud.Barony {
	return c.波罗波那迦迪Parappanangadi
}
    
func (c *古里QalqutCounty) BUrakam尤罗甘() feud.Barony {
	return c.尤罗甘Urakam
}
    
func (c *古里QalqutCounty) BVatakara伐多伽罗() feud.Barony {
	return c.伐多伽罗Vatakara
}
    
var CQalqut古里 QalqutCounty = &古里QalqutCounty{}

func init() {
	f := CQalqut古里.(*古里QalqutCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1117",
		Title:     "qalqut",
		TitleName: "古里",
		TitleCode: "c_qalqut",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡利卡特Calicut = BCalicut卡利卡特
	f.卡利卡特Calicut.SetParent(f)
	
	f.伽利耶Chaliyam = BChaliyam伽利耶
	f.伽利耶Chaliyam.SetParent(f)
	
	f.甘努尔Kannur = BKannur甘努尔
	f.甘努尔Kannur.SetParent(f)
	
	f.内迪耶卢补Nediyiruppu = BNediyiruppu内迪耶卢补
	f.内迪耶卢补Nediyiruppu.SetParent(f)
	
	f.波罗波那迦迪Parappanangadi = BParappanangadi波罗波那迦迪
	f.波罗波那迦迪Parappanangadi.SetParent(f)
	
	f.尤罗甘Urakam = BUrakam尤罗甘
	f.尤罗甘Urakam.SetParent(f)
	
	f.伐多伽罗Vatakara = BVatakara伐多伽罗
	f.伐多伽罗Vatakara.SetParent(f)
	
}
