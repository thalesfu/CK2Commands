package kongu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KonguCounty interface {
    feud.County
    BDharapuri陀罗补罗() 	feud.Barony
    BErode厄卢德() 	feud.Barony
    BHemambika诃曼迦毗() 	feud.Barony
    BKaravur伽罗补() 	feud.Barony
    BNamakkal那摩迦尔() 	feud.Barony
    BPerur吠卢尔() 	feud.Barony
    BSankagiri商迦耆厘() 	feud.Barony
    BSendamangalam西陀曼迦楞() 	feud.Barony
    BTiruppur蒂卢补罗() 	feud.Barony
}

type 贡古KonguCounty struct {
	feud.BaseCounty
	陀罗补罗Dharapuri 	feud.Barony
	厄卢德Erode 	feud.Barony
	诃曼迦毗Hemambika 	feud.Barony
	伽罗补Karavur 	feud.Barony
	那摩迦尔Namakkal 	feud.Barony
	吠卢尔Perur 	feud.Barony
	商迦耆厘Sankagiri 	feud.Barony
	西陀曼迦楞Sendamangalam 	feud.Barony
	蒂卢补罗Tiruppur 	feud.Barony
}

func (c *贡古KonguCounty) BDharapuri陀罗补罗() feud.Barony {
	return c.陀罗补罗Dharapuri
}
    
func (c *贡古KonguCounty) BErode厄卢德() feud.Barony {
	return c.厄卢德Erode
}
    
func (c *贡古KonguCounty) BHemambika诃曼迦毗() feud.Barony {
	return c.诃曼迦毗Hemambika
}
    
func (c *贡古KonguCounty) BKaravur伽罗补() feud.Barony {
	return c.伽罗补Karavur
}
    
func (c *贡古KonguCounty) BNamakkal那摩迦尔() feud.Barony {
	return c.那摩迦尔Namakkal
}
    
func (c *贡古KonguCounty) BPerur吠卢尔() feud.Barony {
	return c.吠卢尔Perur
}
    
func (c *贡古KonguCounty) BSankagiri商迦耆厘() feud.Barony {
	return c.商迦耆厘Sankagiri
}
    
func (c *贡古KonguCounty) BSendamangalam西陀曼迦楞() feud.Barony {
	return c.西陀曼迦楞Sendamangalam
}
    
func (c *贡古KonguCounty) BTiruppur蒂卢补罗() feud.Barony {
	return c.蒂卢补罗Tiruppur
}
    
var CKongu贡古 KonguCounty = &贡古KonguCounty{}

func init() {
	f := CKongu贡古.(*贡古KonguCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1201",
		Title:     "kongu",
		TitleName: "贡古",
		TitleCode: "c_kongu",
		Baronies:  map[string]feud.Barony{},
	}

	f.陀罗补罗Dharapuri = BDharapuri陀罗补罗
	f.陀罗补罗Dharapuri.SetParent(f)
	
	f.厄卢德Erode = BErode厄卢德
	f.厄卢德Erode.SetParent(f)
	
	f.诃曼迦毗Hemambika = BHemambika诃曼迦毗
	f.诃曼迦毗Hemambika.SetParent(f)
	
	f.伽罗补Karavur = BKaravur伽罗补
	f.伽罗补Karavur.SetParent(f)
	
	f.那摩迦尔Namakkal = BNamakkal那摩迦尔
	f.那摩迦尔Namakkal.SetParent(f)
	
	f.吠卢尔Perur = BPerur吠卢尔
	f.吠卢尔Perur.SetParent(f)
	
	f.商迦耆厘Sankagiri = BSankagiri商迦耆厘
	f.商迦耆厘Sankagiri.SetParent(f)
	
	f.西陀曼迦楞Sendamangalam = BSendamangalam西陀曼迦楞
	f.西陀曼迦楞Sendamangalam.SetParent(f)
	
	f.蒂卢补罗Tiruppur = BTiruppur蒂卢补罗
	f.蒂卢补罗Tiruppur.SetParent(f)
	
}
