package epirus

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus/epirus/arta"
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus/epirus/epeiros"
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus/epirus/naupactus"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EpirusDuke interface {
	feud.Duke
	CArta阿尔塔() arta.ArtaCounty
	CEpeiros伊庇鲁斯() epeiros.EpeirosCounty
	CNaupactus纳夫帕克托斯() naupactus.NaupactusCounty
}

type 伊庇鲁斯EpirusDuke struct {
	feud.BaseDuke
	阿尔塔Arta         arta.ArtaCounty
	伊庇鲁斯Epeiros     epeiros.EpeirosCounty
	纳夫帕克托斯Naupactus naupactus.NaupactusCounty
}

func (d *伊庇鲁斯EpirusDuke) CArta阿尔塔() arta.ArtaCounty {
	return d.阿尔塔Arta
}

func (d *伊庇鲁斯EpirusDuke) CEpeiros伊庇鲁斯() epeiros.EpeirosCounty {
	return d.伊庇鲁斯Epeiros
}

func (d *伊庇鲁斯EpirusDuke) CNaupactus纳夫帕克托斯() naupactus.NaupactusCounty {
	return d.纳夫帕克托斯Naupactus
}

var DEpirus伊庇鲁斯 EpirusDuke = &伊庇鲁斯EpirusDuke{}

func init() {
	f := DEpirus伊庇鲁斯.(*伊庇鲁斯EpirusDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "epirus",
		TitleName: "伊庇鲁斯",
		TitleCode: "d_epirus",
		Counties:  map[string]feud.County{},
	}

	f.阿尔塔Arta = arta.CArta阿尔塔
	f.阿尔塔Arta.SetParent(f)

	f.伊庇鲁斯Epeiros = epeiros.CEpeiros伊庇鲁斯
	f.伊庇鲁斯Epeiros.SetParent(f)

	f.纳夫帕克托斯Naupactus = naupactus.CNaupactus纳夫帕克托斯
	f.纳夫帕克托斯Naupactus.SetParent(f)

}
