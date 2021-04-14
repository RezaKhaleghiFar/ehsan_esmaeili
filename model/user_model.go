package model
import(

)

type Error struct {
	Is_error  bool
	Error_cod int
}

//104=not fonde user;
//101=fild all date;
//102=duplicat username;
//103=error server;
//200=succesfull
//105=در یک روز بیشتر از سه بار تقاضای بیام
//106=token is expired

type User struct {

	User_Phon       string
	Name            string
	Last_Name       string
	Post_Cod        string
	State           string
	City            string
	Locition        string
	Shaba           string
	Energy_Count    int
	Samsicoin_Count float64
	Alert_Count     int
	User_Id        string
	Active           bool
}
type GetaUser struct {
	Error_cod     int
	Id string
}

type Aid struct {
	Aid_Id          int
	User_Id         string
	Pric            int
	Aid_description string
	Transaction_id  string
}
type Alert struct {
	Alert_Id int
	User_Id  string
	Pric     int
	Coin_Id  int
	Activ    bool
	Time     string
}
type Buy_Charg struct {
	Buy_Charge_Id int
	User_Id       string
	Pric          int
	Phone_Number  string
}
type Buy_Pack struct {
	Buy_Pack_Id    int
	User_Id        string
	Transaction_Id string
	Pack_Id        int
	Time           string
}
type Buy_Samsicoin struct {
	Buy_Pack_Id     int
	User_Id         string
	Transaction_Id  string
	Samsicoin_Count float32
	Pric            int
}
type Clearing struct {
	Clearing_Id     int
	User_Id         string
	Time            string
	Samsicoin_Count float32
	Price            int
}
type Energy struct {
	Energy_Id      int
	User_Id        string
	Description    string
	Low_High       bool
	Low_High_Count int
}
type Samsicoin struct {
	Samsicoin_Id   int
	User_Id        string
	Description    string
	Low_High       bool
	Low_High_Count int
	Transaction_Id string
	Typ            string
}
type Transaction struct {
	Transaction_Id           string
	Transaction_User_Id    string
	Transaction_Time         string
	Transaction_Price        int
	Transaction_IsSuccessful bool
	Tracking_Number          string
}
type User_Coin struct {
	User_Coin_Id int
	User_Id      string
	Coin_Id      int
}
type Suport struct {
	Support_Id int
	Support_Name      string
	Support_Link      string
}
type Product struct {
	Product_Id int
	Product_Name      string
	Product_Description      string
	Pric      int
	Discount      float32	
	Final_Price      int
	Image      string
}
type Pack struct {
	Pack_Id int
	Alert_Count      int
	Pric_Toman      int
	Pric_Samsicoin      float32
}
type Message struct {
	Message_Id int
	Titel      string
	Text      string
	Time      string
	User_Id      string
}
type Draw struct {
	Draw_Id int
	Draw_Titel      string
	Draw_Description      string
	Time      string
	Date      string
}
type Coin struct {
	Coin_Id int
	Coin_Name      string
	Pric      float32
	Image      string
	Abb_Name      string
}
type Samsicoin_List struct {
	List_Id int
	Samsicoin_Count      float32
	Pric      int

}

