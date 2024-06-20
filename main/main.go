package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/thalesfu/CK2Commands/culture"
	"github.com/thalesfu/CK2Commands/family/bi"
	"github.com/thalesfu/CK2Commands/family/chu"
	"github.com/thalesfu/CK2Commands/family/fu"
	"github.com/thalesfu/CK2Commands/family/li"
	"github.com/thalesfu/CK2Commands/family/lin"
	"github.com/thalesfu/CK2Commands/family/lou"
	"github.com/thalesfu/CK2Commands/family/wu"
	"github.com/thalesfu/CK2Commands/family/wuyibu"
	"github.com/thalesfu/CK2Commands/feudal"
	"github.com/thalesfu/CK2Commands/historydynasty"
	"github.com/thalesfu/CK2Commands/historypeople"
	"github.com/thalesfu/CK2Commands/job"
	"github.com/thalesfu/CK2Commands/people"
	"github.com/thalesfu/CK2Commands/religion"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
	utils2 "github.com/thalesfu/nebulagolang/utils"
	"github.com/thalesfu/paradoxtools/utils"
	"log"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const ck2Folder = "/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II"
const saveFolder = "/Users/thalesfu/Documents/Paradox Interactive/Crusader Kings II/save games"
const saveFolder2 = "/Users/thalesfu/Windows/steam/userdata/94993760/203770/remote/save games"

var CoreFamily = map[int]string{
	1000103393: "lou",
	1000103382: "yuan",
	1000103379: "chu",
	1000103374: "wu",
	1000103360: "zhang",
	1000103348: "lin",
	1000103339: "bi",
	1000103336: "yin",
	1051150:    "li",
	1040018:    "Óengus",
}

func main() {
	defer ck2nebula.SPACE.Nebula.Close()
	forceLoadDataMode := false
	watchMode := false
	titleMode := false
	cultureMode := false
	religionMode := false
	nameMode := false
	washReligionMode := false

	if len(os.Args) > 0 {
		for _, arg := range os.Args {
			if arg == "-f" {
				forceLoadDataMode = true
				continue
			}

			if arg == "-w" {
				watchMode = true
				continue
			}

			if arg == "-t" {
				titleMode = true
				continue
			}

			if arg == "-c" {
				cultureMode = true
				continue
			}

			if arg == "-r" {
				religionMode = true
				continue
			}

			if arg == "-n" {
				nameMode = true
				continue
			}

			if arg == "-wr" {
				washReligionMode = true
				continue
			}
		}
	}

	if washReligionMode {
		people.BuildWashReligionScript(ck2nebula.SPACE, religion.Religion_东方宗教_道教_taoist)
		return
	} else if nameMode {
		people.ChangePeopleName()
	} else if religionMode {
		people.Taoist()
	} else if cultureMode {
		people.HanPictish()
	} else if titleMode {
		feudal.BuildTitle()
	} else if watchMode {

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()

		done := make(chan bool)
		stop := make(chan bool)

		go func() {
			for {
				select {
				case event := <-watcher.Events:
					fmt.Println("Event:", time.Now().Format("2006-01-02 15:04:05"), event)
					if event.Op&fsnotify.Create == fsnotify.Create {
						fn := filepath.Base(event.Name)
						if strings.HasSuffix(event.Name, ".ck2") && fn != "oldautosave.ck2" && fn != "olderautosave.ck2" {
							loadAndAutoBuild(forceLoadDataMode)
						}
					}
				case err := <-watcher.Errors:
					fmt.Println("Error:", err)
				case <-stop:
					fmt.Println("Stopping watcher...")
					done <- true
					return
				}
			}
		}()

		err = watcher.Add(saveFolder)
		if err != nil {
			log.Fatal(err)
		}
		err = watcher.Add(saveFolder2)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Watching", saveFolder, "for changes...")

		// 在另一个协程中等待用户输入
		go func() {
			fmt.Println("Type 'exit' to stop the program.")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				text := scanner.Text()
				if text == "exit" {
					stop <- true
					return
				}
			}
		}()

		// 等待监视器协程优雅地停止
		<-done
		fmt.Println("Program exited.")
	} else {
		loadAndAutoBuild(forceLoadDataMode)
	}

	//people.CureFriends(ck2nebula.SPACE, player)

	//people.SetFriendsAndFriendsFriendsBuildPeopleFlag(ck2nebula.SPACE, player)

	//people.PollinateFriendsAndFriendsFriends(ck2nebula.SPACE, player)

	//BuildFriends(player)
}

func loadAndAutoBuild(forceLoadData bool) {
	start := time.Now()
	story, player, err := GetStory(forceLoadData)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(story.PlayerName)

	people.AutoBuild(ck2nebula.SPACE, player, CoreFamily)
	people.ForceCurePeople(ck2nebula.SPACE, player, CoreFamily)
	end := time.Now()
	duration := end.Sub(start)
	log.Printf("%sDuration: %f seconds from %s to %s %s", utils2.PrintColorCyan, duration.Seconds(), start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"), utils2.PrintColorReset)
}

func BuildFriends(player *ck2nebula.People) {
	fr := player.GetFriends(ck2nebula.SPACE)

	if fr.Ok {
		friends := make([]*ck2nebula.People, 0)
		for _, f := range fr.Data {
			friends = append(friends, f)
		}
		people.BuildHealthAndFertility(ck2nebula.SPACE, friends...)
	}
}

func GetStory(force bool) (*ck2nebula.Story, *ck2nebula.People, error) {

	sr := ck2nebula.GetLatestStory(ck2nebula.SPACE)

	if !sr.Ok && !errors.As(sr.Err, &nebulagolang.NoDataErr) {
		return nil, nil, sr.Err
	}

	dirs := []string{saveFolder, saveFolder2}

	var filePath string
	var modifTime int64

	for _, d := range dirs {
		fs, err := os.ReadDir(d)
		if err != nil {
			return nil, nil, err
		}

		for _, f := range fs {
			if !f.IsDir() && strings.HasSuffix(f.Name(), ".ck2") {
				info, _ := f.Info()
				if info.ModTime().Unix() > modifTime {
					filePath = strings.ReplaceAll(filepath.Join(d, f.Name()), "\\", "/")
					modifTime = info.ModTime().Unix()
				}
			}
		}
	}

	if force || errors.As(sr.Err, &nebulagolang.NoDataErr) || !isSameStory(filePath, sr.Data) {
		log.Printf("%sloading save file \"%s\"%s\n", utils2.PrintColorCyan, filePath, utils2.PrintColorReset)
		ck2nebula.BuildStory(ck2Folder, filePath, culture.CultureMap, religion.ReligionMap, historypeople.HistoryPeopleMap, historydynasty.HistoryDynastyMap)

		sr = ck2nebula.GetLatestStory(ck2nebula.SPACE)

		if !sr.Ok {
			return nil, nil, sr.Err
		}
	}

	pr := sr.Data.GetPlayer(ck2nebula.SPACE)

	if !pr.Ok {
		return nil, nil, pr.Err
	}

	return sr.Data, pr.Data, nil
}

func isSameStory(fp string, story *ck2nebula.Story) bool {

	if fp != story.FilePath {
		return false
	}

	hash, err := utils.GetFileHash(fp)

	if err != nil {
		return false
	}

	if hash != story.FileHash {
		return false
	}

	fileInfo, err := os.Stat(fp)

	if err != nil {
		return false
	}

	return fileInfo.ModTime().Format("2006-01-02 15:04:05") == story.FileUpdateTime.Format("2006-01-02 15:04:05")
}

func buildjobs() {
	people.BuildMyLoad(job.GetMyloads()...)
	people.BuildAmbassador(job.GetAmbassadors()...)
	people.BuildGeneral(job.GetGenerals()...)
	people.BuildManager(job.GetManagers()...)
	people.BuildSpy(job.GetSpys()...)
	people.BuildFather(job.GetFathers()...)
	people.BuildDoctor(job.GetDoctors()...)
	people.BuildRandom(job.GetRandoms()...)
}

func makefriends() {
	people.MakeFriends(getMyBrothersAndSisters(),
		getMySonsAndDaughters(),
		getMyBigFamiliesFriends(),
		bi.GetBiFriends(),
		wu.GetFriends(),
		lou.GetLouFriends(),
		li.GetFriends(),
		chu.GetFriends(),
	)
}

func getMySonsAndDaughters() []int {
	var peoples []int

	peoples = append(peoples, fu.Me)
	peoples = append(peoples, fu.MySon1)
	peoples = append(peoples, fu.MySon2)
	peoples = append(peoples, fu.MySon3)
	peoples = append(peoples, fu.MySon4)
	peoples = append(peoples, fu.MySon5)
	peoples = append(peoples, fu.MySon6)
	peoples = append(peoples, fu.MySon7)
	peoples = append(peoples, fu.MySon8)
	peoples = append(peoples, fu.MySon9)
	peoples = append(peoples, fu.MySon10)
	peoples = append(peoples, fu.MyDaughter1)
	peoples = append(peoples, fu.MyDaughter2)
	peoples = append(peoples, fu.MyDaughter3)
	peoples = append(peoples, fu.MyDaughter4)
	peoples = append(peoples, fu.MyDaughter5)
	peoples = append(peoples, fu.MyDaughter6)
	peoples = append(peoples, fu.MyDaughter7)
	peoples = append(peoples, fu.MyDaughter8)
	peoples = append(peoples, fu.MyDaughter9)
	peoples = append(peoples, fu.MyDaughter10)

	return peoples
}

func getMyBrothersAndSisters() []int {
	var peoples []int

	peoples = append(peoples, fu.Me)
	peoples = append(peoples, fu.MyWife)
	peoples = append(peoples, fu.MyBrother1)
	peoples = append(peoples, fu.MyBrother2)
	peoples = append(peoples, fu.MyBrother3)
	peoples = append(peoples, fu.MyBrother4)
	peoples = append(peoples, fu.MyBrother5)
	peoples = append(peoples, fu.MyBrother6)
	peoples = append(peoples, fu.MyBrother7)
	peoples = append(peoples, fu.MyBrother8)
	peoples = append(peoples, fu.MyBrother9)
	peoples = append(peoples, fu.MyBrother10)
	peoples = append(peoples, fu.MySister1)
	peoples = append(peoples, fu.MySister2)
	peoples = append(peoples, fu.MySister3)
	peoples = append(peoples, fu.MySister4)
	peoples = append(peoples, fu.MySister5)
	peoples = append(peoples, fu.MySister6)
	peoples = append(peoples, fu.MySister7)
	peoples = append(peoples, fu.MySister8)
	peoples = append(peoples, fu.MySister9)
	peoples = append(peoples, fu.MySister10)

	return peoples
}

func getMyBigFamiliesFriends() []int {
	var peoples []int

	peoples = append(peoples, fu.Me)
	peoples = append(peoples, bi.M毕元振2608618)
	peoples = append(peoples, li.M李适217709)
	peoples = append(peoples, li.M李诵217710)
	peoples = append(peoples, lou.M楼彦玮2611890)

	return peoples
}

func getMyBigLoads() []int {
	var peoples []int

	peoples = append(peoples, fu.Me)
	peoples = append(peoples, 2830311)
	peoples = append(peoples, 2840690)
	peoples = append(peoples, lin.M林承庆10960227)

	return peoples
}

func getGermanFriends() []int {
	var peoples []int

	peoples = append(peoples, 2845721)
	peoples = append(peoples, 2840086)

	return peoples
}

func getYiwubuFriends() []int {
	var peoples []int

	peoples = append(peoples, wuyibu.F善理_守礼II)
	peoples = append(peoples, wuyibu.M多梅希_守礼II)
	peoples = append(peoples, 2843524)
	peoples = append(peoples, 2841943)
	peoples = append(peoples, 2838684)
	peoples = append(peoples, 2847861)
	peoples = append(peoples, 2847861)
	peoples = append(peoples, 2864693)
	peoples = append(peoples, 2853577)
	peoples = append(peoples, 2856195)
	peoples = append(peoples, 2850307)

	return peoples
}

func pollinate() {
	//yuan.Hy()
	//pictish.Hy()
	//fu.Hy()
	//lin.Pollinate()
	//wu.Pollinate()
	//bi.Pollinate()
	//lou.Pollinate()
}
