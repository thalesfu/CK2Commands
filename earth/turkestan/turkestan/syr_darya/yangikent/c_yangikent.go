package yangikent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YangikentCounty interface {
    feud.County
    BBogot巴哥特() 	feud.Barony
    BDjend毡的() 	feud.Barony
    BItchankila伊钱卡拉() 	feud.Barony
    BQorovul绮罗乌() 	feud.Barony
    BShovot沙瓦特() 	feud.Barony
    BXazorasp扎佐拉尔() 	feud.Barony
    BXonqa哈纳咖() 	feud.Barony
    BYangiariq杨吉尔奇() 	feud.Barony
    BYangikent养吉干() 	feud.Barony
}

type 养吉干YangikentCounty struct {
	feud.BaseCounty
	巴哥特Bogot 	feud.Barony
	毡的Djend 	feud.Barony
	伊钱卡拉Itchankila 	feud.Barony
	绮罗乌Qorovul 	feud.Barony
	沙瓦特Shovot 	feud.Barony
	扎佐拉尔Xazorasp 	feud.Barony
	哈纳咖Xonqa 	feud.Barony
	杨吉尔奇Yangiariq 	feud.Barony
	养吉干Yangikent 	feud.Barony
}

func (c *养吉干YangikentCounty) BBogot巴哥特() feud.Barony {
	return c.巴哥特Bogot
}
    
func (c *养吉干YangikentCounty) BDjend毡的() feud.Barony {
	return c.毡的Djend
}
    
func (c *养吉干YangikentCounty) BItchankila伊钱卡拉() feud.Barony {
	return c.伊钱卡拉Itchankila
}
    
func (c *养吉干YangikentCounty) BQorovul绮罗乌() feud.Barony {
	return c.绮罗乌Qorovul
}
    
func (c *养吉干YangikentCounty) BShovot沙瓦特() feud.Barony {
	return c.沙瓦特Shovot
}
    
func (c *养吉干YangikentCounty) BXazorasp扎佐拉尔() feud.Barony {
	return c.扎佐拉尔Xazorasp
}
    
func (c *养吉干YangikentCounty) BXonqa哈纳咖() feud.Barony {
	return c.哈纳咖Xonqa
}
    
func (c *养吉干YangikentCounty) BYangiariq杨吉尔奇() feud.Barony {
	return c.杨吉尔奇Yangiariq
}
    
func (c *养吉干YangikentCounty) BYangikent养吉干() feud.Barony {
	return c.养吉干Yangikent
}
    
var CYangikent养吉干 YangikentCounty = &养吉干YangikentCounty{}

func init() {
	f := CYangikent养吉干.(*养吉干YangikentCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1794",
		Title:     "yangikent",
		TitleName: "养吉干",
		TitleCode: "c_yangikent",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴哥特Bogot = BBogot巴哥特
	f.巴哥特Bogot.SetParent(f)
	
	f.毡的Djend = BDjend毡的
	f.毡的Djend.SetParent(f)
	
	f.伊钱卡拉Itchankila = BItchankila伊钱卡拉
	f.伊钱卡拉Itchankila.SetParent(f)
	
	f.绮罗乌Qorovul = BQorovul绮罗乌
	f.绮罗乌Qorovul.SetParent(f)
	
	f.沙瓦特Shovot = BShovot沙瓦特
	f.沙瓦特Shovot.SetParent(f)
	
	f.扎佐拉尔Xazorasp = BXazorasp扎佐拉尔
	f.扎佐拉尔Xazorasp.SetParent(f)
	
	f.哈纳咖Xonqa = BXonqa哈纳咖
	f.哈纳咖Xonqa.SetParent(f)
	
	f.杨吉尔奇Yangiariq = BYangiariq杨吉尔奇
	f.杨吉尔奇Yangiariq.SetParent(f)
	
	f.养吉干Yangikent = BYangikent养吉干
	f.养吉干Yangikent.SetParent(f)
	
}
