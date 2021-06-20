package keystone

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/opensds/multi-cloud/aksk/pkg/model"
	"github.com/opensds/multi-cloud/aksk/pkg/utils"
	pb "github.com/opensds/multi-cloud/aksk/proto"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const KEYSTONE_URL = "http://192.168.20.108/identity/v3/credentials"

type blob struct {
	Access string `json:"access"`
	Secret string `json:"secret"`
}


type CredBody struct {
	ProjectId string `json:"project_id"`
	UserId string `json:"user_id"`
	Blob string `json:"blob"`
	Type string `json:"type"`
}

type Credential struct {
	Credential CredBody `json:"credential"`
}

type Client struct {

}

func CreateAKSK(aksk *model.AkSk, req *pb.CreateAkSkRequest) (*http.Response, error) {

	akey:= utils.GenerateRandomString(16)
	skey := utils.GenerateRandomString(32)

	blb := blob{Access: akey, Secret: skey}
	log.Info("RAJAT - AK SK  - " , akey , skey)

	byts, err := json.Marshal(blb)
	if err != nil {
		panic(err)
	}

	in := string(byts)
	byts, err = json.Marshal(in)
	if err != nil {
		panic(err)
	}

	u, err := json.Marshal(Credential{Credential: CredBody{ProjectId: aksk.ProjectId,
		UserId: aksk.UserId,
		Blob: in,
		Type: "ec2"}})

	fmt.Println(string(u))
	client := &http.Client{}

	postreq , err := http.NewRequest("POST", KEYSTONE_URL, bytes.NewBuffer(u))
	postreq.Header.Add("X-Auth-Token", req.Aksk.Token)
	postreq.Header.Set("Content-Type", "application/json")


	akskresp, _ := client.Do(postreq)
	log.Info("RAJAT - RESPONSE  - " , akskresp)

	defer akskresp.Body.Close()

	if err != nil {
	   return nil, err
	}

	log.Info("RAJAT - Create AKSK Response - ", akskresp)
	return akskresp, nil
}

func ListAKSK(ctx context.Context)(*http.Response, error) {

	postreq , err := http.NewRequest("GET", KEYSTONE_URL, bytes.NewBuffer(nil))
	postreq.Header.Add("X-Auth-Token", req.Aksk.Token)
	postreq.Header.Set("Content-Type", "application/json")

	akskresp, err := http.Get(KEYSTONE_URL)
	if err != nil {
		return nil, err
	}

	bodyBytes, _ := ioutil.ReadAll(akskresp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)
	defer akskresp.Body.Close()
	return akskresp, nil
}

func DeleteAKSK(ctx context.Context, req *pb.DeleteAkSkRequest)(*http.Response, error){

	client := &http.Client{}
	log.Info("RAJAT - DELETE  - URL " , KEYSTONE_URL+"/" + req.GetId())
	getreq , err := http.NewRequest("DELETE", KEYSTONE_URL+"/" + req.GetId(), bytes.NewBuffer(nil))
	getreq.Header.Add("X-Auth-Token", req.AkSkDetail.Token)
	getreq.Header.Set("Content-Type", "application/json")

	akskresp, err := client.Do(getreq)
	log.Info("RAJAT - RESPONSE  - " , akskresp)

	if err != nil {
		return nil, err
	}

	log.Info("RAJAT - Delete AKSK Response - ", akskresp)
	bodyBytes, _ := ioutil.ReadAll(akskresp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("AKSK API Response as String:\n" + bodyString)
	//defer akskresp.Body.Close()
	return akskresp, nil
}

func GetAKSK(ctx context.Context, req *pb.GetAkSkRequest) (*http.Response, error) {

	client := &http.Client{}

	getreq , err := http.NewRequest("GET", KEYSTONE_URL+"/"+req.GetId(), bytes.NewBuffer(nil))
	getreq.Header.Add("X-Auth-Token", req.AkSkDetail.Token)
	getreq.Header.Set("Content-Type", "application/json")

	akskresp, err := client.Do(getreq)
	log.Info("RAJAT - RESPONSE  - " , akskresp)

	defer akskresp.Body.Close()

	if err != nil {
		return nil, err
	}

	log.Info("RAJAT - Create AKSK Response - ", akskresp)
	bodyBytes, _ := ioutil.ReadAll(akskresp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("AKSK API Response as String:\n" + bodyString)

	return akskresp, nil
}
