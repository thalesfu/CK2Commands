package odessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OdessaCounty interface {
    feud.County
    BBerezivka别列佐夫卡() 	feud.Barony
    BHolovkivka霍洛夫基夫卡() 	feud.Barony
    BIllichivsk伊利奇夫斯克() 	feud.Barony
    BOdessa敖德萨() 	feud.Barony
    BPetroverovka彼得罗韦罗夫卡() 	feud.Barony
    BRozdilna拉兹杰利纳亚() 	feud.Barony
    BYuzhne尤日涅() 	feud.Barony
}

type 敖德萨OdessaCounty struct {
	feud.BaseCounty
	别列佐夫卡Berezivka 	feud.Barony
	霍洛夫基夫卡Holovkivka 	feud.Barony
	伊利奇夫斯克Illichivsk 	feud.Barony
	敖德萨Odessa 	feud.Barony
	彼得罗韦罗夫卡Petroverovka 	feud.Barony
	拉兹杰利纳亚Rozdilna 	feud.Barony
	尤日涅Yuzhne 	feud.Barony
}

func (c *敖德萨OdessaCounty) BBerezivka别列佐夫卡() feud.Barony {
	return c.别列佐夫卡Berezivka
}
    
func (c *敖德萨OdessaCounty) BHolovkivka霍洛夫基夫卡() feud.Barony {
	return c.霍洛夫基夫卡Holovkivka
}
    
func (c *敖德萨OdessaCounty) BIllichivsk伊利奇夫斯克() feud.Barony {
	return c.伊利奇夫斯克Illichivsk
}
    
func (c *敖德萨OdessaCounty) BOdessa敖德萨() feud.Barony {
	return c.敖德萨Odessa
}
    
func (c *敖德萨OdessaCounty) BPetroverovka彼得罗韦罗夫卡() feud.Barony {
	return c.彼得罗韦罗夫卡Petroverovka
}
    
func (c *敖德萨OdessaCounty) BRozdilna拉兹杰利纳亚() feud.Barony {
	return c.拉兹杰利纳亚Rozdilna
}
    
func (c *敖德萨OdessaCounty) BYuzhne尤日涅() feud.Barony {
	return c.尤日涅Yuzhne
}
    
var COdessa敖德萨 OdessaCounty = &敖德萨OdessaCounty{}

func init() {
	f := COdessa敖德萨.(*敖德萨OdessaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1646",
		Title:     "odessa",
		TitleName: "敖德萨",
		TitleCode: "c_odessa",
		Baronies:  map[string]feud.Barony{},
	}

	f.别列佐夫卡Berezivka = BBerezivka别列佐夫卡
	f.别列佐夫卡Berezivka.SetParent(f)
	
	f.霍洛夫基夫卡Holovkivka = BHolovkivka霍洛夫基夫卡
	f.霍洛夫基夫卡Holovkivka.SetParent(f)
	
	f.伊利奇夫斯克Illichivsk = BIllichivsk伊利奇夫斯克
	f.伊利奇夫斯克Illichivsk.SetParent(f)
	
	f.敖德萨Odessa = BOdessa敖德萨
	f.敖德萨Odessa.SetParent(f)
	
	f.彼得罗韦罗夫卡Petroverovka = BPetroverovka彼得罗韦罗夫卡
	f.彼得罗韦罗夫卡Petroverovka.SetParent(f)
	
	f.拉兹杰利纳亚Rozdilna = BRozdilna拉兹杰利纳亚
	f.拉兹杰利纳亚Rozdilna.SetParent(f)
	
	f.尤日涅Yuzhne = BYuzhne尤日涅
	f.尤日涅Yuzhne.SetParent(f)
	
}
