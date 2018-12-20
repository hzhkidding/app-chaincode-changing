package  main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type  SimpleChaincode struct  {

}
type Tb_members struct {
	memberId                 string `json:"memberid"`                 //用户ID
	password                 string `json:"password"`               //用户名字
	rank                     string    `json:"rank"` // 证件类型
	credit                   string `json:"credit"`     //证件号码
	phone                    string    `json:"phone"`                //性别
	imgPath                  string `json:"imgpath"`           //生日
	
}
type Tb_business struct {
//	businessID                string `json:"businessid"`                 //用户ID
	businessName               string `json:"businessname"`               //用户名字
	sendOutPrice               string    `json:"sendoutprice"` // 证件类型
	distributionPrice         string `json:"distributionprice"`     //证件号码
	shopHours                 string    `json:"shophours"`                //性别
	businessAddress           string `json:"businessaddress"`           //生日
	businessDepict             string `json:"businessDepict"`           //银行卡号
	notice                 string 'json:"notice"'
	businessScenery         string 'json:"businessscenery"'
	logo                      string   'json:"logo"'
	categoryID                 string  'json:"categoryid"'
	
}
type Tb_category struct {
	categoryID                 string `json:"categoryid"`                 //用户ID
	categoryName               string `json:"categoryname"`               //用户名字
	
}
type Tb_menus struct {
	menusID                 string `json:"menusid"`                 //用户ID
	menusName               string `json:"menusname"`               //用户名字
	menusPrice              string    `json:"menusprice"` // 证件类型
	menusDepict             string `json:"menusdepict"`     //证件号码
	menusImagePath                 string    `json:"menusimagepath"`                //性别
	businessID                 string `json:"businessid"`           //生日

}
type Tb_order struct {
	orderID                 string `json:"id"`                 //用户ID
	commitTime               string `json:"committime"`               //用户名字
	amount             string    `json:"amount"` // 证件类型
	totalPrice         string `json:"totalprice"`     //证件号码
	status                 string    `json:"status"`                //性别
	menusID             string `json:"menusid"`           //生日
	memberID           string `json:"memberid"`           //银行卡号
	
}

var buSinessKey = "business_list"

func (t *SimpleChaincode)  register(stub shim.ChaincodeStubInterface,args []string) pb.Response {

	fmt.Println("register")
	 if len(args) != 6 {
		return shim.Error("CreateUser：Incorrect number of arguments. Expecting 10")
	}

	var  tb_members Tb_members
	tb_members.memberID = args[0]
	tb_members.password = arg[1]
	tb_members.rank = args[2]
	tb_members.credit = args[3]
	tb_members.phone = args[4]
	tb_members.imgpath =args[5]

	jsons_tb_members,err := json.Marshal(tb_members)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(args[0],jsons_tb_members)
	return shim.Success("success")
   
}

func (t *SimpleChaincode)  loginCheck(stub shim.ChaincodeStubInterface,args []string) pb.Response {
       
       fmt.Println("exam logincheck")
	 if len(args) != 2 {
		return shim.Error("CreateUser：Incorrect number of arguments. Expecting 10")
	}
      var id  string
      var password string
 
      id = args[0]
      password = args[1]
       
       tb_members,err = stub.GetState(id)

       if(tb_members !=nil){
                  return shim.Success("1")

       }

}

func (t *SimpleChaincode) addBusiness(stub shim.ChaincodeStubInterface,args []string) pb.Response{

     var businessID string
     var buSinessIds []string
     businessID := args[0]                

     allBuSinessBytes,err := stub.GetState(buSinessKey)
     err = json.Unmarshal(buSinessIds,&allBuSinessBytes)

     buSinessIds : = append (buSinessIds,businessID)
     allBuSinessBytes,err = json.Marshal(buSinessIds)
     err = stub.PutState(buSinessKey,allBuSinessBytes)

     var tb_business  Tb_business
   
     tb_business.businessName = args [1]
     tb_business.sendOutPrice = args [2]
     tb_business.distributionPrice = args [3]
     tb_business.shopHours = args [4]
     tb_business.businessAddress = args [5]
     tb_business.businessDepict = arg[6]
     tb_business.notice = args[7]
     tb_business.businessScenery = args[8]
     tb_business.logo = args[9]
     tb.categoryID = arg[10]
     
     buSinessBytes := json.Marshal(tb_business)
     err := stub.PutState(businessID,buSinessBytes)

     return shim.Success("Success")




}


func (t *SimpleChaincode)  getAllBusiness(stub shim.ChaincodeStubInterface,args []string) pb.Response {

	            allBuSinessBytes,err := stub.GetState(buSinessKey)

	            var buSinessIds []string
	            err = json.Unmarshal(buSinessIds,&allBuSinessBytes)

	            buSinessMap : []Tb_business{}
	            for index := range buSinessIds {

	            	id :=buSinessIds[index]
	            	buSinessBytes,err :=stub.GetState(id)
	            	tb_business := Tb_business{}

	            	buSinessMap = append (buSinessMap,tb_business)

	            }

	            allBuSinessJson,err :=json.Marshal(buSinessMap)
	            return shim.Success(allBuSinessJson)
               
}

func (t *SimpleChaincode) getMenus(stub shim.ChaincodeStubInterface,args []string) pb.Response {

	               
	
}
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nul {
		fmt.Printf("Error starting simple chaincode %s",err)
	}
}

