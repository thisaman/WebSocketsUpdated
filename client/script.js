const url= "ws://localhost:8080/ws"
let socket = new WebSocket(url)


//Textarea 
var msgDiv = document.getElementById('msgDiv');
var msg = document.getElementById('input_text')
// Select Category 
var divCategory=document.getElementById('divCategory')
document.getElementById('btnStudent')
document.getElementById('btnTeacher')

//Register Yourself
var divStuReg=document.getElementById('divStuReg')
var rollno = document.getElementById('rollNo')
document.getElementById('btnSubmit')

//username & password
var divTeacherAuth=document.getElementById('divTeacherAuth')
var username=document.getElementById('username')
var password=document.getElementById('password')

//Teacher Table
var showDataDiv =document.getElementById('showData')
var tbl = document.getElementById('stuTable')
//data received
var rNo,wCount,charCount,wpmCount
var flag=0

function showStudent(){
  divCategory.style.display="none"
divStuReg.style.display = "block"
}

function getRollNo(){
  //rNo = rollNo.value; 
    divStuReg.style.display = "none";
    msgDiv.style.display= "block";
}
function showTeacher(){
  divTeacherAuth.style.display="block"
  divCategory.style.display="none"
}
function btnSubmit(){
  if(username.value=='admin' && password.value=='admin'){
    divTeacherAuth.style.display="none"
    showDataDiv.style.display="block"
}
}

socket.onopen =function(){
  console.log("Successfully connected")
    /*enabling textarea when websocket connection is opened */
    msg.disabled =false;
    msg.addEventListener('input', sendMessage)
    //socket.send("hi from the client")
}

  function sendMessage(){
    var text =msg.value;
    socket.send(text)
    //console.log(text)
  }
  socket.onclose =(event)=>{
    console.log("Socket Closed Connection", event);
    
  }

 socket.onmessage =(event) =>{
  //console.log("messsage  is ",event)
  var jsonObject = JSON.parse(event.data);
  wCount= jsonObject.Words
  charCount= jsonObject.Chars
  wpmCount =jsonObject.WordsPerMin
  rNo= rollno.value
  //console.log(rNo)
  //console.log(wCount)
  //console.log(charCount)
  //console.log(wpmCount)
  if(flag ==0){
    addItemsToTable(rNo,wCount,charCount,wpmCount)
    flag=1
  }
  if(flag ==1){
    addItemsToTable(rNo,wCount,charCount,wpmCount)
    flag=0
  }
  
  } 

  socket.onerror =(error)=>{
    console.log("Socket Error", error);
  }
  
  /* Creating table in which student data is to be displayed*/
  function addItemsToTable( rollNo, words, chars, wordsPM){
   var tblBody,row, _rollNo, _words, _chars, _wordsPM;
   if(flag==0){
     tblBody = document.createElement("tbody")
     row = document.createElement('tr')
     _rollNo = document.createElement('td')
     _words = document.createElement('td')
     _chars = document.createElement('td')
     _wordsPM = document.createElement('td')
    _rollNo.id= _rollNo;
    _words.id= _words;
    _chars.id= _chars;
    _wordsPM.id= _wordsPM;
    _rollNo.innerHTML =rollNo;
    _words.innerHTML= words;
    _chars.innerHTML =chars;
    _wordsPM.innerHTML =wordsPM;

    row.appendChild(_rollNo)
    row.appendChild(_words)
    row.appendChild(_chars)
    row.appendChild(_wordsPM)

    tblBody.appendChild(row)
    tbl.appendChild(tblBody)
    }
    
    }
  