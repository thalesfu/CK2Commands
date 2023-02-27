package gurjara_mandala

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/gurjara_mandala/khetaka"
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/gurjara_mandala/mohadavasaka"
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/gurjara_mandala/sarasvata_mandala"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Gurjara_mandalaDuke interface {
    feud.Duke
    CKhetaka契吒迦() 	khetaka.KhetakaCounty
    CMohadavasaka慕诃陀婆娑迦() 	mohadavasaka.MohadavasakaCounty
    CSarasvata_mandala萨罗萨伐多曼荼罗() 	sarasvata_mandala.Sarasvata_mandalaCounty
}

type 瞿折罗曼荼罗Gurjara_mandalaDuke struct {
	feud.BaseDuke
	契吒迦Khetaka 	khetaka.KhetakaCounty
	慕诃陀婆娑迦Mohadavasaka 	mohadavasaka.MohadavasakaCounty
	萨罗萨伐多曼荼罗Sarasvata_mandala 	sarasvata_mandala.Sarasvata_mandalaCounty
}

func (d *瞿折罗曼荼罗Gurjara_mandalaDuke) CKhetaka契吒迦() khetaka.KhetakaCounty {
	return d.契吒迦Khetaka
}
    
func (d *瞿折罗曼荼罗Gurjara_mandalaDuke) CMohadavasaka慕诃陀婆娑迦() mohadavasaka.MohadavasakaCounty {
	return d.慕诃陀婆娑迦Mohadavasaka
}
    
func (d *瞿折罗曼荼罗Gurjara_mandalaDuke) CSarasvata_mandala萨罗萨伐多曼荼罗() sarasvata_mandala.Sarasvata_mandalaCounty {
	return d.萨罗萨伐多曼荼罗Sarasvata_mandala
}
    
var DGurjara_mandala瞿折罗曼荼罗 Gurjara_mandalaDuke = &瞿折罗曼荼罗Gurjara_mandalaDuke{}

func init() {
	f := DGurjara_mandala瞿折罗曼荼罗.(*瞿折罗曼荼罗Gurjara_mandalaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gurjara_mandala",
		TitleName: "瞿折罗曼荼罗",
		TitleCode: "d_gurjara_mandala",
		Counties:  map[string]feud.County{},
	}

	f.契吒迦Khetaka = khetaka.CKhetaka契吒迦
	f.契吒迦Khetaka.SetParent(f)
	
	f.慕诃陀婆娑迦Mohadavasaka = mohadavasaka.CMohadavasaka慕诃陀婆娑迦
	f.慕诃陀婆娑迦Mohadavasaka.SetParent(f)
	
	f.萨罗萨伐多曼荼罗Sarasvata_mandala = sarasvata_mandala.CSarasvata_mandala萨罗萨伐多曼荼罗
	f.萨罗萨伐多曼荼罗Sarasvata_mandala.SetParent(f)
	
}
