package main

import (
	"fmt"
	"fut-otw/futbin"
	"fut-otw/sofascore"
	"log"
)

type Player struct {
	BaseId      string
	OtwId       string
	SofascoreId string
}

var Players = [...]Player{
	Player{"20801", "50352449", "750"},
	Player{"167664", "50499312", "19438"},
	Player{"204485", "50536133", "158213"},
	Player{"172879", "50504527", "13108"},
	Player{"201995", "50533643", "152276"},
	Player{"210514", "50542162", "138892"},
	Player{"178518", "50510166", "44001"},
	Player{"177413", "50509061", "35612"},
	Player{"205498", "50537146", "132874"},
	Player{"208808", "50540456", "106161"},
	Player{"209658", "50541306", "184661"},
	Player{"213565", "50545213", "191182"},
	Player{"220971", "50552619", "365772"},
	Player{"222737", "50554385", "551772"},
	Player{"209297", "50540945", "243623"},
	Player{"227055", "67335919", "364618"},
	Player{"187491", "50519139", "24601"},
	Player{"201956", "50533604", "143398"},
	Player{"203890", "50535538", "70630"},
	Player{"231943", "50563591", "840217"},
	Player{"196889", "67305753", "98058"},
	Player{"221639", "50553287", "175753"},
	Player{"225663", "50557311", "318963"},
	Player{"211591", "50543239", "318653"},
}

func main() {
	for _, p := range Players {
		link := generateLink(p.BaseId, p.OtwId)

		price, err := futbin.GetPrice(p.OtwId)
		if err != nil {
			log.Fatalln(err)
		}

		ratings, err := sofascore.GetRatings(p.SofascoreId)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(link, price, ratings)
	}
}

func generateLink(baseId, otwId string) string {
	return fmt.Sprintf("https://www.easports.com/fifa/ultimate-team/fut/database/player/%s#%s", baseId, otwId)
}
