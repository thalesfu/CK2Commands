package pelusia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PelusiaCounty interface {
    feud.County
    BAlsalihiyah萨利希耶() 	feud.Barony
    BIsmaillia伊斯梅利亚() 	feud.Barony
    BPelusia廷尼斯() 	feud.Barony
    BPithom皮索姆() 	feud.Barony
    BSaid塞得() 	feud.Barony
    BSerapeum塞拉比尤姆() 	feud.Barony
    BSile塞勒() 	feud.Barony
}

type 廷尼斯PelusiaCounty struct {
	feud.BaseCounty
	萨利希耶Alsalihiyah 	feud.Barony
	伊斯梅利亚Ismaillia 	feud.Barony
	廷尼斯Pelusia 	feud.Barony
	皮索姆Pithom 	feud.Barony
	塞得Said 	feud.Barony
	塞拉比尤姆Serapeum 	feud.Barony
	塞勒Sile 	feud.Barony
}

func (c *廷尼斯PelusiaCounty) BAlsalihiyah萨利希耶() feud.Barony {
	return c.萨利希耶Alsalihiyah
}
    
func (c *廷尼斯PelusiaCounty) BIsmaillia伊斯梅利亚() feud.Barony {
	return c.伊斯梅利亚Ismaillia
}
    
func (c *廷尼斯PelusiaCounty) BPelusia廷尼斯() feud.Barony {
	return c.廷尼斯Pelusia
}
    
func (c *廷尼斯PelusiaCounty) BPithom皮索姆() feud.Barony {
	return c.皮索姆Pithom
}
    
func (c *廷尼斯PelusiaCounty) BSaid塞得() feud.Barony {
	return c.塞得Said
}
    
func (c *廷尼斯PelusiaCounty) BSerapeum塞拉比尤姆() feud.Barony {
	return c.塞拉比尤姆Serapeum
}
    
func (c *廷尼斯PelusiaCounty) BSile塞勒() feud.Barony {
	return c.塞勒Sile
}
    
var CPelusia廷尼斯 PelusiaCounty = &廷尼斯PelusiaCounty{}

func init() {
	f := CPelusia廷尼斯.(*廷尼斯PelusiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "789",
		Title:     "pelusia",
		TitleName: "廷尼斯",
		TitleCode: "c_pelusia",
		Baronies:  map[string]feud.Barony{},
	}

	f.萨利希耶Alsalihiyah = BAlsalihiyah萨利希耶
	f.萨利希耶Alsalihiyah.SetParent(f)
	
	f.伊斯梅利亚Ismaillia = BIsmaillia伊斯梅利亚
	f.伊斯梅利亚Ismaillia.SetParent(f)
	
	f.廷尼斯Pelusia = BPelusia廷尼斯
	f.廷尼斯Pelusia.SetParent(f)
	
	f.皮索姆Pithom = BPithom皮索姆
	f.皮索姆Pithom.SetParent(f)
	
	f.塞得Said = BSaid塞得
	f.塞得Said.SetParent(f)
	
	f.塞拉比尤姆Serapeum = BSerapeum塞拉比尤姆
	f.塞拉比尤姆Serapeum.SetParent(f)
	
	f.塞勒Sile = BSile塞勒
	f.塞勒Sile.SetParent(f)
	
}
