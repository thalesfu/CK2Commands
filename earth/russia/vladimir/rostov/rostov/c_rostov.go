package rostov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RostovCounty interface {
    feud.County
    BKondakovo孔达科沃() 	feud.Barony
    BLvy利维() 	feud.Barony
    BRezanka列赞卡() 	feud.Barony
    BRostov罗斯托夫() 	feud.Barony
    BSarskoyegorodishche萨尔斯科耶戈罗季谢() 	feud.Barony
    BSpasoyakovlevsky斯帕索雅科夫列夫斯基() 	feud.Barony
    BYepikharka叶皮哈尔卡() 	feud.Barony
}

type 罗斯托夫RostovCounty struct {
	feud.BaseCounty
	孔达科沃Kondakovo 	feud.Barony
	利维Lvy 	feud.Barony
	列赞卡Rezanka 	feud.Barony
	罗斯托夫Rostov 	feud.Barony
	萨尔斯科耶戈罗季谢Sarskoyegorodishche 	feud.Barony
	斯帕索雅科夫列夫斯基Spasoyakovlevsky 	feud.Barony
	叶皮哈尔卡Yepikharka 	feud.Barony
}

func (c *罗斯托夫RostovCounty) BKondakovo孔达科沃() feud.Barony {
	return c.孔达科沃Kondakovo
}
    
func (c *罗斯托夫RostovCounty) BLvy利维() feud.Barony {
	return c.利维Lvy
}
    
func (c *罗斯托夫RostovCounty) BRezanka列赞卡() feud.Barony {
	return c.列赞卡Rezanka
}
    
func (c *罗斯托夫RostovCounty) BRostov罗斯托夫() feud.Barony {
	return c.罗斯托夫Rostov
}
    
func (c *罗斯托夫RostovCounty) BSarskoyegorodishche萨尔斯科耶戈罗季谢() feud.Barony {
	return c.萨尔斯科耶戈罗季谢Sarskoyegorodishche
}
    
func (c *罗斯托夫RostovCounty) BSpasoyakovlevsky斯帕索雅科夫列夫斯基() feud.Barony {
	return c.斯帕索雅科夫列夫斯基Spasoyakovlevsky
}
    
func (c *罗斯托夫RostovCounty) BYepikharka叶皮哈尔卡() feud.Barony {
	return c.叶皮哈尔卡Yepikharka
}
    
var CRostov罗斯托夫 RostovCounty = &罗斯托夫RostovCounty{}

func init() {
	f := CRostov罗斯托夫.(*罗斯托夫RostovCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "574",
		Title:     "rostov",
		TitleName: "罗斯托夫",
		TitleCode: "c_rostov",
		Baronies:  map[string]feud.Barony{},
	}

	f.孔达科沃Kondakovo = BKondakovo孔达科沃
	f.孔达科沃Kondakovo.SetParent(f)
	
	f.利维Lvy = BLvy利维
	f.利维Lvy.SetParent(f)
	
	f.列赞卡Rezanka = BRezanka列赞卡
	f.列赞卡Rezanka.SetParent(f)
	
	f.罗斯托夫Rostov = BRostov罗斯托夫
	f.罗斯托夫Rostov.SetParent(f)
	
	f.萨尔斯科耶戈罗季谢Sarskoyegorodishche = BSarskoyegorodishche萨尔斯科耶戈罗季谢
	f.萨尔斯科耶戈罗季谢Sarskoyegorodishche.SetParent(f)
	
	f.斯帕索雅科夫列夫斯基Spasoyakovlevsky = BSpasoyakovlevsky斯帕索雅科夫列夫斯基
	f.斯帕索雅科夫列夫斯基Spasoyakovlevsky.SetParent(f)
	
	f.叶皮哈尔卡Yepikharka = BYepikharka叶皮哈尔卡
	f.叶皮哈尔卡Yepikharka.SetParent(f)
	
}
