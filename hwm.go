package main

import (
	"sort"
	"time"
	"math/rand"
	"fmt"
)

type Supply struct {
	ID     int
	Geo    string // sh,bj,sz
	Gender string // F or M
	Count  int
}

type Demand struct {
	ID         int
	Geo        string
	Gender     string
	Eligible   int
	Count      int
	Rate       float64
}

type DemandLink struct {
	SupplyId int
	DemandId int
}

func HwmPlan(supplies[]Supply, demands []*Demand, demandLinks []DemandLink) {
	for _, demand := range demands {
		for _, link := range demandLinks {
			if link.DemandId == demand.ID {
				demand.Eligible += supplies[link.SupplyId].Count
			}
		}
	}

	sort.Slice(demands, func(i, j int) bool {
		return float64(demands[i].Count)/float64(demands[i].Eligible) > float64(demands[j].Count)/float64(demands[i].Eligible)
	})

	for _, demand := range demands {
		for _, link := range demandLinks {
			if link.DemandId != demand.ID {
				continue
			}

			// calculate total remain of one demand
			totalRemain := 0
			for _, link := range demandLinks {
				if link.DemandId == demand.ID {
					totalRemain += supplies[link.SupplyId].Count
				}
			}

			// assign rate
			if totalRemain < demand.Count {
				demand.Rate = 1.0
			} else {
				demand.Rate = float64(demand.Count) / float64(totalRemain)
			}

		    // re-calculate supplies
		    supplies[link.SupplyId].Count = int(float64(supplies[link.SupplyId].Count) * (1 - demand.Rate))

		}
	}
}

func HwmServe(orderedDemands []*Demand) *Demand {
	truncatedIndex := 0
	rateSum := 0.0
	for i, demand := range orderedDemands {
		rateSum += demand.Rate
		if rateSum >= 1 {
			truncatedIndex = i - 1
			break
		}
	}
	if len(orderedDemands) > truncatedIndex + 1 {
		orderedDemands[truncatedIndex+1].Rate = 1 - rateSum
	}

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	rVal := rand.Float64()
	randSum := 0.0
	for _, demand := range orderedDemands {
		randSum += demand.Rate
		if rVal > randSum {
			return demand
		}
	}

	return nil
}


func main() {
	supplies := []Supply{
		Supply{0, "sh", "F", 100},
		Supply{1, "bj", "F", 200},
		Supply{2, "bj", "M", 300},
		Supply{3, "sz", "F", 200},
		Supply{4, "sz", "M", 200},
		Supply{5, "sh", "M", 100},
	}

	demands := []*Demand{
		&Demand{0, "bj", "M", 0, 50, 0},
		&Demand{1, "sh", "M", 0, 50, 0},
		&Demand{2, "bj", "M", 0, 200, 0},
		&Demand{3, "sh", "F", 0, 50, 0},
		&Demand{4, "bj", "M", 0, 80, 0},

	}

	demandLinks := []DemandLink{}
	for _, demand := range demands {
		for _, supply := range supplies {
			if demand.Gender == supply.Gender && demand.Geo == supply.Geo {
				demandLinks = append(demandLinks, DemandLink{SupplyId: supply.ID, DemandId: demand.ID})
			}
		}
	}

	HwmPlan(supplies, demands, demandLinks)

	for _, demand := range demands {
		fmt.Println(demand)
	}

	fmt.Println("************")

	adFrom := supplies[2]
	candidateDemands := []*Demand{}
	for _, link := range demandLinks {
		if link.SupplyId == adFrom.ID {
			candidateDemands = append(candidateDemands, demands[link.DemandId])
		}
	}
	for _, c := range candidateDemands {
		fmt.Println(c)
	}

	sort.Slice(candidateDemands, func(i, j int) bool { return candidateDemands[i].Rate > candidateDemands[j].Rate })

	for i := 0; i < 100; i++ {
		demand := HwmServe(candidateDemands)
		if demand == nil {
			fmt.Print("n")
		} else {
			fmt.Print(demand.ID)
		}

	}

}