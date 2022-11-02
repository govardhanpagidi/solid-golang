package main
import "fmt"
import "encoding/json"

//Product interface
type Product interface{
	Price() float64
	Description() string

}

//Product 1
type SmokeSensor struct{
	TotalPrice float64
	Discount float64
}

func (s *SmokeSensor) Price() float64{
	return (s.TotalPrice - s.Discount)
}

func (s *SmokeSensor) Description() string{
	return "Smoke and CO sensor with With WiFi connectivity for Remote Monitoring, price is "+fmt.Sprintf("%f",s.Price())
}

//Product 2
type Camera struct{
	TotalPrice float64
	Discount float64
}

func (c *Camera) Price() float64{
	return (c.TotalPrice - c.Discount)
}

func (c *Camera) Description() string{
	return "Smoke and CO sensor with With WiFi connectivity for Remote Monitoring"
}



//Reusable Functions
type ProductDetails struct{}

func (pd *ProductDetails) Text( p Product) string{
	return "Price:"+fmt.Sprintf("%f",p.Price())+ " "+ "Description:"+p.Description()
}
func (pd *ProductDetails) Json() string{
	data,_ := json.Marshal(pd)
	return string(data)
}

func main(){
	s := &SmokeSensor{
		TotalPrice : 150,
		Discount : 5,
	}

	pDetails := ProductDetails{}
	fmt.Println("Product Details:", pDetails.Text(s))
}
