package routes

import (
	cartcontroller "blanja_api/src/controllers/CartController"
	categorycontroller "blanja_api/src/controllers/CategoryController"
	ordercontroller "blanja_api/src/controllers/OrderController"
	paymentcontroller "blanja_api/src/controllers/PaymentController"
	productcontroller "blanja_api/src/controllers/ProductController"
	usercontroller "blanja_api/src/controllers/UserController"
	"blanja_api/src/middleware"
	"fmt"
	"net/http"

	"github.com/goddtriffin/helmet"
)

func Router() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Nothing Here , Try Another Page")
	})

	//Routes User

	//login
	http.HandleFunc("/login", usercontroller.Login)

	//register
	http.Handle("/register-seller", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(usercontroller.SellerRegister))))
	http.Handle("/register-customer", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(usercontroller.CustomerRegister))))
	
	//update
	http.Handle("/update-seller", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(usercontroller.Update_seller))))
	http.Handle("/update-customer", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(usercontroller.Update_customer))))
	
	//get
	http.Handle("/users", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(usercontroller.Data_users))))
	http.Handle("/user/", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(usercontroller.Data_user))))
	
	//Routes Product
	http.Handle("/products", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(productcontroller.Data_products))))
	http.Handle("/product/", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(productcontroller.Data_product))))
	
	//Routes Category
	http.Handle("/categories", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(categorycontroller.Data_categories))))
	http.Handle("/category/", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(categorycontroller.Data_category))))
	
	//Routes Cart
	http.Handle("/carts", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(cartcontroller.Data_carts))))
	http.Handle("/cart/", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(cartcontroller.Data_cart))))
	
	//Routes Order
	http.Handle("/orders", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(ordercontroller.Data_orders))))
	http.Handle("/order/", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(ordercontroller.Data_order))))
	
	//Routes Payment
	http.Handle("/payments", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(paymentcontroller.Data_payments))))
	http.Handle("/payment/", helmet.Default().Secure(middleware.XssMiddleware(http.HandlerFunc(paymentcontroller.Data_payment))))

}
