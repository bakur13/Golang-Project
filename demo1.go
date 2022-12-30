
func main(){
	products,_:=GetAllProducts()

       for i := 0; i < len(products) ; i++ {
		fmt.Println(products[i],ProductName)
	   }

	
   product,_:= AddProducts()
   fmt.Println(product)
   
}
type products structs{

	Id int `json:"id"`
	productName string  `json:"productName"`
	CategoryId int  `json:"category"`
	UnitPrice float64  `json:"UnitPrice"`
}
type Category structs{

	Id int `json:"id"`
	CategoryName string  `json:"categoryName"`
}


func GetAllProducts() ([]Product,error) {
	response,err:=http.Get("http://localhost:3000/products")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	bodyBytes,_:=ioutil.ReadAll(response.Body)

	var products []Products 
	
	json.Unmarshal(bodyBytes,&product)
	return products,nil
}


func AddProducts() (Product,error) {
	product:=Product{Id:4,ProductName:"Telephone",CategoryId:1,UnitPrice:5617}
	jsonProduct,err:=json.Marshal(product)
	
	response,err:=http.Post("http://localhost:3000/products","application/json;charset=utf-8",bytes.NewBuffer(jsonProduct))
if err != nil {
	return product{}, err
}
defer response.Body.Close()

return productResponse,err

}