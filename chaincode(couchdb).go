package  main

import (
	"encoding/json"
	"fmt"
	

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type  SimpleChaincode struct  {

}
type question struct {
	ObjectType string  `json:"docType"`
	ID                 string `json:"id"`                 //用户ID
	TITLE                 string `json:"title"`               //用户名字
	CONTENT                     string    `json:"content"` // 证件类型
	USER_ID                   string `json:"user_id"`     //证件号码
	CREAT_DATE                    string    `json:"create_date"`                //性别
	COMMENT_COUNT                  string `json:"comment_count"`           //生日
	
}
type user struct {
	ObjectType string `json:"docType"`
	ID                string `json:"id"`                 //用户ID
	NAME               string `json:"name"`               //用户名字
	PASSWORD               string    `json:"password"` // 证件类型
	SALT         string `json:"salt"`     //证件号码
	HEAD_URL                 string    `json:"head_url"`                //性别

	
}
type login_ticket struct {
	ObjectType string  `json:"docType"`
	ID                 string `json:"id"`                 //用户ID
	USER_ID               string `json:"user_id"`               //用户名字
	TICKET               string `json:"ticket"`
	EXPIRED               string `json:"expired"`
	STATUS               string `json:"status"`
	
}
type comment struct {
	ObjectType string `json:"docType"`
	ID                 string `json:"id"`                 //用户ID
	CONTENT               string `json:"content"`               //用户名字
	USER_ID              string    `json:"user_id"` // 证件类型
	ENTITY_ID             string `json:"entity_id"`     //证件号码
	ENTITY_TYPE                 string    `json:"entity_type"`                //性别
	CREATED_DATE                 string `json:"create_date"`           //生日
	STATUS                 string `json:"status"`

}
type message struct {
	ObjectType string `json:"docType"`
	ID                 string `json:"id"`                 //用户ID
	FROM_ID               string `json:"from_id"`               //用户名字
	TO_ID             string    `json:"to_id"` // 证件类型
	CONTENT         string `json:"content"`     //证件号码
	CREATED_DATE                 string    `json:"created_date"`                //性别
	HAS_READ             string `json:"has_read"`           //生日
	CONVERSATION_ID           string `json:"conversation_id"`           //银行卡号
	
}
type feed struct {
	ObjectType string `json:"docType"`
	ID                 string `json:"id"`                 //用户ID
	CREATED_DATE               string `json:"created_date"`               //用户名字
	USER_ID             string    `json:"user_id"` // 证件类型
	DATE         string `json:"date"`     //证件号码
	TYPE                 string    `json:"type"`                //性别
}

var userid = 0
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
    }else if function == "addOrders" {
		// the old "Query" is now implemtned in invoke
		return t.addOrders(stub, args)
    }else if function == "getOrdersByMemberID" {
		// the old "Query" is now implemtned in invoke
		return t.getOrdersByMemberID(stub, args)
    }
	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}
func (t *SimpleChaincode)  invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    return shim.Success(nil)
}

func (t *SimpleChaincode)  addUser(stub shim.ChaincodeStubInterface,args []string) pb.Response {
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
    // var err error
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

    businessID = args[0]
      businessName = args[1]
     sendOutPrice = args[2]
     distributionPrice = args[3]
     shopHours = args[4]
     businessAddress = args[5]
     businessDepict = args[6]
     notice = args[7]
     businessScenery = args[8]
     logo = args[9]
     categoryID = args[10]

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

     allBusinessBytes,_ := stub.GetState(businessKey)
      json.Unmarshal(allBusinessBytes,&businessIDs)
     businessIDs = append (businessIDs,businessID)
     allBusinessBytes,_= json.Marshal(businessIDs)
    stub.PutState(businessKey,allBusinessBytes)

     businessBytes,_:= json.Marshal(tb_business)
     stub.PutState(businessID,businessBytes)


     return shim.Success(businessBytes)


}


func (t *SimpleChaincode)  getAllBusiness(stub shim.ChaincodeStubInterface,args []string) pb.Response {
     
	   businessListBytes,_ := stub.GetState(businessKey)

	    var businessIds []string
	    json.Unmarshal(businessListBytes,&businessIds)

	    businessMap := []Tb_business{}
	    for index := range businessIds {

	       id :=businessIds[index]
	       businessBytes,_ :=stub.GetState(id)
	       tb_business := Tb_business{}
	       json.Unmarshal(businessBytes,&tb_business)

	       businessMap = append (businessMap,tb_business)

	       }

	     allBusinessBytes,_ :=json.Marshal(businessMap)
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
func (t *SimpleChaincode)  addOrders(stub shim.ChaincodeStubInterface,args []string) pb.Response {
	var orderID       string          
	var commitTime         string     
	var amount                string
	var totalPrice              string
	var status                  string
	var menusID                 string
	var membersID               string


	orderID = args[0]
	commitTime = args[1]
	amount = args[2]
	totalPrice = args[3]
	status = args[4]
	menusID = args[5]
	membersID = args[6]

     tb_order := Tb_order{OrderID:orderID,CommitTime:commitTime,Amount:amount,TotalPrice:totalPrice,Status:status,MenusID:menusID,MembersID:membersID}

	 key,err :=stub.CreateCompositeKey("Order~Member:",[]string{membersID,orderID})
     
     if err !=nil {
     	return shim.Error(err.Error())
     }
     tb_orderBytes,_ := json.Marshal(tb_order)
     err = stub.PutState(key,tb_orderBytes)
     if err != nil {
     	return shim.Error(err.Error())
     }

     return shim.Success(tb_orderBytes)


}
func (t *SimpleChaincode)  getOrdersByMemberID(stub shim.ChaincodeStubInterface,args []string) pb.Response {
     var membersID string
     membersID = args[0]
     tb_orderMap := []Tb_order{}

      resultIterator, err := stub.GetStateByPartialCompositeKey("Order~Member:", []string{membersID})
    	defer resultIterator.Close()
	    for resultIterator.HasNext() {
		item, _ := resultIterator.Next()
		fmt.Printf("key=%s\n", item.Key)
		tb_orderBytes, err := stub.GetState(item.Key)
		if err != nil {
			return shim.Error("Failed to get state")
		}
		tb_order := Tb_order{}
	   	err  = json.Unmarshal(tb_orderBytes, &tb_order)
		if err != nil {
   			return shim.Error(err.Error())
   		}

	    tb_orderMap = append(tb_orderMap, tb_order)
	}
	tb_orderMapBytes, err := json.Marshal(tb_orderMap)
	if err != nil {
		shim.Error("Failed to decode json of productMap")
	}
    return shim.Success(tb_orderMapBytes)
	               


}
 func main() {

	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting simple chaincode %s",err)
	}

}


