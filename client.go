package main

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/kr/pretty"
)

type Donor struct {
	DonorID               string        `json:"donorId"`
	DonorName             string        `json:"donorName"`
	DonorPhotoURL         string        `json:"donorPhotoURL"`
	DonorProfileURL       string        `json:"donorProfileURL"`
	Favorites             []interface{} `json:"favorites"`
	GivingPages           []interface{} `json:"givingPages"`
	MemberGivingPages     []interface{} `json:"memberGivingPages"`
	NumDonorsAcquired     string        `json:"numDonorsAcquired"`
	NumProposalsSupported string        `json:"numProposalsSupported"`
	NumStudentsSupported  string        `json:"numStudentsSupported"`
	ProposalMessages      []interface{} `json:"proposalMessages"`
	Proposals             []interface{} `json:"proposals"`
	Stickers              []struct {
		Description string `json:"description"`
		ImageURL    string `json:"imageURL"`
		Name        string `json:"name"`
		Value       string `json:"value"`
	} `json:"stickers"`
	TotalFavorites string `json:"totalFavorites"`
	TotalProposals string `json:"totalProposals"`
}

type DonorParams struct {
	Key     string `url:"APIKey"`
	DonorID string `url:"donorid"`
}

type DonorError struct{}

func main() {
	const donorschooseURI = "https://api.donorschoose.org/common/json-donor.html"

	donor := new(Donor)
	params := &DonorParams{Key: "DONORSCHOOSE", DonorID: "4433774"}

	resp, err := sling.New().Get(donorschooseURI).QueryStruct(params).ReceiveSuccess(donor)

	if err != nil {
		fmt.Println("hello")
		fmt.Println(resp)
	} else {
		fmt.Println("err")
		fmt.Println(err)
	}

	fmt.Printf("%# v", pretty.Formatter(*donor))
}
