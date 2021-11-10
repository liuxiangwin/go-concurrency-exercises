package channel_singleton 

type singleton struct {}

var instance singleton
var addCh chan bool =make(chan bool)
var getCountCh chan chan int = make(chan chan int)
var quitCh chan bool = make(chan bool)

func GetInstance() *singleton{

   return &instance
}


func (s*singleton) GetCount()  int{
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh
	
}

func (s*singleton) AddOne(){

   addCh <- true
}

func (s*singleton) Stop(){

   quitCh <- true
   close(addCh)
   close(getCountCh)
   close(quitCh)
}


func init(){
  var count int
  go func(addCh <-chan bool, getCountCh<-chan chan int, quitCh <-chan bool){
     for{
			select{
			case <-addCh:
				   count++
			case ch:= <-getCountCh:
				   ch <- count
			case <-quitCh:
				   break
			
			}
	 
	 }
  
  }(addCh,getCountCh,quitCh)

}


