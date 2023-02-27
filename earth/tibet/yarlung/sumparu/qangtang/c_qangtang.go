package qangtang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QangtangCounty interface {
    feud.County
    BBairab拜惹布() 	feud.Barony
    BBurog布若() 	feud.Barony
    BLaxong拉雄() 	feud.Barony
    BMargai玛尔盖() 	feud.Barony
    BSatso萨错() 	feud.Barony
    BYurba涌波() 	feud.Barony
    BZhoyon雪源() 	feud.Barony
}

type 羌塘QangtangCounty struct {
	feud.BaseCounty
	拜惹布Bairab 	feud.Barony
	布若Burog 	feud.Barony
	拉雄Laxong 	feud.Barony
	玛尔盖Margai 	feud.Barony
	萨错Satso 	feud.Barony
	涌波Yurba 	feud.Barony
	雪源Zhoyon 	feud.Barony
}

func (c *羌塘QangtangCounty) BBairab拜惹布() feud.Barony {
	return c.拜惹布Bairab
}
    
func (c *羌塘QangtangCounty) BBurog布若() feud.Barony {
	return c.布若Burog
}
    
func (c *羌塘QangtangCounty) BLaxong拉雄() feud.Barony {
	return c.拉雄Laxong
}
    
func (c *羌塘QangtangCounty) BMargai玛尔盖() feud.Barony {
	return c.玛尔盖Margai
}
    
func (c *羌塘QangtangCounty) BSatso萨错() feud.Barony {
	return c.萨错Satso
}
    
func (c *羌塘QangtangCounty) BYurba涌波() feud.Barony {
	return c.涌波Yurba
}
    
func (c *羌塘QangtangCounty) BZhoyon雪源() feud.Barony {
	return c.雪源Zhoyon
}
    
var CQangtang羌塘 QangtangCounty = &羌塘QangtangCounty{}

func init() {
	f := CQangtang羌塘.(*羌塘QangtangCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1564",
		Title:     "qangtang",
		TitleName: "羌塘",
		TitleCode: "c_qangtang",
		Baronies:  map[string]feud.Barony{},
	}

	f.拜惹布Bairab = BBairab拜惹布
	f.拜惹布Bairab.SetParent(f)
	
	f.布若Burog = BBurog布若
	f.布若Burog.SetParent(f)
	
	f.拉雄Laxong = BLaxong拉雄
	f.拉雄Laxong.SetParent(f)
	
	f.玛尔盖Margai = BMargai玛尔盖
	f.玛尔盖Margai.SetParent(f)
	
	f.萨错Satso = BSatso萨错
	f.萨错Satso.SetParent(f)
	
	f.涌波Yurba = BYurba涌波
	f.涌波Yurba.SetParent(f)
	
	f.雪源Zhoyon = BZhoyon雪源
	f.雪源Zhoyon.SetParent(f)
	
}
