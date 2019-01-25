package  main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	//"strings"
	//"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type  SimpleChaincode struct  {

}
type question struct {
	OBJECTTYPE string  `json:"docType"`
	ID                 string `json:"id"`                 //用户ID
	TITLE                 string `json:"title"`               //用户名字
	CONTENT                     string    `json:"content"` // 证件类型
	USER_ID                   string `json:"user_id"`     //证件号码
	CREAT_DATE                    string    `json:"create_date"`                //性别
	COMMENT_COUNT                  string `json:"comment_count"`           //生日
	
}
type user struct {
	OBJECTTYPE string `json:"docType"`
	ID                string `json:"id"`                 //用户ID
	NAME               string `json:"name"`               //用户名字
	PASSWORD               string    `json:"password"` // 证件类型
	SALT         string `json:"salt"`     //证件号码
	HEAD_URL                 string    `json:"head_url"`                //性别

	
}
type login_ticket struct {
	OBJECTTYPE string  `json:"docType"`
	ID                 string `json:"id"`                 //用户ID
	USER_ID               string `json:"user_id"`               //用户名字
	TICKET               string `json:"ticket"`
	EXPIRED               string `json:"expired"`
	STATUS               string `json:"status"`
	
}
type comment struct {
	OBJECTTYPE string `json:"docType"`
	ID                 string `json:"id"`                 //用户ID
	CONTENT               string `json:"content"`               //用户名字
	USER_ID              string    `json:"user_id"` // 证件类型
	ENTITY_ID             string `json:"entity_id"`     //证件号码
	ENTITY_TYPE                 string    `json:"entity_type"`                //性别
	CREATED_DATE                 string `json:"create_date"`           //生日
	STATUS                 string `json:"status"`

}
type message struct {
	OBJECTTYPE         string `json:"docType"`
	ID                 string `json:"id"`
	FROM_ID            string `json:"from_id"`
	TO_ID              string    `json:"to_id"`
	CONTENT            string `json:"content"`
	CREATED_DATE       string    `json:"created_date"`
	HAS_READ           string `json:"has_read"`
	CONVERSATION_ID    string `json:"conversation_id"`
}

type feed struct {
	OBJECTTYPE string `json:"docType"`
	ID                 string `json:"id"`                 //用户ID
	CREATED_DATE               string `json:"created_date"`               //用户名字
	USER_ID             string    `json:"user_id"` // 证件类型
	DATE         string `json:"date"`     //证件号码
	TYPE                 string    `json:"type"`                //性别
}
var userId = 0
var questionId = 0
var commentId=0
var messageId=0
var loginTicketId = 0
var feedId = 0
var userPrefix = "user"
var questionPrefix = "question"
var commentPrefix = "comment"
var messagePrefix = "message"
var feedPrefix = "feed"
var logintickerPrefix = "loginTicket"
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
	} else if function == "addUser" {
		// Deletes an entity from its state
		return t.addUser(stub, args)
	} else if function == "selectById" {
		// the old "Query" is now implemtned in invoke
		return t.selectById(stub, args)
	}else if function == "selectByName" {
		// the old "Query" is now implemtned in invoke
		return t.selectByName(stub, args)
	}else if function == "updatePassword" {
		// the old "Query" is now implemtned in invoke
		return t.updatePassword(stub, args)
	}else if function == "deleteById" {
		// the old "Query" is now implemtned in invoke
		return t.deleteById(stub, args)
	}else if function == "queryString" {
		return t.queryString(stub, args)
	}else if function == "addMessage" {
		return t.addMessage(stub, args)
	}else if function == "getConversationDetail" {
		return t.getConversationDetail(stub,args)
	}
	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}
func (t *SimpleChaincode)  invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    return shim.Success(nil)
}
func (t *SimpleChaincode) selectByName(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	// var err error
	name := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"user\",\"name\":\"%s\"}}", name)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

}

func (t *SimpleChaincode) selectById(stub shim.ChaincodeStubInterface,args []string) pb.Response {

	var id = userPrefix+args[0]
	userbytes,err := stub.GetState(id)
	if err != nil {
		return shim.Error("Failed")
	}
	return shim.Success(userbytes)
}

func (t *SimpleChaincode)  updatePassword(stub shim.ChaincodeStubInterface,args []string) pb.Response {
   // var err error
	var id = userPrefix+args[0]
	userbytes,_ := stub.GetState(id)
	var tb_user user
	json.Unmarshal(userbytes,&tb_user)
	tb_user.PASSWORD = args[2]

	userbytes,_ = json.Marshal(tb_user)

	return shim.Success(userbytes)


}
func  (t *SimpleChaincode) deleteById(stub shim.ChaincodeStubInterface,args []string) pb.Response {

   var err error
	var id =  userPrefix+args[0]
	err = stub.DelState(id) //remove the marble from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}
	return shim.Success(nil)

}
func main() {
	err := shim.Start(new(SimpleChaincode))
	fmt.Println("hahahha");
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
func (t *SimpleChaincode)  addUser(stub shim.ChaincodeStubInterface,args []string) pb.Response {
	     var tb_user user
         objectType := "user"
         id := userPrefix+strconv.Itoa(userId)
         userId++
	     数据读取与序列化
         name := args[1]
         password := args[2]
         salt := args[3]
         head_url := args[4]
         tb_user = user{OBJECTTYPE:objectType,ID: id,
         NAME:name,PASSWORD:password,SALT:salt,HEAD_URL:head_url}
         userbytes,_ := json.Marshal(tb_user)
	     插入数据
	     stub.PutState(id,userbytes)
	     fmt.Println("addUser")
         return shim.Success([]byte("5"))

}
func (t *SimpleChaincode)  addQuestion(stub shim.ChaincodeStubInterface,args []string) pb.Response {

	var tb_question question
	objectType := "question"
	id := questionPrefix+strconv.Itoa(questionId)
	questionId++
	title := args[0]
	content := args[1]
	created_date := args[2]
	user_id  := args[3]
	comment_count := args[4]
	tb_question = question{OBJECTTYPE:objectType,ID: id,TITLE:title,CONTENT:content,CREAT_DATE:created_date,USER_ID:user_id,COMMENT_COUNT:comment_count}
	questionbytes,_ := json.Marshal(tb_question)

	stub.PutState(id,questionbytes)
	fmt.Println("addQuestion")
	return shim.Success([]byte("1"))

}
func (t *SimpleChaincode) getById(stub shim.ChaincodeStubInterface,args []string) pb.Response {

	id := questionPrefix+args[0]
	questionbytes,err := stub.GetState(id)
	if err != nil {
		return shim.Error("Failed")
	}
	return shim.Success(questionbytes)

}
func (t *SimpleChaincode)  updateCommentCount(stub shim.ChaincodeStubInterface,args []string) pb.Response {
	// var err error
	id := questionPrefix+args[0]
	questionbytes,_ := stub.GetState(id)
	var tb_question question
	json.Unmarshal(questionbytes,&tb_question)
	tb_question.COMMENT_COUNT = args[2]
	questionbytes,_ = json.Marshal(tb_question)
	return shim.Success(questionbytes)

}
func (t *SimpleChaincode) selectLatestQuestions(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	// var err error
	var queryString string
	if args[0] == "" {
		queryString = fmt.Sprintf("{\"selector\":{\"docType\":\"question\"}}")
	} else {
	user_id := userPrefix+args[0]
	queryString = fmt.Sprintf("{\"selector\":{\"docType\":\"question\",\"name\":\"%s\"}}", user_id)
	}
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

}


func (t *SimpleChaincode)  addMessage(stub shim.ChaincodeStubInterface,args []string) pb.Response {
	var tb_message message
	objectType := "message"
	id := messagePrefix+strconv.Itoa(messageId)
	messageId++
	fromId := args[0]
	toId := args[1]
	content := args[2]
	hasRead  := args[3]
	conversationId := args[4]
	createDate := args[5]
	tb_message = message{OBJECTTYPE:objectType,ID: id,FROM_ID:fromId,TO_ID:toId,CONTENT:content,HAS_READ:hasRead,CONVERSATION_ID:conversationId,CREATED_DATE:createDate}
	messagebytes,_ := json.Marshal(tb_message)

	stub.PutState(id,messagebytes)
	fmt.Println("addMessage")
	return shim.Success([]byte("1"))

}
func (t *SimpleChaincode)  getConversationDetail(stub shim.ChaincodeStubInterface,args []string) pb.Response {
//	@Select({"select ", SELECT_FIELDS, " from ", TABLE_NAME, " where conversation_id=#{conversationId} order by id desc limit #{offset}, #{limit}"})
       conversationId := args[0]  //  1_2  要处理一下
     // var  tb_message message
	 queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"message\",\"conversion_id\":\"%s\"}, " +
	 	"\"sort\": [{\"conversation_id\": \"desc\"}]，\"limit\":\"%s\",\"skip\":\"%s\"}", conversationId)
	 queryResults, err := getQueryResultForQueryString(stub, queryString)
	 if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)



}
func (t *SimpleChaincode)  getConvesationUnreadCount(stub shim.ChaincodeStubInterface,args []string) pb.Response {
	//	@Select({"select ", SELECT_FIELDS, " from ", TABLE_NAME, " where conversation_id=#{conversationId} order by id desc limit #{offset}, #{limit}"})
	conversationId := args[0]  //  1_2  要处理一下
	// var  tb_message message
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"question\",\"conversation_id\":\"%s\"}}", conversationId)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)



}
//@Select({"select ", INSERT_FIELDS, " ,count(id) as id from ( select * from ", TABLE_NAME, " where from_id=#{userId} or to_id=#{userId} order by id desc) tt group by conversation_id  order by created_date desc limit #{offset}, #{limit}"})
func (t *SimpleChaincode) getConversationList(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	// var err error
	userId := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"message\",\"$or\": [{\"from_id\": \"%s\" },{ \"to_id\": \"%s\" }]}}",userId)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

}
func (t *SimpleChaincode) queryString(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	// var err error
	queryString := args[0]
	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return  shim.Error(err.Error())
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())

}
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	var sum = 0
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		sum++
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}
func constructQueryResponseFromIteratorCount(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	var sum = 0
	buffer.WriteString("[")

	//bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		 resultsIterator.Next()
            sum++
	}
	buffer.WriteString(string(sum))
	buffer.WriteString("]")

	return &buffer, nil
}
