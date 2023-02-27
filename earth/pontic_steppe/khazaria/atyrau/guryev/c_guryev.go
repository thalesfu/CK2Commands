package guryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GuryevCounty interface {
    feud.County
    BArlik艾里克() 	feud.Barony
    BAtyrau阿特劳() 	feud.Barony
    BBesikty比谢克特() 	feud.Barony
    BKarabatyr哈拉巴特尔() 	feud.Barony
    BMantyube曼图比() 	feud.Barony
    BSarayjuk小萨赖() 	feud.Barony
    BYesmakan叶斯马坎() 	feud.Barony
    BZhanakush扎那库什() 	feud.Barony
}

type 阿特劳GuryevCounty struct {
	feud.BaseCounty
	艾里克Arlik 	feud.Barony
	阿特劳Atyrau 	feud.Barony
	比谢克特Besikty 	feud.Barony
	哈拉巴特尔Karabatyr 	feud.Barony
	曼图比Mantyube 	feud.Barony
	小萨赖Sarayjuk 	feud.Barony
	叶斯马坎Yesmakan 	feud.Barony
	扎那库什Zhanakush 	feud.Barony
}

func (c *阿特劳GuryevCounty) BArlik艾里克() feud.Barony {
	return c.艾里克Arlik
}
    
func (c *阿特劳GuryevCounty) BAtyrau阿特劳() feud.Barony {
	return c.阿特劳Atyrau
}
    
func (c *阿特劳GuryevCounty) BBesikty比谢克特() feud.Barony {
	return c.比谢克特Besikty
}
    
func (c *阿特劳GuryevCounty) BKarabatyr哈拉巴特尔() feud.Barony {
	return c.哈拉巴特尔Karabatyr
}
    
func (c *阿特劳GuryevCounty) BMantyube曼图比() feud.Barony {
	return c.曼图比Mantyube
}
    
func (c *阿特劳GuryevCounty) BSarayjuk小萨赖() feud.Barony {
	return c.小萨赖Sarayjuk
}
    
func (c *阿特劳GuryevCounty) BYesmakan叶斯马坎() feud.Barony {
	return c.叶斯马坎Yesmakan
}
    
func (c *阿特劳GuryevCounty) BZhanakush扎那库什() feud.Barony {
	return c.扎那库什Zhanakush
}
    
var CGuryev阿特劳 GuryevCounty = &阿特劳GuryevCounty{}

func init() {
	f := CGuryev阿特劳.(*阿特劳GuryevCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "618",
		Title:     "guryev",
		TitleName: "阿特劳",
		TitleCode: "c_guryev",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾里克Arlik = BArlik艾里克
	f.艾里克Arlik.SetParent(f)
	
	f.阿特劳Atyrau = BAtyrau阿特劳
	f.阿特劳Atyrau.SetParent(f)
	
	f.比谢克特Besikty = BBesikty比谢克特
	f.比谢克特Besikty.SetParent(f)
	
	f.哈拉巴特尔Karabatyr = BKarabatyr哈拉巴特尔
	f.哈拉巴特尔Karabatyr.SetParent(f)
	
	f.曼图比Mantyube = BMantyube曼图比
	f.曼图比Mantyube.SetParent(f)
	
	f.小萨赖Sarayjuk = BSarayjuk小萨赖
	f.小萨赖Sarayjuk.SetParent(f)
	
	f.叶斯马坎Yesmakan = BYesmakan叶斯马坎
	f.叶斯马坎Yesmakan.SetParent(f)
	
	f.扎那库什Zhanakush = BZhanakush扎那库什
	f.扎那库什Zhanakush.SetParent(f)
	
}
