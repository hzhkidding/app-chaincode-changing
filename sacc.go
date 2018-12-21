package  main

import (
	"encoding/json"
	"fmt"
	

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type  SimpleChaincode struct  {

}
type Tb_members struct {
	MemberID                 string `json:"memberid"`                 //用户ID
	Password                 string `json:"password"`               //用户名字
	Rank                     string    `json:"rank"` // 证件类型
	Credit                   string `json:"credit"`     //证件号码
	Phone                    string    `json:"phone"`                //性别
	ImgPath                  string `json:"imgpath"`           //生日
	
}
type Tb_business struct {
	BusinessID                string `json:"businessid"`                 //用户ID
	BusinessName               string `json:"businessname"`               //用户名字
	SendOutPrice               string    `json:"sendoutprice"` // 证件类型
	DistributionPrice         string `json:"distributionprice"`     //证件号码
	ShopHours                 string    `json:"shophours"`                //性别
	BusinessAddress           string `json:"businessaddress"`           //生日
	BusinessDepict             string `json:"businessDepict"`           //银行卡号
	Notice                 string `json:"notice"`
	BusinessScenery         string `json:"businessscenery"`
	Logo                      string   `json:"logo"`
	CategoryID                 string  `json:"categoryid"`
	
}
type Tb_category struct {
	CategoryID                 string `json:"categoryid"`                 //用户ID
	CategoryName               string `json:"categoryname"`               //用户名字
	
}
type Tb_menus struct {
	MenusID                 string `json:"menusid"`                 //用户ID
	MenusName               string `json:"menusname"`               //用户名字
	MenusPrice              string    `json:"menusprice"` // 证件类型
	MenusDepict             string `json:"menusdepict"`     //证件号码
	MenusImagePath                 string    `json:"menusimagepath"`                //性别
	BusinessID                 string `json:"businessid"`           //生日

}
type Tb_order struct {
	OrderID                 string `json:"id"`                 //用户ID
	CommitTime               string `json:"committime"`               //用户名字
	Amount             string    `json:"amount"` // 证件类型
	TotalPrice         string `json:"totalprice"`     //证件号码
	Status                 string    `json:"status"`                //性别
	MenusID             string `json:"menusid"`           //生日
	MembersID           string `json:"memberid"`           //银行卡号
	
}

var businessKey = "business_list"

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}
	return shim.Success(nil)
}
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	} else if function == "register" {
		// Deletes an entity from its state
		return t.register(stub, args)
	} else if function == "loginCheck" {
		// the old "Query" is now implemtned in invoke
		return t.loginCheck(stub, args)
	}else if function == "addBusiness" {
		// the old "Query" is now implemtned in invoke
		return t.addBusiness(stub, args)
	}else if function == "getAllBusiness" {
		// the old "Query" is now implemtned in invoke
		return t.getAllBusiness(stub, args)
	}else if function == "addMenus" {
		// the old "Query" is now implemtned in invoke
		return t.addMenus(stub, args)
	}else if function == "getMenusByBusinessID" {
		// the old "Query" is now implemtned in invoke
		return t.getMenusByBusinessID(stub, args)
	

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}
func (t *SimpleChaincode)  invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    return shim.Success(nil)
}

func (t *SimpleChaincode)  register(stub shim.ChaincodeStubInterface,args []string) pb.Response {
         var memberid string 
         var password string
         var rank string
         var credit string
         var phone string
         var imagpath string
         var  tb_members Tb_members

         memberid = args[0]
         password = args[1]
         rank = args[2]
         credit = args[3]
         phone = args[4]
         imagpath = args[5]
	fmt.Println("register")
	 if len(args) != 6 {
		return shim.Error("CreateUser：Incorrect number of arguments. Expecting 10")
	}

	
	tb_members.MemberID = memberid
	tb_members.Password = password
	tb_members.Rank =  rank
	tb_members.Credit = credit
	tb_members.Phone = phone
	tb_members.ImgPath =imagpath
        

	jsons_tb_members,err := json.Marshal(tb_members)
         fmt.Println(jsons_tb_members)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(args[0],jsons_tb_members)
	return shim.Success(jsons_tb_members)
   
}

func (t *SimpleChaincode)  loginCheck(stub shim.ChaincodeStubInterface,args []string) pb.Response {
       var id  string
       fmt.Println("exam logincheck")
	 if len(args) != 2 {
         return shim.Error("CreateUser：Incorrect number of arguments. Expecting 10")
	 }
      
      //var password string
     // var tb_membersbytes string
      id = args[0]
      //password = args[1]
       
       memberbytes,_ := stub.GetState(id)

       if(memberbytes !=nil){
                  return shim.Success(memberbytes)

       }
       return shim.Success(nil)

}

func (t *SimpleChaincode) addBusiness(stub shim.ChaincodeStubInterface,args []string) pb.Response{

     var businessID string
     
     var businessIDs []string
     
     var businessName string
     var sendOutPrice string
     var distributionPrice string
     var shopHours string
     var businessAddress string
     var businessDepict string
     var notice string
     var businessScenery string
     var logo string
     var categoryID string 

     var tb_business Tb_business   

     tb_business.BusinessID =businessID
     tb_business.BusinessName = businessName
     tb_business.SendOutPrice = sendOutPrice
     tb_business.DistributionPrice = distributionPrice
     tb_business.ShopHours = shopHours
     tb_business.BusinessAddress = businessAddress
     tb_business.BusinessDepict =businessDepict
     tb_business.Notice = notice
     tb_business.BusinessScenery = businessScenery
     tb_business.Logo = logo
     tb_business.CategoryID = categoryID

     allBusinessBytes,err := stub.GetState(businessKey)
     err = json.Unmarshal(businessIDs,&allBusinessBytes)
     businessIDs : = append (businessIDs,businessID)
     allBusinessBytes,err = json.Marshal(businessIDs)
     err = stub.PutState(businessKey,allBusinessBytes)

     businessBytes,_:= json.Marshal(tb_business)
     err := stub.PutState(businessID,businessBytes)


     return shim.Success(businessBytes)


}


func (t *SimpleChaincode)  getAllBusiness(stub shim.ChaincodeStubInterface,args []string) pb.Response {

	   businessListBytes,err := stub.GetState(businessKey)

	    var businessIds []string
	    err = json.Unmarshal(businessIds,&allBusinessBytes)

	    businessMap := []Tb_business{}
	    for index := range businessIds {

	       id :=businessIds[index]
	       businessBytes,err :=stub.GetState(id)
	       tb_business := Tb_business{}
	       err := json.Unmarshal(businessBytes,&tb_business)

	       businessMap = append (businessMap,tb_business)

	       }

	     allBusinessBytes,err :=json.Marshal(businessMap)
	     return shim.Success(allBusinessBytes)
               
}
func  (t *SimpleChaincode) addMenus(stub shim.ChaincodeStubInterface,args []string) pb.Response {

   
    var menusID string
    var menusName string
    var menusPrice string
    var menusDepict string
    var menusImagePath string
    var businessID string

    menusID  = args[0]
    menusName = args[1]
    menusPrice = args[2]
    menusDepict = args[3]
    menusImagePath = args[4]
    businessID = args[5]

    tb_menus := Tb_menus{MenusID:menusID,MenusName:menusName,MenusPrice:menusPrice,MenusDepict:menusDepict,MenusImagePath:menusImagePath,BusinessID:businessID}
    key,err :=stub.CreateCompositeKey("Menus~Business:",[]string{businessID,menusID})
     
     if err !=nil {
     	return shim.Error(err.Error())
     }
     tb_menusBytes,_ := json.Marshal(tb_menus)
     err = stub.PutState(key,tb_menusBytes)
     if err != nil {
     	return shim.Error(err.Error())
     }

     return shim.Success(tb_menusBytes)

}
	
}

func (t *SimpleChaincode) getMenusByBusinessID(stub shim.ChaincodeStubInterface,args []string) pb.Response {
      
      var businessID string

      businessID = args[0]

      tb_menusMap := []Tb_menus{}

      resultIterator, err := stub.GetStateByPartialCompositeKey("Menus~Business:", []string{businessID})
    	defer resultIterator.Close()
	    for resultIterator.HasNext() {
		item, _ := resultIterator.Next()
		fmt.Printf("key=%s\n", item.Key)
		tb_menusBytes, err := stub.GetState(item.Key)
		if err != nil {
			return shim.Error("Failed to get state")
		}
		tb_menus := Tb_menus{}
	   	err  = json.Unmarshal(tb_menusBytes, &tb_menus)
		if err != nil {
   			return shim.Error(err.Error())
   		}

	    tb_menusMap = append(tb_menusMap, tb_menus)
	}
	tb_menusMapBytes, err := json.Marshal(tb_menusMap)
	if err != nil {
		shim.Error("Failed to decode json of productMap")
	}
    return shim.Success(tb_menusMapBytes)
	               
	
}
func main() {

	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting simple chaincode %s",err)
	}

}


