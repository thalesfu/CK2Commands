package tell_bashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Tell_bashirCounty interface {
    feud.County
    BBirtha别尔哈() 	feud.Barony
    BBuzaa布扎阿() 	feud.Barony
    BJarabulus杰拉布卢斯() 	feud.Barony
    BMakedonopolis马其顿诺波利斯() 	feud.Barony
    BManbij曼比季() 	feud.Barony
    BNizip尼济普() 	feud.Barony
    BTurbessel图比瑟尔() 	feud.Barony
    BZeugma泽乌格马() 	feud.Barony
}

type 巴希尔丘Tell_bashirCounty struct {
	feud.BaseCounty
	别尔哈Birtha 	feud.Barony
	布扎阿Buzaa 	feud.Barony
	杰拉布卢斯Jarabulus 	feud.Barony
	马其顿诺波利斯Makedonopolis 	feud.Barony
	曼比季Manbij 	feud.Barony
	尼济普Nizip 	feud.Barony
	图比瑟尔Turbessel 	feud.Barony
	泽乌格马Zeugma 	feud.Barony
}

func (c *巴希尔丘Tell_bashirCounty) BBirtha别尔哈() feud.Barony {
	return c.别尔哈Birtha
}
    
func (c *巴希尔丘Tell_bashirCounty) BBuzaa布扎阿() feud.Barony {
	return c.布扎阿Buzaa
}
    
func (c *巴希尔丘Tell_bashirCounty) BJarabulus杰拉布卢斯() feud.Barony {
	return c.杰拉布卢斯Jarabulus
}
    
func (c *巴希尔丘Tell_bashirCounty) BMakedonopolis马其顿诺波利斯() feud.Barony {
	return c.马其顿诺波利斯Makedonopolis
}
    
func (c *巴希尔丘Tell_bashirCounty) BManbij曼比季() feud.Barony {
	return c.曼比季Manbij
}
    
func (c *巴希尔丘Tell_bashirCounty) BNizip尼济普() feud.Barony {
	return c.尼济普Nizip
}
    
func (c *巴希尔丘Tell_bashirCounty) BTurbessel图比瑟尔() feud.Barony {
	return c.图比瑟尔Turbessel
}
    
func (c *巴希尔丘Tell_bashirCounty) BZeugma泽乌格马() feud.Barony {
	return c.泽乌格马Zeugma
}
    
var CTell_bashir巴希尔丘 Tell_bashirCounty = &巴希尔丘Tell_bashirCounty{}

func init() {
	f := CTell_bashir巴希尔丘.(*巴希尔丘Tell_bashirCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "708",
		Title:     "tell_bashir",
		TitleName: "巴希尔丘",
		TitleCode: "c_tell_bashir",
		Baronies:  map[string]feud.Barony{},
	}

	f.别尔哈Birtha = BBirtha别尔哈
	f.别尔哈Birtha.SetParent(f)
	
	f.布扎阿Buzaa = BBuzaa布扎阿
	f.布扎阿Buzaa.SetParent(f)
	
	f.杰拉布卢斯Jarabulus = BJarabulus杰拉布卢斯
	f.杰拉布卢斯Jarabulus.SetParent(f)
	
	f.马其顿诺波利斯Makedonopolis = BMakedonopolis马其顿诺波利斯
	f.马其顿诺波利斯Makedonopolis.SetParent(f)
	
	f.曼比季Manbij = BManbij曼比季
	f.曼比季Manbij.SetParent(f)
	
	f.尼济普Nizip = BNizip尼济普
	f.尼济普Nizip.SetParent(f)
	
	f.图比瑟尔Turbessel = BTurbessel图比瑟尔
	f.图比瑟尔Turbessel.SetParent(f)
	
	f.泽乌格马Zeugma = BZeugma泽乌格马
	f.泽乌格马Zeugma.SetParent(f)
	
}
